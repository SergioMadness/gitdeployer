package commands

import (
	"fmt"
	"gitdeployer/helpers"
	"os"
)

// Codeception command
// 
// Init testing by codeception
type CodeceptionCommand struct {
	BaseCommand
}

// Execute command
func (i *CodeceptionCommand) Execute(path string) (string, error) {
	var out string
	var err error

	fmt.Println("Codeception tests")

	dir := i.Get("Dir")

	currentDir, _ := os.Getwd()
	os.Chdir(path + "/" + dir)

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

// Execute codeception.phar
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

// Download codeception.phar
func DownloadCodecept(path string) error {
	fmt.Println("Codeception download")

	err := helpers.DownloadFile("http://codeception.com/codecept.phar", path+"codecept.phar")
	fmt.Println("Codeception downloaded")

	return err
}
