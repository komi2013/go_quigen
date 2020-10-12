package app

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// Generate page
func Generate(w http.ResponseWriter, r *http.Request) {

	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	uID := common.GetUser(w, r)
	query := `SELECT usr_img FROM t_usr WHERE usr_id = $1`
	rows, err := db.Query(query, uID)
	if err != nil {
		log.Print(err)
	}
	usrImg := ""
	is := false
	for rows.Next() {
		if err := rows.Scan(&usrImg); err != nil {
			log.Print(err)
		}
		is = true
	}
	if !is {
		log.Print("you are not registered yet")
		return
	}
	query = `SELECT category_id, category_name FROM m_category_name
             WHERE category_id in (
                 SELECT leaf_id FROM m_category_tree GROUP BY leaf_id
             ) ORDER BY category_id`
	rows, err = db.Query(query)
	if err != nil {
		log.Print(err)
	}
	category := map[int]string{}
	for rows.Next() {
		categoryID := 0
		categoryName := ""
		if err := rows.Scan(&categoryID, &categoryName); err != nil {
			log.Print(err)
		}
		category[categoryID] = categoryName
	}
	res, _ := json.Marshal(category)
	questionTxt := ""
	choice0 := ""
	choice1 := ""
	choice2 := ""
	choice3 := ""
	reference := ""
	reference2 := ""
	explanation := ""
	questionImg := ""
	questionID := ""

	if r.FormValue("q") != "" {
		query = `SELECT question_txt,choice_0,choice_1 ,choice_2 ,choice_3 ,
        reference ,reference2,explanation,question_img, question_id
        FROM t_question WHERE question_id = $1`
		rows, err := db.Query(query, r.FormValue("q"))
		if err != nil {
			log.Print(err)
		}
		for rows.Next() {
			if err := rows.Scan(&questionTxt, &choice0, &choice1, &choice2, &choice3,
				&reference, &reference2, &explanation, &questionImg, &questionID); err != nil {
				log.Print(err)
			}
		}
	}
	// log.Print(choice0)
	// fmt.Printf("%#v\n", string(res))
	m := map[string]string{
		"CacheV":       common.CacheV,
		"CSRF":         common.MakeCSRF(w, r),
		"Category":     string(res),
		"question_txt": questionTxt,
		"choice_0":     choice0,
		"choice_1":     choice1,
		"choice_2":     choice2,
		"choice_3":     choice3,
		"reference":    reference,
		"reference2":   reference2,
		"explanation":  explanation,
		"question_img": questionImg,
		"question_id":  questionID,
	}
	tpl := template.Must(template.ParseFiles("html/generate.html"))
	tpl.Execute(w, m)
}
