package controllers

import(
	"fmt"
	"net/http"
)
func MauricioController(w http.ResponseWriter, r *http.Request) bool{
	fmt.Fprintf(w, "{\"mauricio\":true}")
	return false
}