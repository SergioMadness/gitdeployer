package config

import (
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
	"strings"
	"gitdeployer/modules/logger"
)

type Configuration struct {
	Host    string
	Port    int
	Servers []*Server
}

var ConfigFilePath, TokenFilePath, CommitFilePath string
var currentConfig = new(Configuration)

func (conf *Configuration) GetServer(params ...string) *Server {
	var result *Server

	serverIndex := conf.FindServer(params...)
	if serverIndex >= 0 {
		result = conf.Servers[serverIndex]
	}

	return result
}

func (conf *Configuration) FindServer(params ...string) int {
	result := -1

	for _, serv := range conf.Servers {
		for index, param := range params {
			if strings.Contains(serv.Name, param) || strings.Contains(serv.Path, param) || strings.Contains(serv.GitUrl, param) {
				result = index
				break
			}
		}
	}

	return result
}

func (conf *Configuration) RemoveServer(params ...string) bool {
	result := false

	serverIndex := conf.FindServer(params...)
	if serverIndex >= 0 {
		conf.Servers = append(conf.Servers[:serverIndex], conf.Servers[serverIndex+1:]...)
	}

	return result
}

func (conf *Configuration) IsServerExists(param string) bool {
	return conf.FindServer(param) >= 0
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

func (conf *Configuration) GetLogger() *logger.Logger {
	return currentLoger
}
