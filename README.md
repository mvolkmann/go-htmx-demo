# go-htmx-demo

See [TEMPL](https://templ.guide/).

Verify that the `GOPATH` environment variable is set to a value like `$HOME/go`.

Verify that `$GOPATH/bin` is in your `PATH` environment variable.

Here are the commands I ran to create this project:

- `brew install go`
- `go install github.com/a-h/templ/cmd/templ@latest`
- `go mod init github.com/a-h/templ-examples/hello-world`
- `go get github.com/a-h/templ`
- create the files `hello.templ` and `main.go`
- `templ generate`

  This creates the file `hello_templ.go`.

- `go run .`

  This outputs "<div>Hello, John</div>".
