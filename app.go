package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/koding/multiconfig"
)

// Server is config interface
type Server struct {
	Port                  int
	GithubReleaseJSONPath string
	APIName               string
}

var serverDefaultConf = new(Server)

// Handler is request Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("error")
	}
	str0 := string(body)
	if strings.Contains(str0, "\"published\"") {
		ioutil.WriteFile(serverDefaultConf.GithubReleaseJSONPath, body, 0777)
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte("no"))
	}
}

func main() {

	defaultConf := multiconfig.NewWithPath("config.default.toml")
	localConf := multiconfig.NewWithPath("config.toml") // supports TOML and JSON
	// Get an empty struct for your configuration

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
		serverDefaultConf.APIName = serverUserConf.APIName
		serverDefaultConf.GithubReleaseJSONPath = serverUserConf.GithubReleaseJSONPath
	}

	r := mux.NewRouter()
	url := "/github-electerm-api/" + serverDefaultConf.APIName
	r.HandleFunc(url, Handler).
		Methods("POST")

	addr := "localhost:" + strconv.Itoa(serverDefaultConf.Port)
	log.Print("server runs on ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
