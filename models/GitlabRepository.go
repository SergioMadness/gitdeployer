package models

type GitlabRepository struct {
	Name            string `json:"name"`
	Url             string `json:"name"`
	Description     string `json:"name"`
	Homepage        string `json:"homepage"`
	GitHttpUrl      string `json:"git_http_url"`
	GitSSHUrl       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}
