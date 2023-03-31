package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// handler function for the home page
func home(w http.ResponseWriter, r *http.Request) {

	// display information about George
	http.ServeFile(w, r, "html/index.html")
}

// handler function for the greeting page
func greeting(w http.ResponseWriter, r *http.Request) {

	// get the current date and time
	now := time.Now().Format("Monday, January 02 2006")

	// read in the template file
	tmpl, _ := template.ParseFiles("html/greeting.html")

	// do the substitution
	tmpl.Execute(w, now)

}

// handler function for the random quote page
func random(w http.ResponseWriter, r *http.Request) {

	// an array of quotes
	quotes := []string{
		"The greatest glory in living lies not in never falling, but in rising every time we fall.",
		"The way to get started is to quit talking and begin doing.",
		"Your time is limited, don't waste it living someone else's life.",
		"We are not defined by our past, but by the choices we make in the present.",
		"It's not about how hard you hit. It's about how hard you can get hit and keep moving.",
	}

	// generate a random number using the current time as the seed
	rand.Seed(time.Now().UnixNano())

	// get a random quote from the array
	quote := quotes[rand.Intn(len(quotes))]

	// read in the template file
	tmpl, _ := template.ParseFiles("html/random.html")

	// do the substitution
	tmpl.Execute(w, quote)

}

func main() {

	// create a new servermux
	mux := http.NewServeMux()

	// register the handler functions for each page
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	//create a web server
	log.Println(("Starting server on port :4000"))
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
