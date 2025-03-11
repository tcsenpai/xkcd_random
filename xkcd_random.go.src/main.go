package main

import (
	"bunster-build/runtime"
	"os"
)

func main() {

	shell := runtime.Shell{
		PID: os.Getpid(),

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,

		Args: os.Args,

		Main: Main,
	}

	os.Exit(shell.Run())
}
