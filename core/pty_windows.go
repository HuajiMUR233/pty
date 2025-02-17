//go:build windows

package core

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/iamacarpet/go-winpty"
)

type Pty struct {
	tty    *winpty.WinPTY
	StdIn  *os.File
	StdOut *os.File
}

func getExecutableFilePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	if executableFileExists(filepath.Dir(ex)+"/winpty-agent.exe") && executableFileExists(filepath.Dir(ex)+"/winpty.dll") {
		return filepath.Dir(ex), nil
	} else {
		return filepath.Dir(ex), errors.New("[MCSMANAGER-TTY] ExecutableFile {winpty-agent.exe,winpty.dll} does not exist")
	}
}

func executableFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func Start(dir, command string) (*Pty, error) {
	path, err := getExecutableFilePath()
	if err != nil {
		return nil, err
	}
	tty, err := winpty.OpenWithOptions(winpty.Options{
		DLLPrefix: path,
		Command:   command,
		Dir:       dir,
		Env:       os.Environ(),
	})
	return &Pty{tty: tty, StdIn: tty.StdIn, StdOut: tty.StdOut}, err
}

func (pty *Pty) Write(p []byte) (n int, err error) {
	return pty.tty.StdIn.Write(p)
}

func (pty *Pty) Read(p []byte) (n int, err error) {
	return pty.tty.StdOut.Read(p)
}

func (pty *Pty) Setsize(cols, rows uint32) error {
	pty.tty.SetSize(cols, rows)
	return nil
}

func (pty *Pty) Close() error {
	pty.tty.Close()
	return nil
}
