package database

import (
	"context"
	"database/sql"
	"employees-manager/internal/server/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

type DB struct {
	Pool *sql.DB
	log  *logrus.Logger
}

func New(config *config.Config, log *logrus.Logger) (*DB, error) {
	pool, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cnl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cnl()

	if err := pool.PingContext(ctx); err != nil {
		return nil, err
	}
	return &DB{Pool: pool}, nil
}

// Close - closes the database connection
func (db DB) Close() error {
	return db.Pool.Close()
}

// Ping - checks the database connection
func (db DB) Ping() error {
	if err := db.Pool.Ping(); err != nil {
		db.log.Error(err)
		return err
	}
	return nil
}
