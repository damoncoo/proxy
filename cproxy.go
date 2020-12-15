package main

import (
	"log"
	"net/http"

	"github.com/smartystreets/cproxy/v2"
)

func initCproxy() bool {
	handler := cproxy.New()
	log.Println("Listening on:", "*:8090")
	_ = http.ListenAndServe(":8090", handler)

	return true
}
