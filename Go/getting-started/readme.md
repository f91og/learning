## go.mod
- manage dependency tracking
- go.mod file stays with your code, including in your source code repository
- To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in
- The name is the module's module path. In most cases, this will be the repository location where your source code will be kept, such as github.com/mymodule

## go cmd
- go mod init mouduleName
- go run fileName
- go help
- go mod tidy: add missing hashes and will remove unnecessary hashes from go.sum. 执行这个命令会自动在go.mod中添加go代码import的包
- go mod edit: provides a command-line interface for editing and formatting go.mod files
- go test
- go build: compiles the packages, along with their dependencies, but it doesn't install the results
- go install: compiles and installs the packages
- go list

## Call code in an external package 
- Visit pkg.go.dev and serach wanted package
- In the Documentation section, under Index, note the list of functions you can call from your code
- At the top of this page, note that package quote is included in the rsc.io/quote module
- directly import the module(such as rsc.io/quote) in Go code file to use those functions

## Add new module requirements and sums
- Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module

