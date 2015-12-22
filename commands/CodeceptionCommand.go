package commands

import (
	"fmt"
	"gitdeployer/helpers"
	"os"
)

type CodeceptionCommand struct {
	BaseCommand
}

func (i *CodeceptionCommand) Execute(path string) (string, error) {
	var out string
	var err error

	fmt.Println("Codeception tests")

	currentDir, _ := os.Getwd()
	os.Chdir(path + "/tests")

	commandExists, _ := helpers.IsCommandExists("codecept")

	if !commandExists {
		DownloadCodecept(path)
		out, err = CodeceptPHP(path)
	} else if out, err = helpers.Exec("codecept", "build"); err == nil {
		out, err = helpers.Exec("codecept", "run", "--steps")
	}
	os.Chdir(currentDir)

	return out, err
}

func CodeceptPHP(path string) (string, error) {
	var output string
	var err error

	currentDir, _ := os.Getwd()
	os.Chdir(path)

	if output, err = helpers.Exec("php", "codecept.phar", "build"); err == nil {
		output, err = helpers.Exec("php", "codecept.phar", "run", "--steps")
	}

	os.Chdir(currentDir)

	return output, err
}

func DownloadCodecept(path string) error {
	fmt.Println("Codeception download")

	err := helpers.DownloadFile("http://codeception.com/codecept.phar", path+"codecept.phar")
	fmt.Println("Codeception downloaded")

	return err
}
