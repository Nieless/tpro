package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDB struct {
	DB *mgo.Database
}

type User struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Age  int           `json:"age" bson:"age"`
}

// UserStore defines the operations of User store.
type UserStore interface {
	CreateUser(User *User) error
	UpdateUser(User *User) error
	DeleteUser(userID bson.ObjectId) error
	GetUsers() ([]*User, error)
	GetUser(userID bson.ObjectId) (*User, error)
}

// AddUser adds a User document to Users collection in Mongo
func (mgo *UserDB) CreateUser(User *User) error {
	// Create variable for collection
	cUserConfig := mgo.DB.C("users")
	return cUserConfig.Insert(&User)
}

// UpdateUser deletes a User document from Mongo
func (mgo *UserDB) UpdateUser(User *User) error {
	// Create variable for collection
	cUserConfig := mgo.DB.C("users")
	return cUserConfig.UpdateId(User.ID, User)
}

// DeleteUser deletes a User document from Mongo
func (mgo *UserDB) DeleteUser(userID bson.ObjectId) error {
	// Create variable for collection
	cUserConfig := mgo.DB.C("users")
	return cUserConfig.RemoveId(userID)
}

// GetUsers gets a Mongo Users document from Mongo
func (mgo *UserDB) GetUsers() ([]*User, error) {
	// Create variable for collection
	cUserConfig := mgo.DB.C("users")
	Users := make([]*User, 0)
	if err := cUserConfig.Find(nil).All(&Users); err != nil {
		return nil, err
	}
	return Users, nil
}

// GetUser gets a Mongo User document from Mongo
func (mgo *UserDB) GetUser(userID bson.ObjectId) (*User, error) {
	// Create variable for collection
	cUserConfig := mgo.DB.C("users")
	user := &User{}

	if err := cUserConfig.FindId(userID).One(user); err != nil {
		return nil, err
	}
	return user, nil
}
