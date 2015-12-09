package models

type GitlabRequest struct {
	ObjectKind   string           `json:"object_kind"`
	Before       string           `json:"before"`
	After        string           `json:"after"`
	Ref          string           `json:"ref"`
	UserId       int              `json:"user_id"`
	UserName     string           `json:"user_name"`
	UserEmail    string           `json:"user_email"`
	ProjectId    int              `json:"project_id"`
	Repository   GitlabRepository `json:"repository"`
	Commits      []GitlabCommit   `json:"commits"`
	TotalCommits int              `json:"total_commits_count"`
}
