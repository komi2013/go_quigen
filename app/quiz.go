package app

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// Quiz page
func Quiz(w http.ResponseWriter, r *http.Request) {

	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	type BreadCrumb struct {
		Level        int
		CategoryID   int
		CategoryName string
	}
	type View struct {
		CacheV       string
		Q            common.TQuestion
		UpdatedAt    string
		Myphoto      string
		EtoColor     string
		Available    int
		BreadCrumb   []BreadCrumb
		Title        string
		Description  string
		OgImage      string
		Qtxt         template.HTML
		Choice0      template.HTML
		Choice1      template.HTML
		Choice2      template.HTML
		Choice3      template.HTML
		Qexplanation template.HTML
		// PleaseClick string
	}
	var view View
	view.CacheV = common.CacheV
	uID := common.GetUser(w, r)
	eto := common.Eto(uID)
	view.Myphoto = eto[0]
	view.EtoColor = eto[1]
	if _, err := strconv.Atoi(r.FormValue("q")); err != nil {
		log.Print(r.URL.Path, r.Form)
		http.NotFound(w, r)
		return
	}
	query := `SELECT question_id,question_txt,usr_id,usr_img,updated_at,choice_0,
        choice_1,choice_2,choice_3,reference,reference2,category_id,question_img,
        previous_question,next_question,sequence,sound,explanation,choice_type
        from t_question WHERE question_id = $1`
	rows, err := db.Query(query, r.FormValue("q"))
	if err != nil {
		fmt.Println(query)
		fmt.Println(err)
	}
	var question common.TQuestion
	nothing := true
	for rows.Next() {
		r := common.TQuestion{}
		if err := rows.Scan(&r.QuestionID, &r.QuestionTxt, &r.UsrID, &r.UsrImg, &r.UpdatedAt, &r.Choice0, &r.Choice1,
			&r.Choice2, &r.Choice3, &r.Reference, &r.Reference2, &r.CategoryID, &r.QuestionImg, &r.PreviousQuestion,
			&r.NextQuestion, &r.Sequence, &r.Sound, &r.Explanation, &r.ChoiceType); err != nil {
			log.Print(err)
		}
		view.UpdatedAt = r.UpdatedAt.Format("2006-01-02 15:04:05")
		if r.Sound != "" {
			view.Available = 0
		} else {
			view.Available = 1
		}
		question = r
		nothing = false
	}
	if nothing {
		fmt.Println("?q=" + r.FormValue("q"))
		http.NotFound(w, r)
		return
	}
	txt := []rune(question.QuestionTxt)
	txtLen := len(txt)
	titleLen := 32
	descLen := 200
	if txtLen < titleLen {
		titleLen = txtLen
		descLen = txtLen
	}
	if txtLen < descLen {
		descLen = txtLen
	}

	// fmt.Printf("descLen %#v\n", descLen)
	// r.QuestionTxt = string(txt[0:titleLen])
	view.Title = string(txt[0:titleLen])
	view.Description = "1." + question.Choice0 + ", 2." + question.Choice1 + ", 3." + question.Choice2 + ", 4." + question.Choice3 + ".." + string(txt[titleLen:descLen])
	if question.QuestionImg != "" {
		view.OgImage = question.QuestionImg
	} else {
		view.OgImage = "/favicon.ico"
	}
	// s := strings.Replace(question.QuestionTxt, "\n", "<br>", -1)
	// escape := template.HTML(strings.Replace(question.QuestionTxt, "\n", "<br>", -1))
	view.Qtxt = template.HTML(strings.Replace(question.QuestionTxt, "\n", "<br>", -1))
	view.Choice0 = template.HTML(strings.Replace(question.Choice0, "\n", "<br>", -1))
	view.Choice1 = template.HTML(strings.Replace(question.Choice1, "\n", "<br>", -1))
	view.Choice2 = template.HTML(strings.Replace(question.Choice2, "\n", "<br>", -1))
	view.Choice3 = template.HTML(strings.Replace(question.Choice3, "\n", "<br>", -1))
	view.Qexplanation = template.HTML(strings.Replace(question.Explanation, "\n", "<br>", -1))
	view.Q = question
	// fmt.Printf("question %#v\n", view.Q)
	convert := make(map[int]string)
	var tree [][2]int
	var x [2]int
	rows, err = db.Query("SELECT * FROM m_category_tree WHERE leaf_id = $1", question.CategoryID)
	if err != nil {
		log.Print(err)
	}
	whereIn := ""
	for rows.Next() {
		r := common.MCategoryTree{}
		if err := rows.Scan(&r.LeafID, &r.Level1, &r.Level2, &r.Level3, &r.Level4, &r.Level5, &r.Level6, &r.Level7, &r.Level8, &r.UpdatedAt); err != nil {
			log.Print(err)
		}
		whereIn = strconv.Itoa(r.Level1)
		x[0] = 1
		x[1] = r.Level1
		tree = append(tree, x)
		convert[r.Level1] = ""
		if r.Level2 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level2)
			x[0] = 2
			x[1] = r.Level2
			tree = append(tree, x)
			convert[r.Level2] = ""
		}
		if r.Level3 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level3)
			x[0] = 3
			x[1] = r.Level3
			tree = append(tree, x)
			convert[r.Level3] = ""
		}
		if r.Level4 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level4)
			x[0] = 4
			x[1] = r.Level4
			tree = append(tree, x)
			convert[r.Level4] = ""
		}
		if r.Level5 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level5)
			x[0] = 5
			x[1] = r.Level5
			tree = append(tree, x)
			convert[r.Level5] = ""
		}
		if r.Level6 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level6)
			x[0] = 6
			x[1] = r.Level6
			tree = append(tree, x)
			convert[r.Level6] = ""
		}
		if r.Level7 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level7)
			x[0] = 7
			x[1] = r.Level7
			tree = append(tree, x)
			convert[r.Level7] = ""
		}
		if r.Level8 > 0 {
			whereIn = whereIn + "," + strconv.Itoa(r.Level8)
			x[0] = 8
			x[1] = r.Level8
			tree = append(tree, x)
			convert[r.Level8] = ""
		}
	}
	if whereIn != "" {
		rows, err = db.Query("SELECT category_id, category_name FROM m_category_name WHERE category_id in (" + whereIn + ")")
		if err != nil {
			log.Print(err)
		}
		for rows.Next() {
			r := common.MCategoryName{}
			if err := rows.Scan(&r.CategoryID, &r.CategoryName); err != nil {
				log.Print(err)
			}
			convert[r.CategoryID] = r.CategoryName
			// categoryName = append(categoryName, r)
			// breadCrumb[r.CategoryID]["name"] = r.CategoryName
		}
		var breadCrumb []BreadCrumb
		for _, v := range tree {
			y := BreadCrumb{}
			y.Level = v[0]
			y.CategoryID = v[1]
			y.CategoryName = convert[v[1]]
			breadCrumb = append(breadCrumb, y)
		}

		sort.Slice(breadCrumb, func(i, j int) bool { return breadCrumb[i].Level < breadCrumb[j].Level }) // DESC
		// fmt.Printf("breadCrumb %#v\n", breadCrumb)
		view.BreadCrumb = breadCrumb
	}

	tpl := template.Must(template.ParseFiles("html/quiz.html"))
	tpl.Execute(w, view)
}
