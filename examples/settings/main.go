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
	root := gotree.New("Settings:")
	_ = root.Appendf("Log level: %d", s.LogLevel)
	serverSettings := s.Server.toNode()
	root.AppendNode(serverSettings)
	lines := root.ToLines()
	return strings.Join(lines, "\n")
}

func (s ServerSettings) toNode() *gotree.Node {
	root := gotree.New("Server settings:")
	root.Appendf("Address: %s", s.Address)
	root.Appendf("Debug: %t", s.Debug)
	return root
}
