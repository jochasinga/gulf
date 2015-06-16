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
 */

package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
    //"io"
    //"io/ioutil"
    "strconv"
    "strings"
    //"time"

    mux "github.com/julienschmidt/httprouter"
)

// This is the slice holding the data. Should belong to db.go
var houses Houses
 	
// Display a welcome message on home page
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	title := "index"
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles(title + ".html")
	t.Execute(w, p)
}

func HouseNew(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	title := "new"
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles(title + ".html")
	t.Execute(w, p)
}

func HouseCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
    title := "new"
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles(title + ".html")
        t.Execute(w, p)
    } else {
        r.ParseForm()
        // logic part of the form
        fmt.Println("address:", r.FormValue("address"))
        fmt.Println("city:", r.FormValue("city"))
        fmt.Println("state:", r.FormValue("state"))
        fmt.Println("state:", r.FormValue("zipcode"))
        fmt.Println("built in:", r.FormValue("built"))
        fmt.Println("Owner:", r.FormValue("firstname"), r.FormValue("lastname"))
        fmt.Println("Status:", r.FormValue("status"))
        fmt.Println("Detail:", r.FormValue("detail"))
    }
}

/*
// Post new house here
func HouseCreate(w http.ResponseWriter, r*http.Request, _ mux.Params) {
	
	var house House

	// Read the request's body for new JSON data
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	// Replace with defer r.Body.Close()
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
*/

// Display all haunted houses
func HouseIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {

    // See model.go
    houses = FindAll()

    // Encode Go struct to JSON
    if err := json.NewEncoder(w).Encode(houses); err != nil {
        panic(err)
    }
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









