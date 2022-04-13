package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	JWT_KEY     string
}

func InitConfiguration() Config {
	// logic dapatin env
	// file(.env, env.yaml, env.json,...), system env

	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "root",
		DB_NAME:     "",
		DB_PORT:     "",
		DB_HOST:     "",
		JWT_KEY:     "l0cal",
	}
}
