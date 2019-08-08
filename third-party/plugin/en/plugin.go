package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/zoncoen-sample/software-design-2019-10/third-party/common"
)

type Greeter struct{}

func (g Greeter) Greet() (string, error) {
	return "Hello!", nil
}

func main() {
	var greeter Greeter
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{Impl: greeter},
		},
	})
}
