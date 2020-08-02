package app

import (
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    // "fmt"
)

func Quiz(w http.ResponseWriter, r *http.Request) {

	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
    u_id := common.GetUser(w,r)
    type View struct {
        CacheV string
        CSRF string
        Q common.TQuestion
        UpdatedAt string
        Uid int
    }
    var view View
    view.CacheV = common.CACHE_V
    view.CSRF = common.MakeCSRF(w,r)
    view.Uid = u_id

    rows, err := db.Query("select question_id,question_txt,usr_id,usr_img,updated_at,choice_0,choice_1,choice_2,choice_3,reference,question_type,category_id,question_img,previous_question,next_question,sequence from t_question WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        panic(err)
    }
    var question common.TQuestion
    for rows.Next() {
        r := common.TQuestion{}
        if err := rows.Scan(&r.QuestionID, &r.QuestionTxt, &r.UsrID, &r.UsrImg, &r.UpdatedAt, &r.Choice0, &r.Choice1, &r.Choice2, &r.Choice3, &r.Reference, &r.QuestionType, &r.CategoryID, &r.QuestionImg, &r.PreviousQuestion, &r.NextQuestion, &r.Sequence); err != nil {
            log.Fatal(err)
        }
        view.UpdatedAt = r.UpdatedAt.Format("2006-01-02 15:04:05")
        question = r
    }
    view.Q = question
    tpl := template.Must(template.ParseFiles("html/quiz.html"))
    tpl.Execute(w, view)
}