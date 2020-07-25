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
    http.HandleFunc("/AnswerShow/", app.AnswerShow)
    http.HandleFunc("/CommentAdd/", app.CommentAdd)
    http.HandleFunc("/generate/", app.Generate)
    http.HandleFunc("/GenerateQuiz/", app.GenerateQuiz)
    http.HandleFunc("/crypttry/", app.CryptTry)
    http.HandleFunc("/SqlTest", app.SqlTest)
    http.HandleFunc("/htm/", app.Htm)
    http.HandleFunc("/quiz/", app.Quiz)
    
    fmt.Println("浄化")
    log.Fatal(http.ListenAndServe(":9001", nil))
}
