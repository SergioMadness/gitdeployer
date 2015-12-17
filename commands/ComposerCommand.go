package commands

import (
	"fmt"
	"gitdeployer/helpers"
	"os"
)

type ComposerCommand struct {
	BaseCommand
}

func (c *ComposerCommand) Execute(path string) (string, error) {
	var out string
	var err error

	fmt.Println("Composer install")

	currentDir, _ := os.Getwd()
	os.Chdir(path)

	commandExists, _ := helpers.IsCommandExists("composer")

	if !commandExists {
		DownloadComposer(path)
		out, err = ComposerPHPInstall(path)
	} else {
		out, err = helpers.Exec("composer", "install")
	}
	os.Chdir(currentDir)
	fmt.Println(out)
	fmt.Println(err)

	return out, err
}

func ComposerPHPInstall(path string) (string, error) {
	currentDir, _ := os.Getwd()
	os.Chdir(path)

	output, err := helpers.Exec("php", "composer.phar", "install")

	os.Chdir(currentDir)

	return output, err
}

func DownloadComposer(path string) error {
	fmt.Println("Composer download")

	err := helpers.DownloadFile("https://getcomposer.org/composer.phar", path+"composer.phar")
	fmt.Println(err)
	fmt.Println("Composer downloaded")

	return err
}
