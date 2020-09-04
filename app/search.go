package app

import (
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "strconv"
    "../common"
    // "net/url"
    // "strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
  }
  defer db.Close()
  rows, err := db.Query("SELECT category_id FROM m_category_name WHERE category_name = $1", r.FormValue("tag"))
  if err != nil {
      log.Print(err)
  }
  for rows.Next() {
    var category_id int
    if err := rows.Scan(&category_id); err != nil {
        log.Print(err)
    }
    http.Redirect(w, r, "/category/1/"+strconv.Itoa(category_id)+"/", 301)
    return
  }
}
