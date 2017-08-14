package main

//exodo application
//author: mauricio

import (
	"./app"
)


func main() {
	//define host & port
	port := 80
	host := "0.0.0.0"

	app.Server(port, host)
}
