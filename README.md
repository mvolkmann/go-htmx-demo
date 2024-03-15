# go-htmx-demo

See [TEMPL](https://templ.guide/).

Verify that the `GOPATH` environment variable is set to a value like `$HOME/go`.

Verify that `$GOPATH/bin` is in your `PATH` environment variable.

Here are the commands I ran to create this project:

- `brew install go`
- create the files `main.go` and `views/hello.templ`
- `go mod init views`
- `go get github.com/a-h/templ`
- `go install github.com/a-h/templ/cmd/templ@latest`
- `go mod init github.com/a-h/templ-examples/hello-world`
- `go mod tidy`
- `go install golang.org/x/lint/golint`

  This installs the command "golint".
  To lint a file, enter `golint {name}.go`.

- `templ generate --watch`

  This creates a `.go` file for every `.templ` file
  in and below the current directory.
  It watches for changes and creates new `.go` files as needed.

- `go get github.com/labstack/echo/v4`
- `go get github.com/labstack/echo/v4/middleware`

  This is a popular HTTP server library.

- `go install github.com/cosmtrek/air@latest`
- `go get -u github.com/cosmtrek/air`

  This provides live reload of Go servers when code changes are detected.

- `air`

  This starts a server that listens on port 3000.

- browse localhost:3000
