package app

import (
    "fmt"
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    // "time"
    "../common"
)

func Quiz(w http.ResponseWriter, r *http.Request) {

	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

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
        // question = append(question,r)
        question = r
    }
    fmt.Printf("%v\n", question)

    type View struct {
        CacheV string
        CSRF string
        Q common.TQuestion
    }
    view := &View{
        CacheV : common.CACHE_V,
        CSRF : common.MakeCSRF(w,r),
        Q : question,
    }
    cookie, err := r.Cookie("xsssr")
    value := ""
    if err == nil {
        value = cookie.Value
	}
    // fmt.Println("koutei")
    fmt.Printf("%#v\n", value)
    tpl := template.Must(template.ParseFiles("html/quiz.html"))
    tpl.Execute(w, view)

}

// func GetTQuestionByPk(db Queryer) (*TQuestion, error) {
// 	var r TQuestion
// 	err := db.QueryRow(
// 		`SELECT question_id, question_txt, usr_id, usr_img, created_at, updated_at, choice_0, choice_1, choice_2, choice_3, reference, question_type, category_id, question_img, previous_question, next_question, sequence FROM t_question WHERE `,
// 	).Scan(&r.QuestionID, &r.QuestionTxt, &r.UsrID, &r.UsrImg, &r.CreatedAt, &r.UpdatedAt, &r.Choice0, &r.Choice1, &r.Choice2, &r.Choice3, &r.Reference, &r.QuestionType, &r.CategoryID, &r.QuestionImg, &r.PreviousQuestion, &r.NextQuestion, &r.Sequence)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to select t_question")
// 	}
// 	return &r, nil
// }