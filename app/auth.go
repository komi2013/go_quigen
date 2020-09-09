package app

import (
    // "fmt"
    // "database/sql"
    // _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "strconv"
)

func Auth(w http.ResponseWriter, r *http.Request) {
    
    if r.FormValue("p") != "85ujd0d" {
        log.Print(r.URL.Path)
        http.NotFound(w, r)
        return
    }

    u_id, _ := strconv.Atoi(r.FormValue("u"))
    common.SetUser(w,r,u_id)
}