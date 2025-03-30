package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	App
	DB
}

type App struct {
	Port string `env:"PORT_APP"`
}

type DB struct {
	SqlDB           *sql.DB
	GormDB          *gorm.DB
	Host            string `env:"DB_HOST"`
	Name            string `env:"DB_NAME"`
	User            string `env:"DB_USER"`
	Password        string `env:"DB_PASSWORD"`
	Port            string `env:"DB_PORT"`
	Driver          string `env:"DB_DRIVER"`
	SSLMode         string `env:"DB_SSL_MODE"`
	ConnMaxLifeTime int    `env:"DB_CONN_MAX_LIFETIME"`
	MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS"`
}

func NewConfig(path string) (*Config, error) {
	if err := godotenv.Load(fmt.Sprintf("%v/.env", path)); err != nil {
		log.Println("Warning: .env file not found, using exported environment variables")
	}

	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
