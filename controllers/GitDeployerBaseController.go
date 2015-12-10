package controllers

import (
	"gitdeployer/commands"
	"gitdeployer/config"
)

type GitDeployerBaseController struct {
}

func (cont *GitDeployerBaseController) PrepareServer(server config.Server) error {
	var err error

	if err = server.Deploy(); err == nil {
		if _, err = commands.ComposerInstall(server.Path); err == nil {

		}
	}

	return err
}
