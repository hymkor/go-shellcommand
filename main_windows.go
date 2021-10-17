package shellcommand

import (
	"os"
	"syscall"
)

func startProcess(appName, commandLine string, stdin, stdout, stderr *os.File) (*os.Process, error) {
	return os.StartProcess(
		appName,
		nil,
		&os.ProcAttr{
			Files: []*os.File{stdin, stdout, stderr},
			Sys:   &syscall.SysProcAttr{CmdLine: commandLine},
		})
}

func system(cmdline string) (*os.Process, error) {
	cmdexe := os.Getenv("COMSPEC")
	return startProcess(cmdexe, `/S /C "`+cmdline+`"`, os.Stdin, os.Stdout, os.Stderr)
}
