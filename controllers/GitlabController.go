package controllers

import (
	"encoding/json"
	"errors"
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

	if err := json.NewDecoder(r.Body).Decode(&gitRequest); err == nil {
		var hookErr error

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
		}
	} else {
		result.Result = 400
		result.ResultMessage = "Can't parse request"
	}

	return result
}

func (c *GitlabController) pushHook(gitlabObject models.GitlabRequest) error {
	var result error

	server := config.GetConfiguration().GetServer(gitlabObject.Repository.GitHttpUrl)

	if server == nil {
		return errors.New("Need server configuration")
	}

	if result = c.PrepareServer(*server); result == nil {
		
	}

	return result
}

func (c *GitlabController) tagPushHook(gitlabObject models.GitlabRequest) error {
	return nil
}

func (c *GitlabController) mergeRequestHook(gitlabObject models.GitlabRequest) error {
	return nil
}
