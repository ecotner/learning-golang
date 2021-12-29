# 01 - Hello

This script is super simple, so I'm not even going to bother explaining it.
Instead, let's take advantage of this simplicity to illustrate some things about your local Go environment.

Assuming you have installed your Go executable correctly, you should see some output when running `which go` and `go version` in the terminal.

## Running from source
The quickest way to run the source code is to simply execute `go run hello.go`.
This will compile to an executable, then run that executable, all in one step.
No intermediate file is created.

## Building from source
You can generate a binary by using `go build hello.go`.
This binary can then be run with `./hello`, and you can delete it with `rm hello`.

## Installing as a package
First, we need to create a `go.mod` file by running `go mod init example/hello`, which keeps track of package dependencies.
You can use `go list` to show a list of packages being used by your project (which should only show `example/hello` for now).
Then you can run `go install hello.go`.
This builds the binary and "installs" it in your `/usr/local/go/bin` directory, where it is typically accessible by anyone on the same machine.
If you have this directory in your `$PATH`, then you should be able to simply execute `hello` to run the program.
Once you grow tired of running it over and over, you can run `go clean -i` to remove the binary and any other associated files, essentially "uninstalling" it.