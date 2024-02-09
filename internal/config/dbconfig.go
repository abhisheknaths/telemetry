package config

type DBConfig struct {
	ConnString string `envconfig:"DB_CONN_STRING"`
}
