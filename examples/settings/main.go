package main

import (
	"fmt"

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
	return s.toNode().String()
}

func (s Settings) toNode() *gotree.Node {
	node := gotree.New("Settings:")
	node.Appendf("Log level: %d", s.LogLevel)
	node.AppendNode(s.Server.toNode())
	return node
}

func (s ServerSettings) toNode() *gotree.Node {
	node := gotree.New("Server settings:")
	node.Appendf("Address: %s", s.Address)
	node.Appendf("Debug: %t", s.Debug)
	return node
}
