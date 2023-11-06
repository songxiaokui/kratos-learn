package redis

type Config struct {
	host     string
	port     string
	password string
	tp       string
	db       int
}

func NewConfig(host string) *Config {
	return &Config{
		host:     host,
		port:     "3306",
		password: "",
		tp:       "cluster",
		db:       0,
	}
}

type Options func(config *Config)

func WithPort(port string) Options {
	return func(c *Config) {
		c.port = port
	}
}

func WithPassword(pd string) Options {
	return func(c *Config) {
		c.password = pd
	}
}
