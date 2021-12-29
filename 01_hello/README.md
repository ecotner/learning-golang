# 01 - Hello

## Step 1
This script is super simple, so I'm not even going to bother explaining it.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```
Instead, let's take advantage of this simplicity to illustrate some things about your local Go environment.

Assuming you have installed your Go executable correctly, you should see some output when running `which go` and `go version` in the terminal.

### Running from source
The quickest way to run the source code is to simply execute `go run hello.go`.
This will compile to an executable, then run that executable, all in one step.
No intermediate file is created.

### Building from source
You can generate a binary by using `go build hello.go`.
This binary can then be run with `./hello`, and you can delete it with `rm hello`.

### Installing as a package
First, we need to create a `go.mod` file by running `go mod init example/hello`, which keeps track of package dependencies.
You can use `go list` to show a list of packages being used by your project (which should only show `example/hello` for now).
Then you can run `go install hello.go`.
This builds the binary and "installs" it in your `/usr/local/go/bin` directory, where it is typically accessible by anyone on the same machine.
If you have this directory in your `$PATH`, then you should be able to simply execute `hello` to run the program.
Once you grow tired of running it over and over, you can run `go clean -i` to remove the binary and any other associated files, essentially "uninstalling" it.

## Step 2
Now we're going to add an external package dependency.
We add `"rsc.io/quote"` to our `import` statement, then change the print call to `fmt.Println(quote.Go())`, which calls the `Go()` function from the `quote` package.
At first, Go will not know that this package even exists and will throw some fits if you try to run things as-is.
If you run `go mod tidy`, the package manager will automatically add the dependency to your `go.mod` file, search for the `rsc.io/quote` package (and any of its dependencies) online at [pkg.go.dev](https://pkg.go.dev), download it (to `$GOPATH/pkg/mod/rsc.io/quote@vX.X.X`), generate an MD5 hash for each dependency installed, and put those in a `go.sum` file.
Then you can run/build/install from source just as in step 1.

At first blush, uninstalling works the same way by running `go clean -i`, only on closer inspection, the dependency source files are not removed.
You can run `go clean -i -modcache`, but it uninstalls ALL downloaded packages, including ones unrelated to the current project...
The best solution seems to be to run `go get rsc.io/quote@none` to uninstall this one specific package.
And if you remove the package from your source altogether, `go mod tidy` should auto-uninstall it for you.

