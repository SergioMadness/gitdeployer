package config

import (
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
	"strings"
)

type Configuration struct {
	Host    string
	Port    int
	Servers []Server
}

var ConfigFilePath, TokenFilePath, CommitFilePath string
var currentConfig = new(Configuration)

func (conf *Configuration) GetServer(params ...string) *Server {
	var result *Server

	for _, serv := range conf.Servers {
		for _, param := range params {
			if strings.Contains(serv.Name, param) || strings.Contains(serv.Path, param) || strings.Contains(serv.GitUrl, param) {
				result = &serv
				break
			}
		}
	}

	return result
}

func (conf *Configuration) RemoveServer(param string) bool {
	result := false

	return result
}

func (conf *Configuration) IsServerExists(param string) bool {
	result := false

	return result
}

func GetConfiguration() *Configuration {
	if helpers.IsPathExists(ConfigFilePath) {
		if confFile, err := ioutil.ReadFile(ConfigFilePath); err == nil {
			json.Unmarshal(confFile, &currentConfig)
		}
	}

	return currentConfig
}

func SaveConfiguration() bool {
	var result bool

	if helpers.IsPathExists(ConfigFilePath) {
		if confYaml, err := json.Marshal(currentConfig); err == nil {
			ioutil.WriteFile(ConfigFilePath, []byte(confYaml), 0777)
		}
	}

	return result
}
