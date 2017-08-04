package main
	
import (
    "fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Mauricio y Rolly FOREVER %s!", r.URL.Path[1:])
}

func main() {
	port := 80
	http.HandleFunc("/dashboard", handler)
	
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d",port), nil)
}
