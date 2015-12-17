package commands

import (
	"gitdeployer/config"
)

func CreateCommand(name string) CommandInterface {
	switch name {
	case "composer":
		return new(ComposerCommand)
	case "codeception":
		return new(CodeceptionCommand)
	case "yii2-init":
		return new(InitYii2Command)
	}

	return nil
}

func ExecuteCommandList(commands []*config.DeployerCommand, executionPath string) (string, error) {
	var result string
	var err error

	for _, dc := range commands {
		if err == nil {
			if comm := CreateCommand(dc.Name); comm != nil {
				comm.SetConfiguration(dc.Config)
				commOutput, commError := comm.Execute(executionPath)
				err = commError
				result += commOutput
			}
		}
	}

	return result, err
}
