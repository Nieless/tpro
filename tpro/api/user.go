package api

import (
	"fmt"
	"github.com/Nieless/tpro/tpro"
	"github.com/Nieless/tpro/tpro/mongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserService struct {
	mongo.UserStore
}

// GetUsers handler
func (svc *UserService) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.UserStore.GetUsers()
		if err != nil {
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}
		NewJSONWriter(w).Write(users, http.StatusOK)
		return
	}
}

// CreateUser handler
func (svc *UserService) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &mongo.User{}
		user.ID = bson.NewObjectId()
		if _, err := tpro.DecodeJSON(r, user); err != nil {
			NewJSONWriter(w).Write(err, http.StatusBadRequest)
			return
		}

		if err := svc.UserStore.CreateUser(user); err != nil {
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}

		NewJSONWriter(w).Write(user, http.StatusOK)
		return
	}
}

// UpdateUser handler
func (svc *UserService) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		if len(userID) == 0 {
			NewJSONWriter(w).Write(fmt.Errorf("user id missing"), http.StatusBadRequest)
			return
		}

		user, err := svc.UserStore.GetUser(bson.ObjectIdHex(userID))
		if err != nil {
			NewJSONWriter(w).Write(err, http.StatusBadRequest)
			return
		}

		if _, err := tpro.DecodeJSON(r, user); err != nil {
			NewJSONWriter(w).Write(err, http.StatusBadRequest)
			return
		}

		if err := svc.UserStore.UpdateUser(user); err != nil {
			fmt.Println(err)
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}

		NewJSONWriter(w).Write(user, http.StatusOK)
	}
}

// DeleteUser handler
func (svc *UserService) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		if len(userID) == 0 {
			NewJSONWriter(w).Write(fmt.Errorf("user id missing"), http.StatusBadRequest)
			return
		}

		// TODO validate user and return after deleting it

		err := svc.UserStore.DeleteUser(bson.ObjectIdHex(userID))
		if err != nil {
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}

		NewJSONWriter(w).Write("deleted", http.StatusOK)
		return
	}
}

// GetUser handler
func (svc *UserService) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		if len(userID) == 0 {
			NewJSONWriter(w).Write(fmt.Errorf("user id missing"), http.StatusBadRequest)
			return
		}

		user, err := svc.UserStore.GetUser(bson.ObjectIdHex(userID))
		if err != nil {
			NewJSONWriter(w).Write(err, http.StatusInternalServerError)
			return
		}

		NewJSONWriter(w).Write(user, http.StatusOK)
		return
	}
}
