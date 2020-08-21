package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    // "encoding/json"
    // "encoding/base64"
    // "os"
    // "strings"
    "time"
    // "strconv"
    // "../table"
)

func MessageAdd(w http.ResponseWriter, r *http.Request) {
    u_id := common.GetUser(w,r)
    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	db, err := sql.Open("postgres", common.DB_CONNECT)
	if err != nil {
		log.Print(err)
    }
    defer db.Close()
    _, err = db.Exec(`INSERT INTO t_message(
        sender,receiver,message_content,created_at,sender_img
        ) VALUES($1,$2,$3,$4,$5)`, 
        u_id,
        r.FormValue("receiver"),
        r.FormValue("message_content"),
        time.Now().Format("2006-01-02 15:04:05"),
        r.FormValue("sender_img") )
    if err != nil {
        log.Print(err)
    }
    fmt.Fprint(w, `{"Status":"1"}`)

}