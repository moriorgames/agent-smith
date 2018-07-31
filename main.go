package main

import (
	"log"
	"net/http"
	"github.com/moriorgames/agent-smith/src"
)

func main() {
	router := createRouter()
	log.Fatal(http.ListenAndServe(":9090", router))
}
