package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type myHandle struct{}

var mux map[string]func(http.ResponseWriter, *http.Request)

func (*myHandle) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if h, ok := mux[req.URL.String()]; ok {
		h(res, req)
		return
	}
}
func InitServer() {
	server := http.Server{
		Addr:        ":7070",
		Handler:     &myHandle{},
		ReadTimeout: 6 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	Router(mux)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
	}
}
