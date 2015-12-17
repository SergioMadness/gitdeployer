package commands

type CommandInterface interface {
	Execute(path string) (string, error)
	SetConfiguration(config map[string]string)
}