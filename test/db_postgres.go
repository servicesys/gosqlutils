package test

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func NewDB() *sql.DB {

	strUser := "postgres"
	strPass := "postgres"
	strHost := "localhost"
	strDatabase := "agenda"
	connStr := "postgres://" + strUser + ":" + strPass + "@" + strHost + "/" + strDatabase + "?sslmode=disable"

	//
	connConfig, _ := pgx.ParseConfig(connStr)
	//	connConfig.Logger = myLogger
	connConfigStr := stdlib.RegisterConnConfig(connConfig)
	//	db, _ := sql.Open("pgx", connStr)
	//
	fmt.Println(connConfigStr)
	db, err := sql.Open("pgx", connConfigStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db

}

func NewDBX() *sqlx.DB {

	strUser := "postgres"
	strPass := "postgres"
	strHost := "localhost"
	strDatabase := "agenda"
	connStr := "postgres://" + strUser + ":" + strPass + "@" + strHost + "/" + strDatabase + "?sslmode=disable"

	dbx, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	dbx.SetMaxOpenConns(25)
	dbx.SetMaxIdleConns(25)
	dbx.SetConnMaxLifetime(5 * time.Minute)

	return dbx

}
