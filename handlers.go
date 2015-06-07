/* Let's plan out our routes
 *
 * Greet user with a stunning home page message
 * "/"
 *
 * GET all Houses here
 * "/houses"
 *
 * POST new House here
 * "/houses"
 *
 * GET a House with a specific Id
 * "/houses/id/:id"
 *
 * GET House(s) in a specific City
 * "/houses/city/:city"
 *
 * GET House(s) in a specific State
 * "/houses/state/:state"
 *
 * GET the JSON as file
 * "/houses.json"
 */

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"
    "strings"
    //"time"

    mux "github.com/julienschmidt/httprouter"
)

// This is the slice holding the data
var houses Houses
 	
// Display a welcome message on home page
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1>Hello, welcome to Spooky.com</h1>")
}

// Display all haunted houses
func HouseIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {

 	// See model.go
	houses = FindAll()

 	// Encode Go struct to JSON
 	if err := json.NewEncoder(w).Encode(houses); err != nil {
 		panic(err)
	}
}

// Post new house here
func HouseCreate(w http.ResponseWriter, r*http.Request, _ mux.Params) {
	
	var house House

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to House struct
	if err := json.Unmarshal(body, &house); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	Create(house)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func HouseById(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		panic(err)
	}

	house := FindById(id, houses)

	if err := json.NewEncoder(w).Encode(house); err != nil {
		panic(err)
	}
}

func HouseByCity(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	// "/city/:city"

	w.WriteHeader(http.StatusOK)

	city := strings.Title(strings.Replace(ps.ByName("city"), "-", " ", -1))

	houses := FindByCity(city, houses)

	if err := json.NewEncoder(w).Encode(houses); err != nil {
 		panic(err)
	}
}

func HouseByState(w http.ResponseWriter, r *http.Request, ps mux.Params){
	// "/state/:state"

	w.WriteHeader(http.StatusOK)

	state := strings.ToUpper(ps.ByName("state"))

	houses := FindByState(state)

	if err := json.NewEncoder(w).Encode(houses); err != nil {
 		panic(err)
	}
}










