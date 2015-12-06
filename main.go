package main

import (
	"gitdeployer/config"
	"gitdeployer/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
* Handle http request
 */
func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is git deployer."))

	w.Header().Set("Access-Control-Allow-Origin", "*")
}

/**
* Chat response handler
 */
func handleRequest(w http.ResponseWriter, r *http.Request) {
	/*
	* All responses in JSON
	 */
	w.Header().Set("Content-Type", "application/json")

	token := r.FormValue("access-token")
	if token != "" {
		config.GetSession().SetCurrentUser(models.NewProfile(config.GetConnection()).GetByToken(token))
	}

	switch r.URL.Path {
	case "/deploy":
		fmt.Println("Deploy")
		cont := controllers.CreateDeployer()
		cont.Deploy(w, r)
		break
	}
}

func consoleCommand(command string) {
	switch command {
	case "create-token":
		config.CreateToken()
		break
	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	command := ""

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if command != "" {
		consoleCommand(command)
	} else {
		// Main page
		http.HandleFunc("/", handleMessage)
		// Registration
		http.HandleFunc("/deploy", handleRequest)

		err := http.ListenAndServe(":81", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
