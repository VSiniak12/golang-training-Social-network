package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siniak/golang-training-Social-network/pkg/data"
)

type stateApi struct {
	data *data.StateData
}

func ServeStateResource(data data.StateData, r *mux.Router) {
	api := &stateApi{data: &data}
	r.HandleFunc("/states", api.getAllStates).Methods("GET")
	r.HandleFunc("/states", api.createState).Methods("POST")
	r.HandleFunc("/states/{id:[0-9]+}", api.deleteState).Methods("DELETE")
	r.HandleFunc("/states/{id:[0-9]+}", api.updateState).Methods("PUT")
	r.HandleFunc("/states/{id:[0-9]+}", api.getState).Methods("GET")
}

func (a stateApi) getAllStates(writer http.ResponseWriter, request *http.Request) {
	states, err := a.data.ReadAll()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(states)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
	writer.WriteHeader(http.StatusOK)
}

func (a stateApi) getState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	state, err := a.data.Read(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(state)
	if err != nil {
		log.Printf("failed writing to JSON: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
	writer.WriteHeader(http.StatusOK)
}

func (a stateApi) createState(writer http.ResponseWriter, request *http.Request) {
	state := new(data.State)
	err := json.NewDecoder(request.Body).Decode(&state)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.Add(state)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/states/%d", id))
	writer.WriteHeader(http.StatusCreated)
}

func (a stateApi) updateState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	state := new(data.State)
	err = json.NewDecoder(request.Body).Decode(&state)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	state.IdState = id
	_, err = a.data.Update(state)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (a stateApi) deleteState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("problem with parse to int : %v\n", err)
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
