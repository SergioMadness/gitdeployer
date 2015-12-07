package models;

type GitlabRequest struct {
	ObjectKind string `json:"object_kind"`;
	Before string `json:"before"`;
	After string `json:"after"`;
	Ref string `json:"ref"`;
	UserId string `json:"user_id"`;
	UserName string `json:"user_name"`;
	UserEmail string `json:"user_email"`;
	ProjectId string `json:"project_id"`;
	Repository GitlabRepository `json:"repository"`;
}