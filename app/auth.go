package app

import (
	// "fmt"
	// "database/sql"
	// _ "github.com/lib/pq" // this driver for postgres
	"log"
	"net/http"
	"strconv"

	"../common"
)

// Auth temporary take session
func Auth(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("p") != "85ujd0d" {
		log.Print(r.URL.Path)
		http.NotFound(w, r)
		return
	}

	uID, _ := strconv.Atoi(r.FormValue("u"))
	common.SetUser(w, r, uID)
}
