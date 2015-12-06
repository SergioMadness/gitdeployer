package config;

type Configuration struct {
	TokenFilePath string;
}

var currentConfig Configuration;

func GetConfiguration() Configuration {
	return currentConfig;
}