package controllers

import (
	"errors"
	"gitdeployer/config"
)

type GitDeployerBaseController struct {
}

func (cont *GitDeployerBaseController) PrepareServer(server config.Server) error {
	if err := server.PrepareServer(); err != nil {
		return errors.New("Can't prepare server")
	}

	if _, err := server.CloneRepo(); err != nil {
		return errors.New("Can't clone repository")
	}
	
	return nil
}
