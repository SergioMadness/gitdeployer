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

func GetConfiguration() *Configuration {
	if helpers.IsFileExists(ConfigFilePath) {
		if confFile, err := ioutil.ReadFile(ConfigFilePath); err == nil {
			json.Unmarshal(confFile, &currentConfig)
		}
	}

	return currentConfig
}

func SaveConfiguration() bool {
	var result bool

	if helpers.IsFileExists(ConfigFilePath) {
		if confYaml, err := json.Marshal(currentConfig); err == nil {
			ioutil.WriteFile(ConfigFilePath, []byte(confYaml), 0777)
		}
	}

	return result
}
