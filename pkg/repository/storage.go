package repository

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"path/filepath"
	"runtime"
)

type Storage interface {
	RunMigrations(connectionString string) error
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return &storage{
		db: db,
	}
}
func (s *storage) RunMigrations(connecionString string) error {
	if connecionString == "" {
		return errors.New("repository: the connectionString was empty")
	}

	// get base path
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../")

	migrationPath := filepath.Join("file://", basepath, "repository/migrations")
	fmt.Println(migrationPath)
	m, err := migrate.New(migrationPath, connecionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}
	return nil
}
