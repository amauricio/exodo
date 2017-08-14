package main

//exodo application
//author: mauricio

import (
    "fmt"
	"./app"
)

func handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "Mauricio y Rolly FOREVER %s!", r.URL.Path[1:])
}


func main() {
	app.Run()
	port := 80
	http.HandleFunc("/dashboard", handler)
	println( fmt.Sprintf("Starting exodo server on port %d", port))
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d",port), nil)
}
