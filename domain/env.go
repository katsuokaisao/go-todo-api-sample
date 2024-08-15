package domain

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Env struct {
	API *APIEnv
	DB  *DBEnv
}

func NewEnv() *Env {
	return &Env{
		API: &APIEnv{},
		DB:  &DBEnv{},
	}
}

func (e *Env) Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".envファイルを読み込めませんでした: %v\n", err)
	} else {
		log.Println(".envファイルが読み込まれました")
	}
	if err := env.Parse(e); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	return nil
}

type APIEnv struct {
	Addr string `env:"API_ADDR" envDefault:":8080"`
}

type DBEnv struct {
	Driver   string `env:"DB_DRIVER"`
	Address  string `env:"DB_ADDRESS"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
	Debug    bool   `env:"DB_DEBUG"`
}
