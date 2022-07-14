package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/index.html")
}

func main() {

	addr := ":8081"

	router := httprouter.New()
	router.GET("/", Index)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
