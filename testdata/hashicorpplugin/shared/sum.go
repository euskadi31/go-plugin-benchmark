package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var PluginMap = map[string]plugin.Plugin{
	"sum": &SumPlugin{},
}

// Sum is the interface that we're exposing as a plugin.
type Sum interface {
	Sum(a int32, b int32) int32
}

// Here is an implementation that talks over RPC
type SumRPC struct {
	client *rpc.Client
}

func (g *SumRPC) Sum(a int32, b int32) int32 {
	var resp int32
	err := g.client.Call("Plugin.Sum", map[string]int32{
		"a": a,
		"b": b,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that SumRPC talks to, conforming to
// the requirements of net/rpc
type SumRPCServer struct {
	// This is the real implementation
	Impl Sum
}

func (s *SumRPCServer) Sum(args map[string]int32, resp *int32) error {
	*resp = s.Impl.Sum(args["a"], args["b"])

	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a SumRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return SumRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type SumPlugin struct {
	// Impl Injection
	Impl Sum
}

func (p *SumPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &SumRPCServer{
		Impl: p.Impl,
	}, nil
}

func (SumPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &SumRPC{
		client: c,
	}, nil
}
