package app

//exodo application
//author: mauricio

import (
	"fmt"
	"net/http"
)

type collect struct{
	url string
	controller func(w http.ResponseWriter, r *http.Request) bool
}

func render(w http.ResponseWriter, r *http.Request){
	//render as html
	w.Header().Set("Content-Type", "text/html")
	routes := Routes()
	url := r.URL.Path
	
	for _, element  := range routes{
		if url == element.url {
			element.controller(w, r)
			return
		}
	}
	fmt.Fprintf(w, "<h1>Error 404</h1>")
}

func api(w http.ResponseWriter, r *http.Request){
	//render as html
	w.Header().Set("Content-Type", "application/json")
	routes := Routes()
	url := r.URL.Path
	
	for _, element  := range routes{
		if url == element.url {
			element.controller(w, r)
			return
		}
	}
	fmt.Fprintf(w, "{\"error\":true}")
}

func Server(port int, host string){

	// get all routes

	//start new server
	mux  := http.NewServeMux()

	
	mux.HandleFunc( "/"  , func(w http.ResponseWriter, r *http.Request){
		render(w, r)
	})
	mux.HandleFunc( "/api/"  , func(w http.ResponseWriter, r *http.Request){
		api(w, r)
	})

	println( fmt.Sprintf("Starting exodo server on port %d", port))
    http.ListenAndServe(fmt.Sprintf("%s:%d",host, port), mux)
}
