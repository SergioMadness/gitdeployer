package commands

import (
	"fmt"
	"gitdeployer/config"
)

func CreateCommand(name string) CommandInterface {
	fmt.Println(name)
	switch name {
	case "composer":
		return new(ComposerCommand)
		break
	}

	return nil
}

func ExecuteCommandList(commands []*config.DeployerCommand, executionPath string) (string, error) {
	var result string
	var err error

	for _, dc := range commands {
		if comm := CreateCommand(dc.Name); comm != nil {
			comm.SetConfiguration(dc.Config)
			commOutput, commError := comm.Execute(executionPath)
			err = commError
			result += commOutput
		}
	}

	return result, err
}
