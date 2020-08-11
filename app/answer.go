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

func Answer(w http.ResponseWriter, r *http.Request) {
    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
    }
    u_id := common.GetUser(w,r)
    if (u_id == 0) {
        rows, err := db.Query("SELECT NEXTVAL('nologin_usr_id_seq')")
        if err != nil {
            log.Print(err)
        }
        var nextVal int
        for rows.Next() {
            if err := rows.Scan(&nextVal); err != nil {
                log.Print(err)
            }
        }
        // fmt.Printf("nextVal %#v\n", nextVal)
        u_id = nextVal * -1
        common.SetUser(w,r,u_id)
    }
    // fmt.Printf("%#v\n", u_id)
    // var choices [4]string
    // json.Unmarshal([]byte(r.FormValue("preChoices")), &choices)
    // fmt.Printf("%#v\n", choices[0])
    // return

    _, err = db.Exec(`INSERT INTO h_answer(
        question_id   
        ,question_txt    
        ,category_id    
        ,respondent_id 
        ,respondent_img
        ,correct
        ,created_at
        ) VALUES($1,$2,$3,$4,$5,$6,$7)`, 
        r.FormValue("question_id"),
        r.FormValue("question_txt"),
        r.FormValue("category_id"),
        u_id,
        r.FormValue("respondent_img"),
        r.FormValue("correct"),
        time.Now().Format("2006-01-02 15:04:05"))
    if err != nil {
        log.Print(err)
    }
    fmt.Fprint(w, `["1"]`)

}