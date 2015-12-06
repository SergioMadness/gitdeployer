package config;

type Configuration struct {
	TokenFilePath string;
	ServersFilePath string;
}

var currentConfig Configuration;

func GetConfiguration() Configuration {
	return currentConfig;
}