package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lbrictson/squint/pkg/conf"
	"github.com/lbrictson/squint/sql/migrations"
)

func New(configuration conf.Config) (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBPort, configuration.DBName)
	_, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		// Handle error if DB doesn't exist
		if strings.Contains(err.Error(), "does not exist") {
			rawDSN := fmt.Sprintf("postgres://%v:%v@%v:%v?sslmode=disable",
				configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBPort)
			builderConnection, err := pgx.Connect(context.Background(), rawDSN)
			if err != nil {
				return nil, err
			}
			defer builderConnection.Close(context.Background())
			_, err = builderConnection.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %v;", configuration.DBName))
			if err != nil {
				return nil, err
			}
			dsn = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
				configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBPort, configuration.DBName)
			_, err = pgx.Connect(context.Background(), dsn)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	err = runMigrations(dsn)
	if err != nil {
		return nil, err
	}
	dbPool, err := pgxpool.Connect(context.Background(), dsn)
	return dbPool, nil
}

func runMigrations(dsn string) error {
	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			return nil
		}
	}
	return err
}
