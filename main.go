package main

import (
	"encoding/json"
	"fmt"
	"gitdeployer/config"
	"gitdeployer/controllers"
	"gitdeployer/models"
	"log"
	"net/http"
	"os"
	"strconv"
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
	if token == "" || !config.IsTokenExists(token) {
		jsonResult, _ := json.Marshal(models.CreateResponse(400, "Need access-token", nil))

		w.Write(jsonResult)
	} else {
		switch r.URL.Path {
		case "/gitlab":
			fmt.Println("Gitlab")
			cont := controllers.CreateGitlabController()
			cont.WebHook(w, r)
			break
		case "/deploy":
			fmt.Println("Deploy")
			cont := controllers.CreateDeployer()
			cont.Deploy(w, r)
			break
		}
	}
}

func consoleCommand(command string) {
	switch command {
	case "create-token":
		fmt.Println(config.CreateToken())
		break
	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	command := ""

	config.ConfigFilePath = "config.json"
	config.TokenFilePath = "tokens.json"

	configuration := config.GetConfiguration()

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if command != "" {
		consoleCommand(command)
	} else {
		// Main page
		http.HandleFunc("/", handleMessage)
		//Gitlab hook
		http.HandleFunc("/gitlab", handleRequest)
		// Deploy
		http.HandleFunc("/deploy", handleRequest)

		err := http.ListenAndServe(configuration.Host+":"+strconv.Itoa(configuration.Port), nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
