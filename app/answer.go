package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "encoding/json"
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
		log.Fatal(err)
    }
    u_id := common.GetUser(w,r)
    if (u_id == 0) {
        rows, err := db.Query("SELECT NEXTVAL('nologin_usr_id_seq')")
        if err != nil {
            log.Fatal(err)
        }
        type Nextval struct {
            Int        int
        }
        var next []Nextval
        for rows.Next() {
            n := Nextval{}
            if err := rows.Scan(&n.Int); err != nil {
                log.Fatal(err)
            }
            next = append(next,n)
        }
        u_id = next[0].Int * -1
        common.SetUser(w,r,u_id)
    }
    fmt.Printf("%#v\n", u_id)
    var choices [4]string
    json.Unmarshal([]byte(r.FormValue("preChoices")), &choices)
    // fmt.Printf("%#v\n", choices[0])
    // return

    _, err = db.Exec(`INSERT INTO h_answer(
        question_id   
        ,question_txt  
        ,generator_id  
        ,generator_img 
        ,asked_at      
        ,choice_0    
        ,choice_1    
        ,choice_2    
        ,choice_3    
        ,reference     
        ,question_type 
        ,category_id   
        ,question_img  
        ,respondent_id 
        ,respondent_img
        ,sequence      
        ,mytext        
        ,spend         
        ,mychoice      
        ,explanation
        ,created_at
        ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)`, 
        r.FormValue("question_id"),
        r.FormValue("question_txt"),
        r.FormValue("generator_id"),
        r.FormValue("generator_img"),
        r.FormValue("asked_at"),
        choices[0],
        choices[1],
        choices[2],
        choices[3],
        r.FormValue("reference"),
        r.FormValue("question_type"),
        r.FormValue("category_id"),
        r.FormValue("question_img"),
        u_id,
        r.FormValue("respondent_img"),
        r.FormValue("sequence"),
        r.FormValue("mytext"),
        r.FormValue("spend"),
        r.FormValue("mychoice"),
        r.FormValue("explanation"),
        time.Now().Format("2006-01-02 15:04:05"))
    if err != nil {
        panic(err)
    }
    fmt.Fprint(w, `["1"]`)

}