package controllers

import (
	"encoding/json"
	"gitdeployer/models"
	"net/http"
)

const HOOK_PUSH = "Push Hook"
const HOOK_PUSH_TAG = "Tag Push Hook"
const HOOK_MERGE = "Merge Request Hook"

type GitlabController struct {
}

func CreateGitlabController() *GitlabController {
	return new(GitlabController)
}

func (c *GitlabController) WebHook(w http.ResponseWriter, r *http.Request) models.Response {
	var result models.Response
	var gitRequest models.GitlabRequest

	eventType := r.Header.Get("X-Gitlab-Event")

	if err := json.NewDecoder(r.Body).Decode(&gitRequest); err == nil {
		switch eventType {
		case HOOK_PUSH:
			c.pushHook(gitRequest, &result)
			break
		case HOOK_PUSH_TAG:
			c.tagPushHook(gitRequest, &result)
			break
		case HOOK_MERGE:
			c.mergeRequestHook(gitRequest, &result)
			break
		}
	} else {
		result.Result = 400
		result.ResultMessage = "Can't parse request"
	}

	return result
}

func (c *GitlabController) pushHook(gitlabObject models.GitlabRequest, result *models.Response) {

}

func (c *GitlabController) tagPushHook(gitlabObject models.GitlabRequest, result *models.Response) {

}

func (c *GitlabController) mergeRequestHook(gitlabObject models.GitlabRequest, result *models.Response) {

}
