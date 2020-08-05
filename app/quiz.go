package app

import (
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "fmt"
    "strconv"
    "encoding/json"
)

func Quiz(w http.ResponseWriter, r *http.Request) {

	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
    u_id := common.GetUser(w,r)
    type View struct {
        CacheV string
        CSRF string
        Q common.TQuestion
        UpdatedAt string
        Myphoto string
        EtoColor string
        Available int
        BreadCrumb string
        // PleaseClick string
    }
    var view View
    view.CacheV = common.CACHE_V
    view.CSRF = common.MakeCSRF(w,r)
    // view.PleaseClick = common.PLEASE_CLICK

    eto := common.Eto(u_id)
    view.Myphoto = eto[0]
    view.EtoColor = eto[1]

    rows, err := db.Query("SELECT question_id,question_txt,usr_id,usr_img,updated_at,choice_0,choice_1,choice_2,choice_3,reference,question_type,category_id,question_img,previous_question,next_question,sequence,sound,explanation from t_question WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        log.Print(err)
    }
    var question common.TQuestion
    for rows.Next() {
        r := common.TQuestion{}
        if err := rows.Scan(&r.QuestionID, &r.QuestionTxt, &r.UsrID, &r.UsrImg, &r.UpdatedAt, &r.Choice0, &r.Choice1, &r.Choice2, &r.Choice3, &r.Reference, &r.QuestionType, &r.CategoryID, &r.QuestionImg, &r.PreviousQuestion, &r.NextQuestion, &r.Sequence, &r.Sound, &r.Explanation); err != nil {
            log.Print(err)
        }
        view.UpdatedAt = r.UpdatedAt.Format("2006-01-02 15:04:05")
        if r.Sound != "" {
            view.Available = 0
        }else{
            view.Available = 1
        }
        question = r
    }
    view.Q = question
    breadCrumb := make(map[int]map[int]string)
    rows, err = db.Query("SELECT * FROM m_category_tree WHERE leaf_id = $1", question.CategoryID)
    if err != nil {
        log.Print(err)
    }
    var whereIn string
    for rows.Next() {
        r := common.MCategoryTree{}
        if err := rows.Scan(&r.LeafID,&r.Level1,&r.Level2,&r.Level3,&r.Level4,&r.Level5,&r.Level6,&r.Level7,&r.Level8,&r.UpdatedAt); err != nil {
            log.Print(err)
        }
        whereIn = strconv.Itoa(r.Level1)
        breadCrumb[r.Level1] = make(map[int]string)
        if(r.Level2 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level2)
            breadCrumb[r.Level2] = make(map[int]string)
        }
        if(r.Level3 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level3)
            breadCrumb[r.Level3] = make(map[int]string)
        }
        if(r.Level4 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level4)
            breadCrumb[r.Level4] = make(map[int]string)
        }
        if(r.Level5 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level5)
            breadCrumb[r.Level5] = make(map[int]string)
        }
        if(r.Level6 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level6)
            breadCrumb[r.Level6] = make(map[int]string)
        }
        if(r.Level7 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level7)
            breadCrumb[r.Level7] = make(map[int]string)
        }
        if(r.Level8 > 0){
            whereIn = whereIn + "," + strconv.Itoa(r.Level8)
            breadCrumb[r.Level8] = make(map[int]string)
        }
    }
    rows, err = db.Query("SELECT category_id, category_name FROM m_category_name WHERE category_id in ("+whereIn+")")
    if err != nil {
        log.Print(err)
    }
    for rows.Next() {
        r := common.MCategoryName{}
        if err := rows.Scan(&r.CategoryID,&r.CategoryName); err != nil {
            log.Print(err)
        }
        // categoryName = append(categoryName, r)
        breadCrumb[r.CategoryID][0] = r.CategoryName
    }
    res, _ := json.Marshal(breadCrumb)
    fmt.Printf("breadCrumb %#v\n", string(res))
    view.BreadCrumb = string(res)
    tpl := template.Must(template.ParseFiles("html/quiz.html"))
    tpl.Execute(w, view)
}