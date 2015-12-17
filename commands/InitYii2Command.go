package commands

import (
	"fmt"
	"gitdeployer/helpers"
	"os"
	"errors"
)

type InitYii2Command struct {
	BaseCommand
}

func (i *InitYii2Command) Execute(path string) (string, error) {
	var out string
	var err error

	env := i.Get("env")

	fmt.Println("Yii2 Init")

	if env == "" {
		return "", errors.New("Need --env param")
	}

	currentDir, _ := os.Getwd()
	os.Chdir(path)

	out, err = helpers.Exec("php", "init --overwrite=All --env="+env)

	os.Chdir(currentDir)
	fmt.Println(out)
	fmt.Println(err)

	return out, err
}
