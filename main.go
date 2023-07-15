package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	fmt.Println("Server started at 9000:")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
