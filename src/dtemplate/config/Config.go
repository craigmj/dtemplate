package config

import (
	`fmt`
	`os`
	`log`

	`gopkg.in/yaml.v3`
)

type Process struct {
	Exec string `yaml:"exec"`
	Prefix string `yaml:"prefix"`
	Suffix string `yaml:"suffix"`
}

func stringFromMap(m map[string]interface{}, k string, def string) string {
	v, ok := m[k]
	if ok {
		s, ok := v.(string)
		if ok {
			return s
		}
	}
	return def
}

func NewProcessFromMap(m map[string]interface{}) (*Process) {
	return &Process{
		Exec: stringFromMap(m, `exec`,``),
		Prefix: stringFromMap(m, `prefix`,``),
		Suffix: stringFromMap(m, `suffix`,``),
	}
}

type Config struct {
	Process map[string]*Process `yaml:"process"`
}

func (c *Config) Read(fn string) error {
	log.Printf(`reading config %s`, fn)
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
		Process: map[string]*Process{},
	}
	if err := c.Read(fn); nil!=err {
		return nil, err
	}
	return c, nil
}