package pluginbench

import (
	"os/exec"
	"testing"

	"github.com/euskadi31/go-plugin-benchmark/testdata/hashicorpplugin/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/stretchr/testify/assert"
)

func TestHashicorpPlugin(t *testing.T) {
	logger := hclog.NewNullLogger()

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("./testdata/hashicorpplugin/plugin/sum.hashiplug"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		t.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("sum")
	if err != nil {
		t.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	sum := raw.(shared.Sum)

	assert.Equal(t, int32(42), sum.Sum(5, 37))
}

func BenchmarkHashicorpPlugin(b *testing.B) {
	logger := hclog.NewNullLogger()

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("./testdata/hashicorpplugin/plugin/sum.hashiplug"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		b.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("sum")
	if err != nil {
		b.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	sum := raw.(shared.Sum)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum.Sum(5, 37)
	}
}
