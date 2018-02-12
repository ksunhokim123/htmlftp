package api

import (
	"net/http"

	htmlftp "github.com/ksunhokim123/htmlftp/server/htmlftp"
)

var Service *htmlftp.Service

func Init(allowedDomains []string, service *htmlftp.Service) {
	Service = service
}

func StartServer() {

}

func AddUser(w http.ResponseWriter, r *http.Request) {

}

func RetrieveUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

}

func GenerateKey(w http.ResponseWriter, r *http.Request) {

}

func RetrieveKeys(w http.ResponseWriter, r *http.Request) {

}

func RemoveKey(w http.ResponseWriter, r *http.Request) {

}
