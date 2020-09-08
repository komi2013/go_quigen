package console

import (
  "fmt"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "../common"
  // "io/ioutil"
  "time"
)

func CacheQuiz() {
  connStr := common.DB_CONNECT
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var query string
  query = "SELECT question_id FROM t_question WHERE question_id > 7752 ORDER BY question_id"
  rows, err := db.Query(query)
  if err != nil {
      fmt.Println(err)
  }
  timeout := time.Duration(60 * time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  for rows.Next() {
    var question_id string
    if err := rows.Scan(&question_id); err != nil {
      fmt.Println(err)
    }
    resp, err := client.Get("https://shikaku.quigen.info/quiz/?q="+question_id)
    if err != nil {
      fmt.Println(err)
    }
    defer resp.Body.Close()
    fmt.Println(question_id)
    // byteArray, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println(string(byteArray)) // htmlをstringで取得
  }

}
