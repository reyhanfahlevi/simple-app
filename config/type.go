package config

type Config struct {
	App       AppConfig       `yaml:"app"`
	Databases DatabasesConfig `yaml:"databases"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type DatabasesConfig struct {
	postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

type ElasticConfig struct {
	Host string `yaml:"host"`
}
