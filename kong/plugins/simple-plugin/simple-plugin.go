package main

import (
	"fmt"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

var Version = "0.2"
var Priority = 1

type Config struct {
	Message string
}

func main() {
	server.StartServer(New, Version, Priority)
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Certificate(kong *pdk.PDK) {
	kong.Log.Info("=== CERTIFICATE BLOCK ===")
	kong.Log.Info(kong)
	kong.Log.Info("/=== CERTIFICATE BLOCK ===/")
}

func (conf Config) Rewrite(kong *pdk.PDK) {
	kong.Log.Info("=== REWRITE BLOCK ===")
	kong.Log.Info(kong)
	kong.Log.Info("/=== REWRITE BLOCK ===/")
}

func (conf Config) Access(kong *pdk.PDK) {
	kong.Log.Info("=== ACCESS BLOCK ===")
	kong.Log.Info(kong)
	message := conf.Message
	if message == "" {
		message = "Hello World"
	}
	kong.Response.SetHeader("x-hello-from-go", fmt.Sprintf("Go says %s", message))
	kong.Log.Info("/=== ACCESS BLOCK ===/")
}

func (conf Config) Preread(kong *pdk.PDK) {
	kong.Log.Info("=== PREREAD BLOCK ===")
	kong.Log.Info(kong)
	kong.Log.Info("/=== PREREAD BLOCK ===/")
}

func (conf Config) Log(kong *pdk.PDK) {
	kong.Log.Info("=== kong.Log BLOCK ===")
	kong.Log.Info(kong)
	kong.Log.Info("/=== kong.Log BLOCK ===/")
}
