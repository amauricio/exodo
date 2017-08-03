package main
	
import . "github.com/amauricio/exodo"
import (
    "fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Malcon rosquete %s!", r.URL.Path[1:])
}

func main() {
	port := 80
	http.HandleFunc("/dashboard", handler)
	
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d",port), nil)
}