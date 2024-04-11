package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Используются массивы т.к. они быстрее в случае нашего сервиса.
// Использование дополнительной таблицы будет менее эффективно, хотя бы из-за join-ов
// Если интересно или есть сомнения насчёт 1НФ, вот статья и бенчмарки которые я нашёл:
// http://www.databasesoup.com/2015/01/tag-all-things.html
// http://www.databasesoup.com/2015/01/tag-all-things-part-2.html
const bannersDDL = `CREATE TABLE IF NOT EXISTS public.banners (
    id serial4 NOT NULL,
    "content" json NOT NULL,
    feature_id int4 NOT NULL,
    tag_ids _int4 NOT NULL,
    is_active bool NOT NULL,
    CONSTRAINT banners_pk PRIMARY KEY (id)
);`

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
	_, err = db.Exec(bannersDDL)
	if err != nil {
		return nil, errors.Join(errors.New("failed to automigrate to the database"), err)
	}
	return db, nil
}

func CleanDatabase(db *sql.DB) (sql.Result, error) {
	res, err := db.Exec("TRUNCATE TABLE banners")
	if err != nil {
		return nil, err
	}
	return res, nil
}
