package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MCSManager/pty/core"
)

var Dir, Cmd string

func init() {
	flag.StringVar(&Dir, "dir", "", "command work path")
	flag.StringVar(&Cmd, "cmd", "", "command")
	flag.StringVar(&core.PtySize, "size", "", "command")
}

func main() {
	flag.Parse()
	Pty, err := core.Start(Dir, Cmd)
	if err != nil {
		fmt.Printf("[MCSMANAGER-TTY] Process Start Error:%s\n", err)
		os.Exit(-1)
	}
	defer Pty.Close()

	if core.PtySize == "" {
		Pty.Setsize(50, 50)
	} else {
		Pty.ResizeWindow(&core.PtySize)
	}
	Pty.HandleStdIO()
}
