package config;

import (
	"gopkg.in/yaml.v2"
	"gitdeployer/helpers"
)

type Server struct {
	Name string;
	Path string;
	DefaultBranch string;
	GitUrl string;
	GitLogin string;
	GitPassword string;
}

func CreateServer(name, path, defaultBranch, gitUrl, gitLogin, gitPassword string) *Server {
	result := new(Server);
	
	result.Name = name;
	result.Path = path;
	result.DefaultBranch = defaultBranch;
	result.GitUrl = gitUrl;
	result.GitLogin = gitLogin;
	result.GitPassword = gitPassword;
	
	return result;
}

func RemoveServer(name string) bool {
	result := false;
	
	return result;
}

func IsServerExists(name string) bool {
	result := false;
	
	return result;
}

func saveServers(servers []*Server) bool {
	result := false;
	
	return result;
}

func getServers() []*Server {
	var result []*Server;
	
	if helpers.IsFileExists(currentConfig.ServersFilePath) {
		yaml.Unmarshal([]byte(currentConfig.TokenFilePath), result);
	}
	
	return result;
}