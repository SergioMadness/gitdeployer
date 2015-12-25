package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitdeployer/commands"
	"gitdeployer/config"
	"gitdeployer/helpers"
	"gitdeployer/models"
	"gitdeployer/modules/logger"
	"net/http"
)

// 'Push' request
const HOOK_PUSH = "Push Hook"

// 'Tag Push' request
const HOOK_PUSH_TAG = "Tag Push Hook"

// 'Merge Request' request
const HOOK_MERGE = "Merge Request Hook"

// Handle request from gitlab
type GitlabController struct {
	GitDeployerBaseController
}

// Constructor
func CreateGitlabController() *GitlabController {
	return new(GitlabController)
}

// Main request handler
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

// 'Push' request handler
func (c *GitlabController) pushHook(gitlabObject models.GitlabRequest) error {
	var result error

	server := config.GetConfiguration().GetServer(gitlabObject.Repository.GitHttpUrl, gitlabObject.Repository.GitSSHUrl)

	if server == nil {
		fmt.Println("No server")
		return errors.New("Need server configuration")
	}

	dir := config.GetConfiguration().ReleaseDir + "/" + helpers.RandomString(8)

	helpers.PrepareDir(dir)

	go commands.TestAndDeploy(server, dir, logger.CreateLogger())

	return result
}

// 'Tag Push' request handler
func (c *GitlabController) tagPushHook(gitlabObject models.GitlabRequest) error {
	return nil
}

// 'Merge Push' request handler
func (c *GitlabController) mergeRequestHook(gitlabObject models.GitlabRequest) error {
	return nil
}
