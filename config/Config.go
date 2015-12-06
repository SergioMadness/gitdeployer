package config;

import (
	"time"
//	"gopkg.in/yaml.v2"
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
)

type Configuration struct {
	TokenFilePath string;
}

type Token struct {
	CreateDate int64;
	Hash string;
}

func CreateToken(token string) *Token {
	result:=new(Token);
	
	result.Hash = token;
	result.CreateDate = time.Now().Unix();
	
	return result;
}

var currentConfig Configuration;

func GetConfiguration() Configuration {
	return currentConfig;
}


func AddToken() string {
	result:=helpers.RandomString(24);
	
	tokens:=getTokens();
	tokens=append(tokens, CreateToken(result));
	saveTokens(tokens);
	
	return result;
}

func RemoveToken(token string) bool {
	result:=false;
	
	tokens:=getTokens();
	for index, cToken:=range tokens {
		if cToken.Hash == token {
			tokens = append(tokens[:index], tokens[index+1:]...);
			result = true;
		}
	}
	if result {
		saveTokens(tokens);
	}
	
	return result;
}

func IsTokenExists(token string) bool {
	result:=false;
	
	tokens:=getTokens();
	for _, cToken:=range tokens {
		if cToken.Hash == token {
			result = true;
		}
	}
	
	return result;
}

func saveTokens(tokens []*Token) bool {
	result:=false;
	
	if ts, err := json.Marshal(tokens); err==nil {
		ioutil.WriteFile(currentConfig.TokenFilePath, ts, 0777);
		result = true;
	}
	
	return result;
}

func getTokens() []*Token {
	var result []*Token;
	
	if helpers.IsFileExists(currentConfig.TokenFilePath) {
		json.Unmarshal([]byte(currentConfig.TokenFilePath), result);
	}
	
	return result;
}