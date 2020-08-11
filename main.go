package main

import (
    "fmt"
    // "html/template"
    "log"
    "net/http"
    // "time"
    "./app"
)

func main() {
    http.HandleFunc("/advertisement/", app.Advertisement)
    http.HandleFunc("/Answer/", app.Answer)
    http.HandleFunc("/AnswerShow/", app.AnswerShow)
    http.HandleFunc("/category/", app.Category)
    http.HandleFunc("/CommentAdd/", app.CommentAdd)
    http.HandleFunc("/generate/", app.Generate)
    http.HandleFunc("/GenerateQuiz/", app.GenerateQuiz)
    http.HandleFunc("/htm/", app.Htm)
    http.HandleFunc("/mydata/", app.Mydata)
    http.HandleFunc("/NewQuestionShow/", app.NewQuestionShow)
    http.HandleFunc("/quiz/", app.Quiz)
    http.HandleFunc("/rank/", app.Rank)
    http.HandleFunc("/", app.Top)
    fmt.Println("starting..")
    log.Fatal(http.ListenAndServe(":9001", nil))
}
