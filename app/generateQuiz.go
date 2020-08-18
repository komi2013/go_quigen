package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    // "encoding/json"
    "encoding/base64"
    "os"
    "strings"
    "time"
    "strconv"
    // "../table"
)

func GenerateQuiz(w http.ResponseWriter, r *http.Request) {
    usr_id := 1
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
	rows, err := db.Query("SELECT NEXTVAL('t_question_question_id_seq')")
    if err != nil {
        log.Print(err)
    }
    var nextval int
    for rows.Next() {
        if err := rows.Scan(&nextval); err != nil {
            log.Print(err)
        }
    }
    question_id := nextval
    img := ""
    if r.FormValue("img") != "no" {
        img = strings.Replace(r.FormValue("img"), "data:image/png;base64,", "", 1)
        data, _ := base64.StdEncoding.DecodeString(img)
        dir := "/data/img_question/" + time.Now().Format("20060102") + "/"
        os.Mkdir("public" + dir, 0777)
        file, _ := os.Create("public" + dir + strconv.Itoa(question_id) + ".png")
        defer file.Close()
        file.Write(data)
    }

    _, err = db.Exec(`INSERT INTO t_question(
        question_id,
        question_txt,
        choice_0,
        choice_1,
        choice_2,
        choice_3,
        question_img,
        category_id,
        reference,
        usr_id,
        usr_img
        ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`, 
        question_id,
        r.FormValue("q_txt"),
        r.FormValue("choice_0"),
        r.FormValue("choice_1"),
        r.FormValue("choice_2"),
        r.FormValue("choice_3"),
        img,
        0,
        r.FormValue("reference"),
        usr_id,
        r.FormValue("myphoto"))
    if err != nil {
        log.Print(err)
    }
    fmt.Fprint(w, `["1"]`)

}