package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/VSiniak12/golang-training-Social-network/pkg/api"
	"github.com/VSiniak12/golang-training-Social-network/pkg/data"
	"github.com/VSiniak12/golang-training-Social-network/pkg/db"
	"github.com/gorilla/mux"
	"github.com/pressly/goose"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "social_network"
	}
	if password == "" {
		password = "1234"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	connDB, err := conn.DB()
	connDB.Exec("SET search_path TO social_network")
	err = goose.Up(connDB, "./pkg/db/migrations")
	if err != nil {
		log.Fatalf("can't connect to goose, error: %v", err)
	}
	r := mux.NewRouter()
	stateData := data.NewStateData(conn)
	api.ServeStateResource(*stateData, r)
	userData := data.NewUserData(conn)
	api.ServeUserResource(*userData, r)
	r.Use(mux.CORSMethodMiddleware(r))
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}
