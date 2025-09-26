go-shellcommand
===============

go-shellcommand is the package providing the function `System` like
the one of the programming language C.

```example.go
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
```

It supports these environments.

- Windows (the shell specfied by %COMSPEC% is used)
- UNIX Like OSes (the shell specfied by $SHELL is used)
