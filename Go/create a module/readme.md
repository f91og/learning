## Developing and publishing modules
- Go code is grouped into packages, and packages are grouped into modules
- Start your module using the go mod init command. eg: go mod init example.com/greetings
- call local moudle from another module
- go mod edit, redirect Go tools from its module path (where the module isn't) to the local directory (where it is)
```go
$ go mod edit -replace=example.com/greetings=../greetings
```

For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system

## go package
- fmt
- errors
- "math/rand"
- time
- regexp
- testing
- log
