package sql

// Config is SQL structure config.
type Config struct {
	Host string `json:"host" env:"SQL_HOST" env-default:"0.0.0.0"`
	Port string `json:"port" env:"SQL_PORT"`

	User     string `json:"user" env:"SQL_USER"`
	Password string `json:"password" env:"SQL_PASSWORD"`

	DBName string `json:"db_name" env:"SQL_DB_NAME"`
}
