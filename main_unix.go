//go:build !windows
// +build !windows

package shellcommand

import (
	"os"
)

func system(cmdline string) (*os.Process, error) {
	shell := os.Getenv("SHELL")
	return os.StartProcess(
		shell,
		[]string{shell, "-c", cmdline},
		&os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		})
}
