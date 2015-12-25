package commands

import (
	"errors"
	"fmt"
	"gitdeployer/helpers"
	"os"
)

// Yii2 initialisation
type InitYii2Command struct {
	BaseCommand
}

// Execute 'php init' on Yii2 project
func (i *InitYii2Command) Execute(path string) (string, error) {
	var out string
	var err error

	env := i.Get("Env")

	fmt.Println("Yii2 Init " + env)

	if env == "" {
		return "", errors.New("Need Env param")
	}

	currentDir, _ := os.Getwd()
	os.Chdir(path)

	out, err = helpers.Exec("php", "init", "--overwrite=All", "--env="+env)

	os.Chdir(currentDir)

	return out, err
}
