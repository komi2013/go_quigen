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

func CommentAdd(w http.ResponseWriter, r *http.Request) {
    usr_id := 1
    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	db, err := sql.Open("postgres", common.DB_CONNECT)
	if err != nil {
		log.Print(err)
	}
    _, err = db.Exec(`INSERT INTO t_comment(
        comment_txt,usr_id,question_id,created_at,usr_img
        ) VALUES($1,$2,$3,$4,$5)`, 
        r.FormValue("comment_txt"),
        usr_id,
        r.FormValue("question_id"),
        time.Now().Format("2006-01-02 15:04:05"),
        r.FormValue("usr_img") )
    if err != nil {
        log.Print(err)
    }
    fmt.Fprint(w, `{"Status":"1"}`)

}