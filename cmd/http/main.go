package main

import (
	"github.com/Nieless/tpro/tpro"
	"github.com/Nieless/tpro/tpro/api"
	"github.com/Nieless/tpro/tpro/mongo"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	mongoConnStr := tpro.MustGetEnv("MONGO_CONNECT_STRING")

	var session *mgo.Session
	session, err := mgo.Dial(mongoConnStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	UserAPI := &api.UserService{
		UserStore: &mongo.UserDB{DB: session.DB("")},
	}

	myRouter.HandleFunc("/users", UserAPI.GetUsers()).Methods("GET")
	myRouter.HandleFunc("/users", UserAPI.CreateUser()).Methods("POST")
	myRouter.HandleFunc("/users/{id}", UserAPI.UpdateUser()).Methods("PUT")
	myRouter.HandleFunc("/users/{id}", UserAPI.DeleteUser()).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", UserAPI.GetUser()).Methods("GET")

	err = tpro.ListenAndServe(tpro.MustGetEnv("MONGO_CRUD_SERVICE_HTTP_PORT"), myRouter)
	if err != nil {
		panic(err)
	}
}
