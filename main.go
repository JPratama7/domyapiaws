package main

import (
	"github.com/domyid/domyapi/route"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	route.URL(w, r)
}
