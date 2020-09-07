package main

import (
    "fmt"
    // "html/template"
    "log"
    "net/http"
    // "time"
    "./app"
    "./common"
    "./console"
)

func main() {
    console.CacheQuiz()
    http.HandleFunc("/advertisement/", app.Advertisement)
    http.HandleFunc("/Answer/", app.Answer)
    http.HandleFunc("/AnswerShow/", app.AnswerShow)
    http.HandleFunc("/category/", app.Category)
    http.HandleFunc("/CommentAdd/", app.CommentAdd)
    http.HandleFunc("/generate/", app.Generate)
    http.HandleFunc("/GenerateQuiz/", app.GenerateQuiz)
    http.HandleFunc("/htm/", app.Htm)
    http.HandleFunc("/message/", app.Message)
    http.HandleFunc("/MessageAdd/", app.MessageAdd)
    http.HandleFunc("/mydata/", app.Mydata)
    http.HandleFunc("/NewQuestionShow/", app.NewQuestionShow)
    http.HandleFunc("/quiz/", app.Quiz)
    http.HandleFunc("/rank/", app.Rank)
    http.HandleFunc("/search/", app.Search)


    http.HandleFunc("/", app.Top)
    fmt.Println("starting.." + common.CACHE_V)
    fmt.Println(common.DB_CONNECT)

    log.Fatal(http.ListenAndServe(common.GO_PORT, nil))
}
