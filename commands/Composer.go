package commands

import (
	"fmt"
	"gitdeployer/config"
	"gitdeployer/helpers"
	"os"
)

func ComposerInstall(server config.Server) error {
	fmt.Println("Composer install")

	currentDir, _ := os.Getwd()
	os.Chdir(server.Path)
	_, error := helpers.Exec("composer", "install")
	os.Chdir(currentDir)

	return error
}
