package app

import (
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "../common"
    "fmt"
    "strconv"
    "sort"
    "strings"
    // "encoding/json"
)

func Category(w http.ResponseWriter, r *http.Request) {

	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
    type BreadCrumb struct {
        Level int
        CategoryID int
        CategoryName string
    }
    type CategoryQuestion struct {
        QuestionID string
        QuestionTitle string
    }
    type CategoryList struct {
        Level string
        CategoryID int
        CategoryName string
        CategoryQuestion []CategoryQuestion
    }

    type View struct {
        CacheV string
        CSRF string
        BreadCrumb []BreadCrumb
        CategoryList []CategoryList
        CategoryName string
        CategoryDescription string
        CategoryTxt template.HTML
    }
    var view View
    view.CacheV = common.CACHE_V
    view.CSRF = ""

    u := strings.Split(r.URL.Path, "/")
    fmt.Printf("url %#v\n", u[2])
    if u[2] != "1" && u[2] != "2" && u[2] != "3" && u[2] != "4" && u[2] != "5" &&
    u[2] != "6" && u[2] != "7" && u[2] != "8" {
        return
    }
    rows, err := db.Query("SELECT * FROM m_category_tree WHERE level_"+u[2]+" = $1", u[3])
    if err != nil {
        log.Print(err)
    }
    var tree [][2]int
    var x [2]int
    convert := make(map[int]string)
    treeList := map[int]map[string]string{}
    whereIn := u[3]
    d := common.MCategoryTree{}
    for rows.Next() {
        if err := rows.Scan(&d.LeafID,&d.Level1,&d.Level2,&d.Level3,&d.Level4,&d.Level5,&d.Level6,&d.Level7,&d.Level8,&d.UpdatedAt); err != nil {
            log.Print(err)
        }
        list := map[string]string{}
        switch u[2] {
        case "1":
            list["level"] = "2"
            list["category_name"] = ""
            treeList[d.Level2] = list
        case "2":
            list["level"] = "3"
            list["category_name"] = ""
            treeList[d.Level3] = list
        case "3":
            list["level"] = "4"
            list["category_name"] = ""
            treeList[d.Level4] = list
        case "4":
            list["level"] = "5"
            list["category_name"] = ""
            treeList[d.Level5] = list
        case "5":
            list["level"] = "6"
            list["category_name"] = ""
            treeList[d.Level6] = list
        }
    }
    if(d.Level1 > 0 && u[2] > "1"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level1)
        x[0] = 1
        x[1] = d.Level1
        tree = append(tree, x)
        convert[d.Level1] = ""
    }
    if(d.Level2 > 0 && u[2] > "2"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level2)
        x[0] = 2
        x[1] = d.Level2
        tree = append(tree, x)
        convert[d.Level2] = ""
    }
    if(d.Level3 > 0 && u[2] > "3"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level3)
        x[0] = 3
        x[1] = d.Level3
        tree = append(tree, x)
        convert[d.Level3] = ""
    }
    if(d.Level4 > 0 && u[2] > "4"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level4)
        x[0] = 4
        x[1] = d.Level4
        tree = append(tree, x)
        convert[d.Level4] = ""
    }
    if(d.Level5 > 0 && u[2] > "5"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level5)
        x[0] = 5
        x[1] = d.Level5
        tree = append(tree, x)
        convert[d.Level5] = ""
    }
    if(d.Level6 > 0 && u[2] > "6"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level6)
        x[0] = 6
        x[1] = d.Level6
        tree = append(tree, x)
        convert[d.Level6] = ""
    }
    if(d.Level7 > 0 && u[2] > "7"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level7)
        x[0] = 7
        x[1] = d.Level7
        tree = append(tree, x)
        convert[d.Level7] = ""
    }
    if(d.Level8 > 0 && u[2] > "8"){
        whereIn = whereIn + "," + strconv.Itoa(d.Level8)
        x[0] = 8
        x[1] = d.Level8
        tree = append(tree, x)
        convert[d.Level8] = ""
    }
    whereIn2 := ""

    for i, _ := range treeList {
        whereIn2 = whereIn2 + "," + strconv.Itoa(i)
    }

    rows, err = db.Query("SELECT category_id, category_name, category_description FROM m_category_name WHERE category_id in ("+whereIn+whereIn2+")")
    if err != nil {
        log.Print(err)
    }
    for rows.Next() {
        r := common.MCategoryName{}
        if err := rows.Scan(&r.CategoryID,&r.CategoryName,&r.CategoryDescription); err != nil {
            log.Print(err)
        }
        if _, ok := convert[r.CategoryID]; ok {
            convert[r.CategoryID] = r.CategoryName
        }
        if _, ok := treeList[r.CategoryID]; ok {
            treeList[r.CategoryID]["category_name"] = r.CategoryName
        }
        if u[3] == strconv.Itoa(r.CategoryID) {
            view.CategoryName = r.CategoryName
            view.CategoryDescription = r.CategoryDescription
            view.CategoryTxt = template.HTML(strings.Replace(r.CategoryDescription, "\n", "<br>", -1))
        }
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
    fmt.Printf("breadCrumb %#v\n", breadCrumb)
    view.BreadCrumb = breadCrumb
    
    var categoryQuestionPre []common.MCategoryQuestion
    if whereIn2 != "" {
        rows, err = db.Query("SELECT question_id, category_id, question_title FROM m_category_question WHERE category_id in ("+whereIn2[1:]+")")
        if err != nil {
            log.Print(err)
        }
        for rows.Next() {  //5,13,16,15,14
            r := common.MCategoryQuestion{}
            if err := rows.Scan(&r.QuestionID,&r.CategoryID,&r.QuestionTitle); err != nil {
                log.Print(err)
            }
            categoryQuestionPre = append(categoryQuestionPre, r)
        }
    }
    var categoryList []CategoryList
    for i, v := range treeList {
        y := CategoryList{}
        y.Level = v["level"]
        y.CategoryID = i
        y.CategoryName = v["category_name"]
        var categoryQuestion []CategoryQuestion
        for _, v2 := range categoryQuestionPre {
            if y.CategoryID == v2.CategoryID {
                y2 := CategoryQuestion{}
                y2.QuestionID = strconv.Itoa(v2.QuestionID)
                y2.QuestionTitle = v2.QuestionTitle
                categoryQuestion = append(categoryQuestion, y2)
            }
        }
        y.CategoryQuestion = categoryQuestion
        categoryList = append(categoryList, y)
    }

    // res, _ := json.Marshal(categoryList)
    
    // var categoryQuestion []common.MCategoryQuestion
    // x = common.MCategoryQuestion{}
    // x.QuestionID = 1
    // x.CategoryID = 1
    // categoryQuestion = append(categoryQuestion, x)
    // categoryQuestion = append(categoryQuestion, x)
    // var categoryList []CategoryList
    // y := CategoryList{}
    // y.Level = 2
    // y.CategoryID = 2
    // y.CategoryName = "name"
    // y.CategoryQuestion = categoryQuestion
    // categoryList = append(categoryList, y)
    // categoryList = append(categoryList, y)
    view.CategoryList = categoryList
    // fmt.Printf("categoryList %#v\n", categoryList)
    tpl := template.Must(template.ParseFiles("html/category.html"))
    tpl.Execute(w, view)
}