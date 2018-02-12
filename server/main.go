package main

import (
	"github.com/ksunhokim123/htmlftp/server/api"
	htmlftp "github.com/ksunhokim123/htmlftp/server/htmlftp"
)

func main() {
	sv := htmlftp.NewService("", "")
	api.Init([]string{}, sv)
	sv.Start()
	api.StartServer()
	for {
	}
}
