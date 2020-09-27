package app

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// Rank page
func Rank(w http.ResponseWriter, r *http.Request) {
	uID := common.GetUser(w, r)
	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	type Usr struct {
		Sum      int
		UsrID    int
		UsrImg   string
		EtoColor string
		MeColor  string
	}

	type View struct {
		CacheV       string
		CategoryID   string
		CategoryName string
		CategoryTxt  template.HTML
		Usr          []Usr
	}
	var view View
	view.CacheV = common.CacheV
	cookie, err := r.Cookie("cat")
	if err != nil {
		log.Print("No cat Cookie: ", err)
		http.Redirect(w, r, "/", 302)
		return
	}
	t := time.Now()
	t1 := t.AddDate(0, -1, 0)
	time.Now().Format("2006-01-02 15:04:05")
	rows, err := db.Query(
		"SELECT SUM(correct), respondent_id "+
			"FROM h_answer                      "+
			"WHERE created_at > $1              "+
			"AND category_id = $2               "+
			"GROUP BY respondent_id             ",
		t1.Format("2006-01-02 15:04:05"), cookie.Value)
	if err != nil {
		log.Print(err)
	}
	var usr []Usr
	for rows.Next() {
		r := Usr{}
		if err := rows.Scan(&r.Sum, &r.UsrID); err != nil {
			log.Print(err)
		}
		eto := common.Eto(r.UsrID)
		r.UsrImg = eto[0]
		r.EtoColor = eto[1]
		r.MeColor = "#EEEEEE"
		if r.UsrID == uID {
			r.MeColor = "white"
		}
		usr = append(usr, r)
	}
	sort.Slice(usr, func(i, j int) bool { return usr[i].Sum > usr[j].Sum }) // DESC
	view.Usr = usr

	rows, err = db.Query("SELECT category_name FROM m_category_name WHERE category_id = $1", cookie.Value)
	if err != nil {
		log.Print(err)
	}
	var categoryName string
	for rows.Next() {
		if err := rows.Scan(&categoryName); err != nil {
			log.Print(err)
		}
	}

	view.CategoryID = cookie.Value
	view.CategoryName = categoryName
	tpl := template.Must(template.ParseFiles("html/rank.html"))
	tpl.Execute(w, view)

}
