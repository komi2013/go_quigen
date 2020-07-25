package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "encoding/json"
    "sort"
    "math"
)

func AnswerShow(w http.ResponseWriter, r *http.Request) {
    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	db, err := sql.Open("postgres", common.DB_CONNECT)
	if err != nil {
		log.Fatal(err)
	}
    rows, err := db.Query("SELECT question_id,question_txt,usr_id,usr_img,updated_at,choice_0,choice_1,choice_2,choice_3,reference,question_type,category_id,question_img,previous_question,next_question,sequence, mytext, mychoice FROM h_answer WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        panic(err)
    }
    var answer []common.HAnswer
    var count float64 = 0
    var correctNum float64 = 0
    for rows.Next() {
        r := common.HAnswer{}
        if err := rows.Scan(&r.QuestionID, &r.QuestionTxt, &r.UsrID, &r.UsrImg, &r.UpdatedAt, &r.Choice0, &r.Choice1, &r.Choice2, &r.Choice3, &r.Reference, &r.QuestionType, &r.CategoryID, &r.QuestionImg, &r.PreviousQuestion, &r.NextQuestion, &r.Sequence, &r.Mytext, &r.Mychoice); err != nil {
            log.Fatal(err)
        }
        answer = append(answer,r)
        if r.QuestionTxt == r.Mytext || r.Mychoice == 0 {
            correctNum++
        }
        count++
    }
    answerRatio := math.Round(correctNum / count * 100)
    answerCalc := [2]float64{0,0}
    if count > 0 {
        answerCalc[0] = answerRatio
        answerCalc[1] = count
    }
    
    // fmt.Printf("%#v\n", answerCalc)
    sort.Slice(answer, func(i, j int) bool { return answer[i].CreatedAt.After(answer[j].CreatedAt) }) // DESC

    var correctUsr [16][2]string
    var incorrectUsr [16][2]string

    fmt.Printf("%#v\n", answer)

    coU := 0
    inU := 0
    for _, d := range answer {
        if d.UsrImg == "" {
            eto := common.Eto(d.UsrID)
            d.UsrImg = eto[0]
            d.EtoColor = eto[1]
        }
        if (d.QuestionTxt == d.Mytext || d.Mychoice == 0) && coU < 16 {
            correctUsr[coU][0] = d.UsrImg
            correctUsr[coU][1] = "background-color:#" + d.EtoColor
            coU++
        } else if (d.QuestionTxt != d.Mytext || d.Mychoice != 0) && inU < 16 {
            incorrectUsr[inU][0] = d.UsrImg
            incorrectUsr[inU][1] = "background-color:#" + d.EtoColor
            inU++
        }
        // fmt.Printf("index: %d, name: %s\n", i, s)
    }
    
    

    rows, err = db.Query("SELECT comment_txt,usr_id,question_id,usr_img,created_at FROM t_comment WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        panic(err)
    }
    var comment []common.TComment
    for rows.Next() {
        r := common.TComment{}
        if err := rows.Scan(&r.CommentTxt, &r.UsrID, &r.QuestionID, &r.UsrImg, &r.CreatedAt); err != nil {
            log.Fatal(err)
        }
        if r.UsrImg == "" {
            eto := common.Eto(r.UsrID)
            r.UsrImg = eto[0]
            r.EtoColor = "background-color:#" + eto[1]
        }
        comment = append(comment,r)
    }

    sort.Slice(comment, func(i, j int) bool { return comment[i].CreatedAt.After(comment[j].CreatedAt) }) // DESC
    // sort.Sort(comment)
    // fmt.Printf("%#v\n", comment)
    type Arr struct {
        Status       int
        CorrectUsr   [16][2]string
        IncorrectUsr [16][2]string
        AnswerCalc   [2]float64
        Comment      []common.TComment
    }
    var arr *Arr
    arr = &Arr{
        Status: 1,
        CorrectUsr: correctUsr,
        IncorrectUsr: incorrectUsr,
        AnswerCalc: answerCalc,
        Comment: comment,
    }
    res, _ := json.Marshal(arr)
    // fmt.Printf("%#v\n", res)
    fmt.Fprint(w, string(res))
}

// func Slice(slice interface{}, less func(i, j int) bool)