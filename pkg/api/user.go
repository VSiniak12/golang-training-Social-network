package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VSiniak12/golang-training-Social-network/pkg/data"
	"github.com/gorilla/mux"
)

type userApi struct {
	data *data.UserData
}

func ServeUserResource(data data.UserData, r *mux.Router) {
	api := &userApi{data: &data}
	r.HandleFunc("/users", api.getAllUsers).Methods("GET")
	r.HandleFunc("/users", api.createUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", api.deleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", api.updateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", api.getUser).Methods("GET")
}

func (a userApi) getAllUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := a.data.ReadAll()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
	writer.WriteHeader(http.StatusOK)
}

func (a userApi) getUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := a.data.Read(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		log.Printf("failed writing to JSON: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
	writer.WriteHeader(http.StatusOK)
}

func (a userApi) createUser(writer http.ResponseWriter, request *http.Request) {
	user := new(data.User)
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.Add(user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	writer.WriteHeader(http.StatusCreated)
}

func (a userApi) updateUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user := new(data.User)
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user.IdUser = id
	_, err = a.data.Update(user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (a userApi) deleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.Delete(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
