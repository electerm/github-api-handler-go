package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/koding/multiconfig"
)

// Server is config interface
type Server struct {
	Port int `default:"7654"`
}

//Handler is request Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", Handler)

	defaultConf := multiconfig.NewWithPath("config.default.toml")
	localConf := multiconfig.NewWithPath("config.toml") // supports TOML and JSON
	// Get an empty struct for your configuration
	serverDefaultConf := new(Server)
	serverUserConf := new(Server)

	// Populated the serverConf struct
	defaultConf.MustLoad(serverDefaultConf)
	err := localConf.Load(serverUserConf) // Check for error
	//m.MustLoad(serverConf) // Panic's if there is any error
	// Bind to a port and pass our router in
	if err != nil {
		log.Print("no config.toml, but it is ok")
	} else {
		serverDefaultConf.Port = serverUserConf.Port
	}

	addr := "localhost:" + strconv.Itoa(serverDefaultConf.Port)
	log.Print("server runs on ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
