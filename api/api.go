package api

import (
	"fmt"
	rest "mini/rest"
	user "mini/user"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey home page")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	user.AllUsers(w, r)
}

func first(w http.ResponseWriter, r *http.Request) {
	rest.First(w, r)
}

func second(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rest.Second(w, r, params["pid"])
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/allUsers", allUsers)
	r.HandleFunc("/first", first)
	r.HandleFunc("/second/{pid}", second)
	// r.HandleFunc("/third", third)
	// r.HandleFunc("/forth", forth)
	// r.HandleFunc("/fifth", fifth)
	fmt.Println(http.ListenAndServe(":3030", r))
}
func Run() {
	handleRequest()
}
