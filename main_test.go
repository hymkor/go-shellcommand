package shellcommand_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hymkor/go-shellcommand"
)

func TestSystem(t *testing.T) {
	workpath := filepath.Join(os.TempDir(), "work.txt")

	process, err := shellcommand.System(fmt.Sprintf(`echo "ahaha" > "%s"`, workpath))
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	process.Wait()

	output, err := os.ReadFile(workpath)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	var expect string
	if runtime.GOOS == "windows" {
		expect = "\"ahaha\" \r\n"
	} else {
		expect = "ahaha\n"
	}
	if string(output) != expect {
		t.Fatalf("result = `%s` (expect %s)", string(output), expect)
		return
	}
	os.Remove(workpath)
}
