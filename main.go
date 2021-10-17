package shellcommand

import (
	"os"
)

// System is the function like `system` function of the programming
// language C.
func System(cmdline string) (*os.Process, error) {
	return system(cmdline)
}
