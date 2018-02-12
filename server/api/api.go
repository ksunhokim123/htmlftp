package api

import (
	"net/http"

	"github.com/gorilla/mux"
	mouse "github.com/ksunhokim123/mouse-hosting/server/mouse"
)

var Service *mouse.Service
var Routes *mux.Router

func Init(allowedDomains []string, service *mouse.Service) {
	Service = service
	Routes = mux.NewRouter()

	Routes.HandleFunc("/api/users", AddUser).Methods("POST")
	Routes.HandleFunc("/api/keys/{key:[a-z0-9-]+}", GetKey).Methods("GET")

	Routes.HandleFunc("/api/users", use(RetrieveUsers, basicAuth)).Methods("GET")
	Routes.HandleFunc("/api/users/{id:[a-z0-9]+}", use(GetUser, basicAuth)).Methods("GET")
	Routes.HandleFunc("/api/users/{id:[a-z0-9]+}", use(RemoveUser, basicAuth)).Methods("DELETE")
	Routes.HandleFunc("/api/keys", use(RetrieveKeys, basicAuth)).Methods("GET")
	Routes.HandleFunc("/api/keys/{key:[a-z0-9-]+}", use(RemoveKey, basicAuth)).Methods("DELETE")
	Routes.HandleFunc("/api/keygen", use(GenerateKey, basicAuth)).Methods("GET")
	//TODO options
}
func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != Service.Config.UserName || password != Service.Config.Password {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func StartServer() {
	http.ListenAndServe(":80", Routes)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

}

func RetrieveUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func GetKey(w http.ResponseWriter, r *http.Request) {

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {

}

func GenerateKey(w http.ResponseWriter, r *http.Request) {

}

func RetrieveKeys(w http.ResponseWriter, r *http.Request) {

}

func RemoveKey(w http.ResponseWriter, r *http.Request) {

}
