package config

import (
	"gitdeployer/helpers"
	"os"
)

type Server struct {
	Name          string
	Path          string
	DefaultBranch string
	GitUrl        string
	GitLogin      string
	GitPassword   string
}

func CreateServer(name, path, defaultBranch, gitUrl, gitLogin, gitPassword string) *Server {
	result := new(Server)

	result.Name = name
	result.Path = path
	result.DefaultBranch = defaultBranch
	result.GitUrl = gitUrl
	result.GitLogin = gitLogin
	result.GitPassword = gitPassword

	return result
}

func (s *Server) PrepareServer() error {
	var result error
	if !helpers.IsPathExists(s.Path) {
		result = os.MkdirAll(s.Path, 0644)
	}
	return result
}

func (s *Server) CloneRepo() (string, error) {
	var resultStr string
	var err error

	currentDir, _ := os.Getwd()
	os.Chdir(s.Path)
	resultStr, err = helpers.Exec("git", "clone", s.GitUrl, ".")
	s.Checkout(s.DefaultBranch)
	os.Chdir(currentDir)

	return resultStr, err
}

func (s *Server) PullRepo(branch string) (string, error) {
	s.Checkout(branch)
	return helpers.Exec("git", "pull", "origin", branch)
}

func (s *Server) Checkout(branch string) (string, error) {
	return helpers.Exec("git", "checkout", branch)
}
