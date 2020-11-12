package config

import (
	`fmt`
	`os`

	`gopkg.in/yaml.v3`
)

type Config struct {
	Process map[string]string `yaml:"process"`
}

func (c *Config) Read(fn string) error {
	in, err := os.Open(fn)
	if nil!=err {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf(`Failed to read %s: %w`, fn, err)
	}
	defer in.Close()
	return yaml.NewDecoder(in).Decode(c)
}

func ReadConfig(fn string) (*Config, error) {
	c := &Config {
		Process: map[string]string{},
	}
	if err := c.Read(fn); nil!=err {
		return nil, err
	}
	return c, nil
}