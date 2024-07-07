package utils

import (
	"log"
)

var DEBUG bool = false


func Error(err error,from string){
	if err != nil {
		log.Printf("[ERROR] Error from function : %s", from)
		log.Fatal(err)
	}
}
