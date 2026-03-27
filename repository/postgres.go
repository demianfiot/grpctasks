package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	log.Println("HOST:", os.Getenv("DB_HOST"))
	log.Println("PORT:", os.Getenv("DB_PORT"))
	log.Println("USER:", os.Getenv("DB_USER"))
	log.Println("PASSWORD:", os.Getenv("DB_PASSWORD"))
	log.Println("DBNAME:", os.Getenv("DB_NAME"))
	log.Println("SSLMODE:", os.Getenv("DB_SSLMODE"))
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE")))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DBConfigFromViper() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLmode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
