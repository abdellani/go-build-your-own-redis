package config

type Config struct {
	RDB RDBConfig
}

type RDBConfig struct {
	Dir      string
	FileName string
}

func New() *Config {
	return &Config{
		RDB: RDBConfig{},
	}
}
