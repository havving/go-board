package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
	// postgres driver
	_ "github.com/lib/pq"
)

func TestConnection(t *testing.T) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open("postgres", conn)
	defer db.Close()

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
	log.Println("Success connection to DB")
}

func TestSqlxConnection(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=heybin password=1234 dbname=board sslmode=disable")
	if err != nil { // sqlx.Connect안에 이미 Ping() 검사를 함
		log.Fatalln(err)
		return
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO todo VALUES ($1, $2, $3)", 2, "reading", true)
	tx.Commit()
}
