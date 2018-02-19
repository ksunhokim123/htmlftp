package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	mouse "github.com/sunho/mouse-hosting/server/mouse"
)

var Service *mouse.Service
var Routes *mux.Router

func Init(allowedDomains []string, service *mouse.Service) {
	Service = service
	Routes = mux.NewRouter()

	Routes.HandleFunc("/api/users", AddUser).Methods("POST")
	Routes.HandleFunc("/api/keys/{key:[a-z-]+}", GetKey).Methods("GET")

	Routes.HandleFunc("/api/users", useBasicAuth(RetrieveUsers)).Methods("GET")
	Routes.HandleFunc("/api/users/{id:[a-z0-9]+}", useBasicAuth(GetUser)).Methods("GET")
	Routes.HandleFunc("/api/users/{id:[a-z0-9]+}", useBasicAuth(RemoveUser)).Methods("DELETE")
	Routes.HandleFunc("/api/keys", useBasicAuth(RetrieveKeys)).Methods("GET")
	Routes.HandleFunc("/api/keys/{key:[a-z-]+}", useBasicAuth(RemoveKey)).Methods("DELETE")
	Routes.HandleFunc("/api/keygen", useBasicAuth(GenerateKey)).Methods("GET")
}

func useBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return use(h, basicAuth)
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

		if username != Service.Config.Username || password != Service.Config.Password {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func StartServer() {
	address := Service.Config.Address.String()
	fmt.Println("running:", address)
	go func() {
		if err := http.ListenAndServe(address, Routes); err != nil {
			fmt.Printf("API server http error: %v", err)
		}
	}()
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	key := r.FormValue("key")
	response := struct {
		Error string `json:"error"`
	}{
		Error: "success",
	}

	if len(username) == 0 {
		response.Error = "empty username"
	} else if len(password) == 0 {
		response.Error = "empty password"
	} else if len(key) == 0 {
		response.Error = "empty key"
	} else if keyindex := Service.KeyContainer.Exist(key); keyindex == -1 {
		response.Error = "no such key"
	} else if err := Service.UserContainer.AddUser(username, password); err != nil {
		response.Error = err.Error()
	} else {
		Service.KeyContainer.Remove(keyindex)
	}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//TODO improve error handling
func GetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := struct {
		Error string `json:"error"`
		Exist bool   `json:"exist"`
	}{
		Error: "success",
		Exist: false,
	}
	key := vars["key"]
	if Service.KeyContainer.Exist(key) != -1 {
		response.Exist = true
	}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func RetrieveKeys(w http.ResponseWriter, r *http.Request) {
	keyList := []string{}
	for _, key := range *(Service.KeyContainer) {
		keyList = append(keyList, key)
	}
	fmt.Println(keyList)
	response := struct {
		Error   string   `json:"error"`
		KeyList []string `json:"keys"`
	}{
		Error:   "success",
		KeyList: keyList,
	}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GenerateKey(w http.ResponseWriter, r *http.Request) {
	key := Service.KeyContainer.Generate()
	response := struct {
		Error string `json:"error"`
		Key   string `json:"key"`
	}{
		Error: "success",
		Key:   key,
	}
	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func RetrieveUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}

func RemoveKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}
