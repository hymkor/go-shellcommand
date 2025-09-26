//go:build run

package main

import (
	"github.com/hymkor/go-shellcommand"
)

func main() {
	process, err := shellcommand.System(`echo "ahaha" > "test"`)
	if err != nil {
		println(err.Error())
		return
	}
	process.Wait()
}
