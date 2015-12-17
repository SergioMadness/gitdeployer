package controllers

import (
	"fmt"
	"gitdeployer/config"
)

type GitDeployerBaseController struct {
}

func (cont *GitDeployerBaseController) PrepareServer(server config.Server) error {
	var err error

	if err = server.Deploy(); err == nil {
		fmt.Println(err)
	}

	return err
}
