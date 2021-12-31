# `gotree`

`gotree` is a library to output tree like string slices in your Go program.

Its use is focused on logging out settings in a 💅 way.

## Usage

```go
package main

import (
    "fmt"
    "strings"

    "github.com/qdm12/gotree"
)

func main() {
    settings := Settings{
        LogLevel: 2,
        Server: ServerSettings{
            Address: ":8000",
            Debug:   true,
        },
    }
    fmt.Println(settings)
}

type Settings struct {
    LogLevel int
    Server   ServerSettings
}

type ServerSettings struct {
    Address string
    Debug   bool
}

func (s Settings) String() string {
    root := gotree.NewRoot("Settings:")
    _ = root.Appendf("Log level: %d", s.LogLevel)
    serverSettings := s.Server.toNode()
    root.AppendNode(serverSettings)
    lines := root.ToLines()
    return strings.Join(lines, "\n")
}

func (s ServerSettings) toNode() *gotree.Node {
    root := gotree.NewRoot("Server settings:")
    root.Appendf("Address: %s", s.Address)
    root.Appendf("Debug: %t", s.Debug)
    return root
}

```

See the [examples](examples) directory for more cases.

## Setup

```sh
go get github.com/qdm12/gotree
```

## Safety to use

- [x] Full unit test coverage
- [x] Full integration test coverage
- [x] Linting with `golangci-lint` with most of its linters enabled
- [ ] In use by the following Go projects:

## Bug and feature request

- [Create an issue](https://github.com/qdm12/gotree/issues/new) or [a discussion](https://github.com/qdm12/gotree/discussions) for feature requests or bugs.
