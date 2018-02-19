package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sunho/mouse-hosting/server/api"
	mouse "github.com/sunho/mouse-hosting/server/mouse"
)

func main() {
	sv := mouse.NewService("config.yaml")
	sv.Start()

	api.Init([]string{}, sv)
	api.StartServer()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	sv.Stop()
}
