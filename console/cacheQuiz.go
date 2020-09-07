package console

import (
  "fmt"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "../common"
  // "io/ioutil"
)

func CacheQuiz() {
  connStr := common.DB_CONNECT
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var query string
  query = "SELECT question_id FROM t_question  ORDER BY question_id"
  rows, err := db.Query(query)
  if err != nil {
      fmt.Println(err)
  }
  
  for rows.Next() {
    var question_id string
    if err := rows.Scan(&question_id); err != nil {
      fmt.Println(err)
    }
    resp, err := http.Get("https://shikaku.quigen.info/quiz/?q="+question_id)
    if err != nil {
      fmt.Println(err)
    }
    defer resp.Body.Close()
    fmt.Println(question_id)
    // byteArray, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println(string(byteArray)) // htmlをstringで取得
  }

}
