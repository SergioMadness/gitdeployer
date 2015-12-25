package commands

import (
	"fmt"
	"gitdeployer/helpers"
	"os"
)

// Composer command
type ComposerCommand struct {
	BaseCommand
}

// Execute composer command
func (c *ComposerCommand) Execute(path string) (string, error) {
	var out string
	var err error

	fmt.Println("Composer install")

	commandExists, _ := helpers.IsCommandExists("composer")

	if !commandExists {
		DownloadComposer(path)
		out, err = ComposerPHPInstall(path)
	} else {
		out, err = helpers.Exec("composer", "install")
	}

	return out, err
}

// Execute 'php composer.phar install' command
func ComposerPHPInstall(path string) (string, error) {
	currentDir, _ := os.Getwd()
	os.Chdir(path)

	output, err := helpers.Exec("php", "composer.phar", "install")

	os.Chdir(currentDir)

	return output, err
}

// Download composer.phar
func DownloadComposer(path string) error {
	fmt.Println("Composer download")
	err := helpers.DownloadFile("https://getcomposer.org/composer.phar", path+"composer.phar")
	fmt.Println("Composer downloaded")

	return err
}
