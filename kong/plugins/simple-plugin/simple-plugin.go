package main

import (
    "fmt"
    "github.com/Kong/go-pdk"
    "github.com/Kong/go-pdk/server"
)

type Config struct {
    Apikey string
    Message string
}

var Version = "0.2"
var Priority = 1

func New() interface{} {
    return &Config{}
}

func main() {
  server.StartServer(New, Version, Priority)
}

func (conf Config) Access(kong *pdk.PDK) {
    key, err := kong.Request.GetQueryArg("key")
    apiKey := conf.Apikey

    if err != nil {
        kong.Log.Err(err.Error())
    }

    x := make(map[string][]string)
    x["Content-Type"] = append(x["Content-Type"], "application/json")

    if apiKey != key {
        kong.Response.Exit(403, "Youu have no correct key", x)
    }

    kong.Log.Info("=== ACCESS BLOCK ===")
    kong.Log.Info(kong)
    message := conf.Message
    if len(message) == 0 {
        message = "Hello World"
    }
    kong.Response.SetHeader("x-hello-from-go", fmt.Sprintf("Go says %s", message))
    kong.Log.Info("/=== ACCESS BLOCK ===/")

}

func (conf Config) Rewrite(kong *pdk.PDK) {
    kong.Log.Info("=== REWRITE BLOCK ===")
    kong.Log.Info(kong)
    kong.Log.Info("/=== REWRITE BLOCK ===/")
}


func (conf Config) Log(kong *pdk.PDK) {
    kong.Log.Info("=== kong.Log BLOCK ===")
    kong.Log.Info(kong)
    kong.Log.Info("/=== kong.Log BLOCK ===/")
}
