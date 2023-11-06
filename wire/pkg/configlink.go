package pkg

type Config struct {
	DB   DBconfig
	Sche ScheConfig
}

type DBconfig struct {
	Host     string
	Type     string
	Port     string
	Password string
	DB       string
}

type ScheConfig struct {
	Type string
}

func NewConfig() Config {
	return Config{
		DB: DBconfig{
			Host:     "127.0.0.1",
			Type:     "redis",
			Port:     "6379",
			Password: "sxk",
			DB:       "aust",
		},
		Sche: ScheConfig{
			Type: "linux",
		},
	}
}
