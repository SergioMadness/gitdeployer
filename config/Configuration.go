package config;

import(
	"gitdeployer/helpers"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Servers []*Server
}

var ConfigFilePath, TokenFilePath string;
var currentConfig Configuration;

func GetConfiguration() Configuration {
	var result Configuration;
	
	if helpers.IsFileExists(ConfigFilePath) {
		yaml.Unmarshal([]byte(ConfigFilePath), result);
	}
	
	return result;
}

func SaveConfiguration() bool {
	var result bool;
	
	if helpers.IsFileExists(ConfigFilePath) {
		if confYaml, err:=yaml.Marshal(currentConfig.Servers); err==nil {
			ioutil.WriteFile(ConfigFilePath, []byte(confYaml), 0777);
		}
	}
	
	return result;
}