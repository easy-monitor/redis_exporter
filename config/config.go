package config

type Config struct {
	RedisModule []*RedisModule `json:"module" yaml:"module"`
}

type RedisModule struct {
	Name     string `json:"name" yaml:"name"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}
