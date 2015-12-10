package commands

type BaseCommand struct {
	config map[string]string
}

func (c *BaseCommand) SetConfiguration(config map[string]string) {
	c.config = config
}

func (c *BaseCommand) GetConfiguration() map[string]string {
	return c.config
}

func (c *BaseCommand) Get(key string) string {
	return c.config[key]
}

func (c *BaseCommand) Set(key, value string) {
	c.config[key] = value
}
