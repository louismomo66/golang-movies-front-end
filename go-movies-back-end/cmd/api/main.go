package main

import (
	"log"
	"net/http"
)

// const port = 8080

type application struct {
	Domain string
}

func main() {
	//set application config
	var app application
	//read from command line

	//connect to the databaase
 app.Domain = "example.com"
 log.Println("Starting application on port 9000")
	//start a web server
	err := http.ListenAndServe("localhost:9000",app.routes())
	if err != nil{
		log.Fatal(err)
	}
}