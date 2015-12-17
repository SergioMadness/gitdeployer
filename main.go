package main

import (
	"encoding/json"
	"fmt"
	"gitdeployer/commands"
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
	var result models.Response

	/*
	* All responses in JSON
	 */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	token := r.FormValue("access-token")
	if token == "" || !config.IsTokenExists(token) {
		result = *models.CreateResponse(400, "Need access-token", nil)
	} else {
		switch r.URL.Path {
		case "/gitlab":
			fmt.Println("Gitlab")
			cont := controllers.CreateGitlabController()
			result = cont.WebHook(w, r)
			fmt.Println("Done")
			break
		case "/deploy":
			fmt.Println("Deploy")
			cont := controllers.CreateDeployer()
			cont.Deploy(w, r)
			break
		}
	}

	jsonResult, _ := json.Marshal(result)

	w.Write(jsonResult)
}

func consoleCommand(command string, params []string) {
	switch command {
	case "create-token":
		fmt.Println(config.CreateToken())
		break
	case "deploy":
		for _, serverName := range params {
			fmt.Println("Starting deploy to '" + serverName + "'")
			server := config.GetConfiguration().GetServer(serverName)
			if err := server.Deploy(); err == nil {
				commands.ExecuteCommandList(server.Commands, server.Path)
			}
		}
		fmt.Println("Done")
		break
	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	command := ""

	config.ConfigFilePath = "config.json"
	config.TokenFilePath = "tokens.json"
	config.CommitFilePath = "commits.json"

	configuration := config.GetConfiguration()

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if command != "" {
		consoleCommand(command, os.Args[2:])
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
