package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "net/http"
    "../common"
    "time"
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
		fmt.Println(err)
    }
    defer db.Close()
    u_id := common.GetUser(w,r)
    if (u_id == 0) {
        rows, err := db.Query("SELECT NEXTVAL('nologin_usr_id_seq')")
        if err != nil {
            fmt.Println(err)
        }
        var nextVal int
        for rows.Next() {
            if err := rows.Scan(&nextVal); err != nil {
                fmt.Println(err)
            }
        }
        u_id = nextVal * -1
        common.SetUser(w,r,u_id)
    }

    if r.FormValue("update") == "1" {
        _, err = db.Exec(`UPDATE h_answer 
            SET correct = 1, created_at = $1
            WHERE respondent_id = $2
            AND question_id = $3`, 
            time.Now().Format("2006-01-02 15:04:05"),
            u_id,
            r.FormValue("question_id"))
        if err != nil {
            fmt.Println(err)
        }
    } else {
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
            fmt.Println(err)
        }
    }

    fmt.Fprint(w, `["1"]`)

}