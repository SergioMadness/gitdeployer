package controllers

import (
	"errors"
	"fmt"
	"gitdeployer/config"
)

type GitDeployerBaseController struct {
}

func (cont *GitDeployerBaseController) PrepareServer(server config.Server) error {
	if err := server.PrepareServer(); err != nil {
		fmt.Println("Can't prepare server")
		return errors.New("Can't prepare server")
	}

	if _, err := server.CloneRepo(); err != nil {
		fmt.Println("Can't clone repo")
		return errors.New("Can't clone repository")
	}

	return nil
}
