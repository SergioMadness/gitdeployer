package commands

import (
	"fmt"
	"gitdeployer/config"
	"gitdeployer/modules/logger/interfaces"
)

// Command fabric
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

// Execute list of commands
func ExecuteCommandList(commands []*config.DeployerCommand, executionPath string, logger interfaces.LoggerInterface) (string, error) {
	var result string
	var err error

	for _, dc := range commands {
		if err == nil {
			if comm := CreateCommand(dc.Name); comm != nil {
				comm.SetConfiguration(dc.Config)
				commOutput, commError := comm.Execute(executionPath)
				err = commError
				result += commOutput
				if err != nil {
					logger.Log(dc.Name, err.Error())
				}
				logger.Log(dc.Name, commOutput)
			}
		}
	}

	logger.Flush()

	return result, err
}

// Prepare release and deploy to server
func TestAndDeploy(server *config.Server, dir string, log interfaces.LoggerInterface) error {
	var err error

	if _, err = server.CloneTo(dir); err == nil {
		fmt.Println("Dir: " + dir + "/")
		if _, err = ExecuteCommandList(server.Commands, dir+"/", log); err == nil {
			if err = server.Deploy(); err != nil {
				log.Log("application", err.Error())
			} else {
				_, err = ExecuteCommandList(server.DeployCommands, server.Path, log)
			}
		}
	}
	log.Log("aplication", err.Error())
	log.Flush()
	return err
}
