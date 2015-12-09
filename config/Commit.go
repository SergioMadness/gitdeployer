package config

import (
	"encoding/json"
	"fmt"
	"gitdeployer/helpers"
	"io/ioutil"
	"time"
)

type Commit struct {
	DeployDate int64
	CommitHash string
}

func AddCommit(hash string) *Commit {
	result := CreateCommit(hash)

	commits := getCommits()
	commits = append(commits, result)
	saveCommits(commits)

	return result
}

func CreateCommit(hash string) *Commit {
	result := new(Commit)

	result.DeployDate = time.Now().Unix()
	result.CommitHash = hash

	return result
}

func IsCommitDeployed(hash string) bool {
	result := false

	commits := getCommits()
	for _, commit := range commits {
		if commit.CommitHash == hash {
			result = true
		}
	}

	return result
}

func saveCommits(commits []*Commit) bool {
	result := false

	if ts, err := json.Marshal(commits); err == nil {
		ioutil.WriteFile(CommitFilePath, ts, 0777)
		result = true
	}

	return result
}

func getCommits() []*Commit {
	var result []*Commit

	if helpers.IsPathExists(CommitFilePath) {
		if tokensStr, err := ioutil.ReadFile(CommitFilePath); err == nil {
			if err := json.Unmarshal(tokensStr, &result); err != nil {
				fmt.Println(err)
			}
		}
	}

	return result
}
