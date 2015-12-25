package commands

// Basic command
type BaseCommand struct {
	config map[string]string
}

// Set command configuration
func (c *BaseCommand) SetConfiguration(config map[string]string) {
	c.config = config
}

// Get command configuration
func (c *BaseCommand) GetConfiguration() map[string]string {
	return c.config
}

// Get param
func (c *BaseCommand) Get(key string) string {
	return c.config[key]
}

// Set param
func (c *BaseCommand) Set(key, value string) {
	c.config[key] = value
}
