package conf

type Config struct {
}

func (c *Config) FindAndLoad(path string) error {
	return nil
}

var Global = &Config{}
