package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database interface {
	// Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
	Exec(query string, args ...any) (sql.Result, error)
	// QueryRow executes a query that is expected to return at most one row.
	// QueryRow always returns a non-nil value. Errors are deferred until Row's
	// Scan method is called. If the query selects no rows, the *Row.Scan will
	// return ErrNoRows. Otherwise, *Row.Scan scans the first selected row and
	// discards the rest.
	QueryRow(query string, args ...any) *sql.Row
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SSLmode  string
}

func DBConfigFromEnv(filepath string) (*DatabaseConfig, error) {
	c := new(DatabaseConfig)
	err := godotenv.Load(filepath)
	if err != nil {
		return nil, err
	}
	c.Host = os.Getenv("DBHOST")
	_, err = fmt.Sscan(os.Getenv("DBPORT"), &c.Port)
	if err != nil {
		return nil, errors.Join(errors.New("error parsing DBPORT env variable to int"), err)
	}
	c.User = os.Getenv("DBUSER")
	c.Password = os.Getenv("DBPASSWORD")
	c.Dbname = os.Getenv("DBNAME")
	c.SSLmode = os.Getenv("DBSSLMODE")
	return c, nil
}

func SetupDatabase(c *DatabaseConfig) (*sql.DB, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Dbname, c.SSLmode,
	)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, errors.Join(errors.New("failed to connect to the database"), err)
	}
	return db, nil
}

func CleanDatabase(db Database) (sql.Result, error) {
	res, err := db.Exec("TRUNCATE TABLE banners")
	if err != nil {
		return nil, err
	}
	return res, nil
}
