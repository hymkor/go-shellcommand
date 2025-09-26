go-shellcommand
===============

go-shellcommand is the package providing the function `System` like
the one of the programming language C.

```go
process, err := shellcommand.System(fmt.Sprintf(`echo "ahaha" > "%s"`, workpath))
if err != nil {
    t.Fatal(err.Error())
    return
}
process.Wait()
```

It supports these environments.

- Windows (the shell specfied by %COMSPEC% is used)
- UNIX Like OSes (the shell specfied by $SHELL is used)
