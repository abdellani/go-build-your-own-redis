package config

import "strings"

type Config struct {
	RDB RDBConfig
}

type RDBConfig struct {
	Dir      string
	FileName string
}

func (c *Config) IsRdbProvided() bool {
	if strings.Compare(c.RDB.FileName, "") == 0 {
		return false
	} else {
		return true
	}
}

func New() *Config {
	return &Config{
		RDB: RDBConfig{},
	}
}
