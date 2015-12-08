package commands

import (
	"errors"
	"fmt"
	"gitdeployer/config"
)

func Deploy(server config.Server) error {
	fmt.Println("Prepare directory")
	if err := server.PrepareServer(); err != nil {
		fmt.Println(err)
		return errors.New("Can't prepare server")
	}
	fmt.Println("Cloning repo")
	if _, err := server.CloneRepo(); err != nil {
		fmt.Println(err)
		return errors.New("Can't clone repository")
	}

	return nil
}
