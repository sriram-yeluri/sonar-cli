package glib

import (
	"log"
)

var DEBUG bool = false

func Error(err error,from string){
	if err != nil {
		log.Printf("Error from : %s", from)
		log.Fatal(err)
	}
}
