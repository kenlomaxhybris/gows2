package main

import (
	"log"
	"net/http"

	"github.com/kenlomaxhybris/goworkshopII/router"
)

func main() {
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":8089", r))
}
