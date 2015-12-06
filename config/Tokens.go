package config;

import (
	"time"
	"encoding/json"
	"gitdeployer/helpers"
	"io/ioutil"
)

type Token struct {
	CreateDate int64;
	Hash string;
}

func AddToken(token string) *Token {
	result:=new(Token);
	
	result.Hash = token;
	result.CreateDate = time.Now().Unix();
	
	return result;
}

func CreateToken() string {
	result:=helpers.RandomString(24);
	
	tokens:=getTokens();
	tokens=append(tokens, AddToken(result));
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
		ioutil.WriteFile(TokenFilePath, ts, 0777);
		result = true;
	}
	
	return result;
}

func getTokens() []*Token {
	var result []*Token;
	
	if helpers.IsFileExists(TokenFilePath) {
		json.Unmarshal([]byte(TokenFilePath), result);
	}
	
	return result;
}