package main

import (
	handler "github.com/da-moon/coe817-dare/cmd/decryptor-plugin/handler"
	grpc "github.com/da-moon/coe817-dare/plugins/decryptor/grpc"
	shared "github.com/da-moon/coe817-dare/plugins/shared"
	plugin "github.com/hashicorp/go-plugin"
)

// ServeConfig - This is the plugin config thet is used in main function of engine
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"decryptor": &grpc.Plugin{Impl: &handler.Decrypt{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
