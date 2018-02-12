package main

import (
	htmlftp "github.com/ksunhokim123/htmlftp/server"
	"github.com/ksunhokim123/htmlftp/server/api"
)

func main() {
	sv := htmlftp.NewService("", "")
	api.Init([]string{}, sv)
	sv.Start()
	api.StartServer()
	for {
	}
}
