package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/functions"
)

func main() {
	http.HandleFunc("/static/", functions.StyleFunc)
	http.HandleFunc("/", functions.HandlerMainFunc)
	http.HandleFunc("/asciiart", functions.HandlerArtFunc)
	fmt.Println("runing server : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error", err)
	}
}
