package app

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// GenerateQuiz insert t_question
func GenerateQuiz(w http.ResponseWriter, r *http.Request) {

	ok := common.CheckCSRF(r, r.FormValue("csrf"))
	if !ok {
		fmt.Fprint(w, `["2","CSRF Error"]`)
		return
	}
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
	questionID := r.FormValue("question_id")
	if questionID != "" {
		query := `SELECT usr_id FROM t_question WHERE question_id = $1`
		rows, err := db.Query(query, r.FormValue("question_id"))
		if err != nil {
			log.Print(query)
			log.Print(err)
		}
		usrID := 0
		for rows.Next() {
			if err := rows.Scan(&usrID); err != nil {
				log.Print(err)
			}
		}
		if uID != usrID {
			log.Print("what are you trying to do?")
			return
		}
		_, err = db.Exec(`UPDATE t_question SET 
			question_txt = $1,
			choice_0     = $2,
			choice_1     = $3,
			choice_2     = $4,
			choice_3     = $5,
			category_id  = $6,
			reference    = $7,
			reference2   = $8,
			explanation  = $9
			WHERE question_id = $10`,
			r.FormValue("q_txt"),
			r.FormValue("choice_0"),
			r.FormValue("choice_1"),
			r.FormValue("choice_2"),
			r.FormValue("choice_3"),
			r.FormValue("category"),
			r.FormValue("reference"),
			r.FormValue("reference2"),
			r.FormValue("explanation"),
			questionID,
		)
		if err != nil {
			log.Print(questionID)
			log.Print(err)
			return
		}
	} else {
		rows, err := db.Query("SELECT NEXTVAL('t_question_question_id_seq')")
		if err != nil {
			log.Print(err)
		}
		var nextval int
		for rows.Next() {
			if err := rows.Scan(&nextval); err != nil {
				log.Print(err)
			}
		}
		questionID = strconv.Itoa(nextval)
		img := ""
		if r.FormValue("img") != "no" {
			img = strings.Replace(r.FormValue("img"), "data:image/png;base64,", "", 1)
			data, _ := base64.StdEncoding.DecodeString(img)
			dir := "/data/img_question/" + time.Now().Format("20060102") + "/"
			os.Mkdir("public"+dir, 0777)
			file, _ := os.Create("public" + dir + questionID + ".png")
			defer file.Close()
			file.Write(data)
		}

		_, err = db.Exec(`INSERT INTO t_question(
			question_id,
			question_txt,
			choice_0,
			choice_1,
			choice_2,
			choice_3,
			question_img,
			category_id,
			reference,
			reference2,
			explanation,
			usr_id,
			usr_img
			) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`,
			questionID,
			r.FormValue("q_txt"),
			r.FormValue("choice_0"),
			r.FormValue("choice_1"),
			r.FormValue("choice_2"),
			r.FormValue("choice_3"),
			img,
			r.FormValue("category"),
			r.FormValue("reference"),
			r.FormValue("reference2"),
			r.FormValue("explanation"),
			uID,
			usrImg)
		if err != nil {
			log.Print(err)
			return
		}
	}
	fmt.Fprint(w, `[1,`+questionID+`]`)

}
