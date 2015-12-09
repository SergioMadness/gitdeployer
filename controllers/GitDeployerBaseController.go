package controllers

import (
	"gitdeployer/commands"
	"gitdeployer/config"
)

type GitDeployerBaseController struct {
}

func (cont *GitDeployerBaseController) PrepareServer(server config.Server) error {
	var err error
	
	if err = commands.Deploy(server); err == nil {
		if err = commands.ComposerInstall(server); err == nil {
			
		}
	}

	return err
}
