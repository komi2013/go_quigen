package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "encoding/json"
)

func NewQuestionShow(w http.ResponseWriter, r *http.Request) {
    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	db, err := sql.Open("postgres", common.DB_CONNECT)
	if err != nil {
		log.Print(err)
    }
    type Question struct {
        QuestionID       int
        QuestionTxt   string
    }
    rows, err := db.Query("SELECT question_id,question_txt FROM t_question ORDER BY updated_at DESC LIMIT 10")
    if err != nil {
        log.Print(err)
    }
    var question []Question
    for rows.Next() {
        r := Question{}
        if err := rows.Scan(&r.QuestionID, &r.QuestionTxt); err != nil {
            log.Print(err)
        }
        txtLen := 100
        s := []rune(r.QuestionTxt)
        if len(s) < 100 {
            txtLen = len(s)
        }
        fmt.Printf("s %#v\n", string(s[0:txtLen]))
        r.QuestionTxt = string(s[0:txtLen])
        question = append(question, r)
    }
    type Res struct {
        Status       int
        Question   []Question
    }
    var res Res
    res.Status = 1
    res.Question = question
    response, _ := json.Marshal(res)
    fmt.Fprint(w, string(response))
}