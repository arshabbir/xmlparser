package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arshabbir/gmux/src/services"
	"github.com/gorilla/mux"
)

type controller struct {
	serv services.UserService
}

type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Controller interface {
	Start(chan int) error
	HandlePing(http.ResponseWriter, *http.Request)
	HandleGetUser(w http.ResponseWriter, r *http.Request)
}

func NewController(serv services.UserService) Controller {

	return &controller{serv: serv}
}

func (c *controller) Start(appStatus chan int) error {
	r := mux.NewRouter()

	addr := ":8080"

	//r.StrictSlash(true)
	r.StrictSlash(true).HandleFunc("/user/{id}/", c.HandleGetUser).Methods("GET")
	r.StrictSlash(true).HandleFunc("/user", c.HandleAddUser).Methods("POST")
	r.HandleFunc("/ping", c.HandlePing).Methods("GET")
	r.HandleFunc("/", http.NotFoundHandler().ServeHTTP)
	s := http.Server{
		Handler:      r,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Addr:         addr,
	}

	fmt.Println("Controller Started..")
	if err := s.ListenAndServe(); err != nil {
		appStatus <- 2
		return err
	}
	appStatus <- 1

	return nil

}

func (c *controller) HandlePing(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))

	return
}

func (c *controller) HandleGetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if vars["id"] == "" {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id is mandatory"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User Id is : %s", vars["id"])))

	return
}

func (c *controller) HandleAddUser(w http.ResponseWriter, r *http.Request) {

	var data user

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to Parse the request"))
		return

	}
	defer r.Body.Close()

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User  : %#v", data)))

	return
}
