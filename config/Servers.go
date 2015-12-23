package config

import (
	"errors"
	"fmt"
	"gitdeployer/helpers"
	"os"
	"strings"
)

type Server struct {
	Name          string
	Path          string
	DefaultBranch string
	GitUrl        string
	GitLogin      string
	GitPassword   string
	Commands      []*DeployerCommand
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
	return helpers.PrepareDir(s.Path)
}

func (s *Server) CloneRepo() (string, error) {
	return s.CloneTo(s.Path)
}

func (s *Server) PullRepo(branch string) (string, error) {
	s.Checkout(branch)
	return helpers.Exec("git", "pull", "origin", branch)
}

func (s *Server) Checkout(branch string) (string, error) {
	return helpers.Exec("git", "checkout", branch)
}

func (s *Server) CloneTo(path string) (string, error) {
	var resultStr string
	var err error
	
	currentDir, _ := os.Getwd()
	os.Chdir(path)
	fmt.Println(path)
	url := s.GitUrl
	url = strings.Replace(url, "http://", "http://"+s.GitLogin+":"+s.GitPassword+"@", 1)
	resultStr, err = helpers.Exec("git", "clone", s.GitUrl, ".")
	s.Checkout(s.DefaultBranch)
	os.Chdir(currentDir)

	return resultStr, err
}

func (s *Server) Deploy() error {
	fmt.Println("Prepare directory")
	if err := s.PrepareServer(); err != nil {
		fmt.Println(err)
		return errors.New("Can't prepare server")
	}
	fmt.Println("Cloning repo")
	if _, err := s.CloneTo(s.Path); err != nil {
		fmt.Println(err)
		return errors.New("Can't clone repository")
	}

	return nil
}
