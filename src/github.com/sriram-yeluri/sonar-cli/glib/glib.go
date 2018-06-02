package glib

import (
	"log"
)

func Error(err error,from string){
	if err != nil {
		log.Printf("Error from : %s", from)
		log.Fatal(err)
	}
}