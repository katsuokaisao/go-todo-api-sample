package rdb

import (
	"fmt"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/katsuokaisao/gin-play/domain"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RDB struct {
	db *gorm.DB
}

func NewRDB(env *domain.DBEnv) (*RDB, error) {
	var (
		db       *gorm.DB
		err      error
		retryNum uint64 = 15
	)

	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), retryNum)
	if err := backoff.Retry(
		func() error {
			db, err = connect(env)
			if err != nil {
				return err
			}
			return nil
		}, b,
	); err != nil {
		return nil, err
	}

	return &RDB{db: db}, nil
}

func connect(env *domain.DBEnv) (*gorm.DB, error) {
	var (
		dialector gorm.Dialector
		dialect   string
		host      string
		port      string
	)

	dialect = env.Driver
	add := strings.Split(env.Address, ":")
	host = add[0]
	if len(add) == 2 {
		port = add[1]
	}

	switch dialect {
	case "postgres":
		if port == "" {
			port = "5432"
		}
		args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, env.Username, env.Password, env.Database)
		dialector = postgres.Open(args)
	case "mysql":
		if port == "" {
			port = "3306"
		}
		args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", env.Username, env.Password, host, port, env.Database)
		dialector = mysql.Open(args)
	default:
		return nil, fmt.Errorf("invalid database dialect: %s", dialect)
	}

	db, err := gorm.Open(dialector, &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	if env.Debug {
		db = db.Debug()
	}

	return db, nil
}

func (r *RDB) NewSession(cfg *gorm.Session) *gorm.DB {
	return r.db.Session(cfg)
}
