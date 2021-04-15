package main

import (
	"log"
	"os"

	"github.com/siniak/golang-training-Social-network/pkg/data"
	"github.com/siniak/golang-training-Social-network/pkg/db"
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
	//userData := data.NewUserData(conn)
	//readAll
	/*users, err := userData.ReadAll()
	if err != nil {
		log.Println(err)
	}
	log.Println("values: ", &users)*/
	/*
		//read
		var readUser *data.User
		readUser, err = userData.Read(45)
		if err != nil {
			log.Println(err)
		}
		log.Println(readUser)*/
	/*//add
	id, err := userData.Add(data.User{
		Login:     "login6",
		Password:  "1234",
		Gender:    false,
		Email:     "e6@mail.ru",
		LastName:  "TestAdd",
		FirstName: "TestAddF",
		Birthday:   time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC),
		CityName:      "Minsk",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("User id is: ", id)*/
	/*//delete
	err = userData.Delete(45)
	if err != nil {
		log.Println(err)
	}
	//*/
	/*//update
	err = userData.Update(46, "newLogin")
	if err != nil {
		log.Println(err)
	}*/

	stateData := data.NewStateData(conn)
	//readAll
	/*states, err := stateData.ReadAll()
	if err != nil {
		log.Println(err)
	}
	log.Println("values: ", &states)*/
	//add
	/*id, err:= stateData.Add(data.State{
		Name:      "Kursk",
		CountryId: 2,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inserted state id is:", id)*/
	//Delete
	err = stateData.Delete(6)
	if err != nil {
		log.Println(err)
	}
	/*//update
	err = stateData.Update(6, "Vietka")
	if err != nil {
		log.Println(err)
	}*/
}
