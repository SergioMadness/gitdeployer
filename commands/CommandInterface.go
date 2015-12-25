package commands

// Common command interface
type CommandInterface interface {
	Execute(path string) (string, error)
	SetConfiguration(config map[string]string)
}