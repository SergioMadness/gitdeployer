package config

import (
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
)

type Configuration struct {
	Host    string
	Port    int
	Servers []Server
}

var ConfigFilePath, TokenFilePath string
var currentConfig = new(Configuration)

func (conf *Configuration) GetServer(param string) *Server {
	var result *Server

	for _, serv := range conf.Servers {
		if serv.Name == param || serv.Path == param || serv.GitUrl == param {
			result = &serv
			break
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
