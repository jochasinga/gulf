// router here

package main  

import (
	"log"
	"net/http"
	"fmt"
	ct "./controllers"

	mux "github.com/julienschmidt/httprouter"
)

func main() {
	
	router := mux.New()

	for _, rt := range ct.routes {
		router.Handle(rt.Method, rt.Path, rt.Handle)
	}

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}