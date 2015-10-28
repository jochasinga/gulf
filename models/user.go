package models

import (
	"math/rand"
	"strconv"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/mattn/go-sqlite3"
	cont "controllers"
)

/**** Enum for level ****/
type Level int

const (
	Free = iota
	Putter
	Pro
	Championship
)

var levels = [...]string{
	"Free",
	"Putter",
	"Pro",
	"Championship",
}

func (s Status) String() string {
	return statuses[s]
}
/**** Enum for level ****/

/***** User and related structs *****/
type User struct {
	Id          int64	 `json:"id"`
	Email		string   `json:"email"`
	Password   	string	 `json:"password"`
	Course		[]Course `json:"course"`
	Level		Level	 `json:"level"`
}

func (u *User) SaveToDB() int {

}

type Course struct {
	// Define what a Course is
	// A Course is an instance of a database cluster

	// Name of the course, i.e. evergreen-8340
	Name		string 	    `json:"name"`
	Route		cont.Route	`json:"route"`
}

func (c *Course) makeName(aliases []string) {
	// generate course's name, i.e. evergreen-8340
	i := rand.Intn( len(aliases) )
	c.Name = aliases[i] + "-" + strconv.Itoa( rand.Intn(9001) + 999 )
}

aliases := []string{
	"evergreen",
	"pinetree",
	"groundhog",
	"gopherhole",
	"greenbirdy",
	"handicap",
	"holeinone",
}
/***** User and related structs *****/

/***** Helper functions *****/
func NewUser(email, password string) (u *User) {
	u := &User{ Email: email, Password: password }
	return
}


func NewCourse() (c *Course) {
	c := &Course{}
	c.Name = c.makeName(aliases)
	return
}
