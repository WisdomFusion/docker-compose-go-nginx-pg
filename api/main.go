package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"database/sql"
	"encoding/json"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	bin, err := ioutil.ReadFile("/run/secrets/pg-secret")
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/pgdemodb?sslmode=disable", string(bin)))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM post")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var titles []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		titles = append(titles, title)
	}
	json.NewEncoder(w).Encode(titles)
}

func main() {
	log.Print("Prepare db...")
	if err := prepare(); err != nil {
		log.Fatal(err)
	}

	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", postHandler)
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
}

func prepare() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS post"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS post (id SERIAL, title VARCHAR)"); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO post (title) VALUES ($1);", fmt.Sprintf("Post #%d", i)); err != nil {
			return err
		}
	}
	return nil
}