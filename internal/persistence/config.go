package persistence

import "fmt"

type Config struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"user"`
	Password string `env:"PASSWORD" envDefault:"password"`
	DBName   string `env:"NAME" envDefault:"mydatabase"`
}

func (c *Config) dsn() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		c.User,
		c.Password,
		c.DBName,
		c.Host,
		c.Port,
	)
}
