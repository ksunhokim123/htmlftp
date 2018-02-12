package main

import (
	"github.com/ksunhokim123/mouse-hosting/server/api"
	mouse "github.com/ksunhokim123/mouse-hosting/server/mouse"
)

func main() {
	sv := mouse.NewService("", "")
	api.Init([]string{}, sv)
	sv.Start()
	api.StartServer()
	for {
	}
}
