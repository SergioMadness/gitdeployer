package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitdeployer/commands"
	"gitdeployer/config"
	"gitdeployer/models"
	"net/http"
)

const HOOK_PUSH = "Push Hook"
const HOOK_PUSH_TAG = "Tag Push Hook"
const HOOK_MERGE = "Merge Request Hook"

type GitlabController struct {
	GitDeployerBaseController
}

func CreateGitlabController() *GitlabController {
	return new(GitlabController)
}

func (c *GitlabController) WebHook(w http.ResponseWriter, r *http.Request) models.Response {
	var result models.Response
	var gitRequest models.GitlabRequest

	eventType := r.Header.Get("X-Gitlab-Event")

	fmt.Println(eventType)

	fmt.Print("Body: ")
	fmt.Println(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&gitRequest); err == nil {
		fmt.Print("Decoded request: ")
		fmt.Println(gitRequest)
		var hookErr error
		var commit string

		if len(gitRequest.Commits) > 0 {
			commit = gitRequest.Commits[len(gitRequest.Commits)-1].Id
		}
		if !config.IsCommitDeployed(commit) {
			config.AddCommit(commit)

			switch eventType {
			case HOOK_PUSH:
				hookErr = c.pushHook(gitRequest)
				break
			case HOOK_PUSH_TAG:
				hookErr = c.tagPushHook(gitRequest)
				break
			case HOOK_MERGE:
				hookErr = c.mergeRequestHook(gitRequest)
				break
			}

			if hookErr != nil {
				result.Result = 500
				result.ResultMessage = "Deploy failed"
				fmt.Println("Failed")
			}
		} else {
			fmt.Println("Commit is already deployed")
			result.Result = 403
			result.ResultMessage = "Commit is already deployed"
		}
	} else {
		fmt.Print("Error: ")
		fmt.Println(err)
		result.Result = 400
		result.ResultMessage = "Can't parse request"
	}

	return result
}

func (c *GitlabController) pushHook(gitlabObject models.GitlabRequest) error {
	var result error

	server := config.GetConfiguration().GetServer(gitlabObject.Repository.GitHttpUrl, gitlabObject.Repository.GitSSHUrl)

	if server == nil {
		fmt.Println("No server")
		return errors.New("Need server configuration")
	}

	if result = c.PrepareServer(*server); result != nil {
		return errors.New("Can't prepare server")
	}
	fmt.Println("Deployed")

	if output, err := commands.ExecuteCommandList(server.Commands, server.Path); err == nil {
		fmt.Println(output)
//		logger := config.GetConfiguration().GetLogger()
//		logger.Log("full", output)
//		logger.Flush()
	}

	return result
}

func (c *GitlabController) tagPushHook(gitlabObject models.GitlabRequest) error {
	return nil
}

func (c *GitlabController) mergeRequestHook(gitlabObject models.GitlabRequest) error {
	return nil
}
