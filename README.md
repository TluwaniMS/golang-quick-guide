# Golang

## What is golang?

Developed by Google, Go, or Golang, stands as an open-source, compiled, and statically typed programming language. Its design focuses on simplicity, high performance, readability, and efficiency.

[Install go](https://go.dev/)

## Basic go commands.

* ### go get command:

In Go, the go get command is primarily used to fetch and install packages from remote repositories into your Go workspace.

`go get <flags> <packages>`

`go get [-t] [-u] [-v] [build flags] [packages]`

* ### go mod init command:

The go mod init command is used in Go (also known as Golang) to initialize a new module, creating a go.mod file within the directory where the command is executed. This command is primarily used when starting a new Go project or when converting an existing Go codebase into a module.

When you run go mod init, it initializes the module by creating a go.mod file that keeps track of dependencies and version information for the project. This file contains metadata about the module, such as its module path and required dependencies. The module path usually refers to the URL where the module's code is hosted (like a Git repository), but it can also be any valid module path you define.

`go mod init [module-path]`

* ### go mod tidy command:

Tidy makes sure go.mod matches the source code in the module.
It adds any missing modules necessary to build the current module's
packages and dependencies, and it removes unused modules that
don't provide any relevant packages. It also adds any missing entries
to go.sum and removes any unnecessary ones.

`go mod tidy [-e] [-v] [-go=version] [-compat=version]`

* ### go build command:

The go build command is used to compile Go source code files into executable binaries or libraries. When you run go build, it compiles the Go code in the current directory (or specific files if specified) and generates an executable file if the code is intended to be an application, or it creates a package if the code is a library.

`go build [-o output] [build flags] [packages]`

