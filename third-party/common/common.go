package common

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Greeter is an interface for plugin.
type Greeter interface {
	Greet() (string, error)
}

// GreeterRPC is a RPC client for plugin.
type GreeterRPC struct{ client *rpc.Client }

func (g *GreeterRPC) Greet() (string, error) {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// GreeterRPCServer is a RPC server for plugin.
type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	var err error
	*resp, err = s.Impl.Greet()
	return err
}

// HandshakeConfig is a configuration for handshake.
var HandshakeConfig = plugin.HandshakeConfig{
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// GreeterPlugin implements plugin.Plugin interface.
type GreeterPlugin struct {
	Impl Greeter
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}
