package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(ctx context.Context, cfg Config) (dbPool *pgxpool.Pool, err error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbPool, err = pgxpool.Connect(ctx, connString)
	if err != nil {
		logrus.Fatalf("unable to connect DB %s", err.Error())
	}
	return dbPool, err
}
