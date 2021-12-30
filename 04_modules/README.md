# Modules

In this repository, there are two modules, `greetings` and `hello`.
`greetings` is exposed as a package by the `package greetings` statement in `greetings.go`, and initialized with `go mod init example.com/greetings`.
It provides a basic function `Hello` that is exposed to be used by whoever imports the package.
(Note that in order to export functions, the first letter of the function name _must_ be capitalized!)

Then in the `hello` module, we import `example.com/greetings`.
We must also point `hello`'s dependencies in the right place to find the `greetings` package, so we run `go mod edit -replace example.com/greetings=../greetings` and `go mod tidy` which fixes up the `go.mod` file.
Executing `go run .` should compile and run the code in `hello.go`.