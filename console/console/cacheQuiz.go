package console

import (
	"fmt"
	"log"
	"net/http"

	// "net/http"
	"database/sql"

	"../../common"
	_ "github.com/lib/pq" // this driver for postgres

	// "io/ioutil"
	// "os/exec"
	"time"
)

func CacheQuiz() {
	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var query string
	query = "SELECT question_id FROM c_resource WHERE category_id = -1"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	questionID := "0"
	for rows.Next() {
		if err := rows.Scan(&questionID); err != nil {
			fmt.Println(err)
		}
	}
	query = "SELECT question_id FROM t_question WHERE question_id > " + questionID + " ORDER BY question_id LIMIT 9500"
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(query)
	// timeout := time.Duration(60 * time.Second)
	// client := http.Client{
	//   Timeout: timeout,
	// }
	var q []string
	for rows.Next() {
		if err := rows.Scan(&questionID); err != nil {
			fmt.Println(err)
		}
		// time.Sleep(40 * time.Millisecond)
		// args := []string{"--spider", "https://english.quigen.info/quiz/?q=" + questionID}
		// err := exec.Command("wget", args...).Start()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(questionID)
		q = append(q, questionID)
		// questionID = questionID
		// byteArray, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(byteArray)) // htmlをstringで取得
	}
	parallel(q)
	_, err = db.Exec(`UPDATE c_resource SET question_id = $1 WHERE category_id = -1`,
		questionID)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(questionID)
}
func parallel(number []string) {
	for _, n := range number {
		time.Sleep(40 * time.Millisecond)
		go func(n string) {
			fmt.Println("parallel: ", n)
			resp, err := http.Get("https://shikaku.quigen.info/quiz/?q=" + n)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
		}(n)
	}
	time.Sleep(5 * time.Second)
}
