package main

import (
	"github.com/euskadi31/go-plugin-benchmark/testdata/hashicorpplugin/shared"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type Plugin struct{}

func (Plugin) Sum(a int32, b int32) int32 {
	return a + b
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"sum": &shared.SumPlugin{
				Impl: &Plugin{},
			},
		},
	})
}
