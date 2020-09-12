package main

import (
  "fmt"
  // "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "../common"
  // "io/ioutil"
  "time"
  "os/exec"
)

func main() {
  connStr := common.DB_CONNECT
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var query string
  query = "SELECT question_id FROM t_question WHERE question_id > 313331 ORDER BY question_id LIMIT 9500"
  rows, err := db.Query(query)
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(query)
  // timeout := time.Duration(60 * time.Second)
  // client := http.Client{
  //   Timeout: timeout,
  // }
  for rows.Next() {
    var question_id string
    if err := rows.Scan(&question_id); err != nil {
      fmt.Println(err)
    }
    time.Sleep(40 * time.Millisecond)
    args := []string{"--spider", "https://english.quigen.info/quiz/?q="+question_id}
    err := exec.Command("wget", args...).Start()
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(question_id)
    // byteArray, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println(string(byteArray)) // htmlをstringで取得
  }

}
