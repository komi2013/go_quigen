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
		log.Print(err)
    }
    defer db.Close()
    rows, err := db.Query("SELECT respondent_id ,respondent_img ,correct ,created_at FROM h_answer WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        log.Print(err)
    }
    var answer []common.HAnswer
    var count float64 = 0
    var correctNum float64 = 0
    for rows.Next() {
        r := common.HAnswer{}
        if err := rows.Scan(&r.RespondentID, &r.RespondentImg, &r.Correct, &r.CreatedAt); err != nil {
            log.Print(err)
        }
        answer = append(answer,r)
        correctNum = correctNum + r.Correct
        count++
    }
    answerRatio := math.Round(correctNum / count * 100)
    answerCalc := [2]float64{0,0}
    if count > 0 {
        answerCalc[0] = answerRatio
        answerCalc[1] = count
    }
    
    sort.Slice(answer, func(i, j int) bool { return answer[i].CreatedAt.After(answer[j].CreatedAt) }) // DESC

    var correctUsr [16][2]string
    var incorrectUsr [16][2]string

    coU := 0
    inU := 0
    for _, d := range answer {
        if d.RespondentImg == "" {
            eto := common.Eto(d.RespondentID)
            d.RespondentImg = eto[0]
            d.EtoColor = eto[1]
        }
        if d.Correct == 1 && coU < 16 {
            correctUsr[coU][0] = d.RespondentImg
            correctUsr[coU][1] = "background-color:#" + d.EtoColor
            coU++
        } else if d.Correct == 0 && inU < 16 {
            incorrectUsr[inU][0] = d.RespondentImg
            incorrectUsr[inU][1] = "background-color:#" + d.EtoColor
            inU++
        }
    }

    rows, err = db.Query("SELECT comment_txt,usr_id,question_id,usr_img,created_at FROM t_comment WHERE question_id = $1", r.FormValue("q"))
    if err != nil {
        log.Print(err)
    }
    var comment []common.TComment
    for rows.Next() {
        r := common.TComment{}
        if err := rows.Scan(&r.CommentTxt, &r.UsrID, &r.QuestionID, &r.UsrImg, &r.CreatedAt); err != nil {
            log.Print(err)
        }
        if r.UsrImg == "" {
            eto := common.Eto(r.UsrID)
            r.UsrImg = eto[0]
            r.EtoColor = "background-color:#" + eto[1]
        }
        comment = append(comment,r)
    }

    sort.Slice(comment, func(i, j int) bool { return comment[i].CreatedAt.After(comment[j].CreatedAt) }) // DESC
    type Arr struct {
        Status       int
        CorrectUsr   [16][2]string
        IncorrectUsr [16][2]string
        AnswerCalc   [2]float64
        Comment      []common.TComment
        CSRF         string
    }
    var arr Arr
    arr.Status = 1
    arr.CorrectUsr = correctUsr
    arr.IncorrectUsr = incorrectUsr
    arr.AnswerCalc = answerCalc
    arr.Comment = comment
    arr.CSRF = common.MakeCSRF(w,r)
    res, _ := json.Marshal(arr)
    fmt.Fprint(w, string(res))
}