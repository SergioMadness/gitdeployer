package models

type GitlabCommit struct {
	Id        string             `json:"id"`
	Message   string             `json:"message"`
	Timestamp string             `json:"timestamp"`
	Url       string             `json:"url"`
	Added     []string           `json:"added"`
	Modified  []string           `json:"modified"`
	Removed   []string           `json:"removed"`
	Author    GitlabChangeAuthor `json:"author"`
}
