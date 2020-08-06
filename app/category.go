package app

import (
    "html/template"
    // "database/sql"
    // _ "github.com/lib/pq"
    // "log"
    "net/http"
    "../common"
    "fmt"
    // "strconv"
    // "sort"
)

func Category(w http.ResponseWriter, r *http.Request) {

	// connStr := common.DB_CONNECT
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Print(err)
	// }
    type BreadCrumb struct {
        Level int
        CategoryID int
        CategoryName string
    }
    type CategoryList struct {
        Level int
        CategoryID int
        CategoryName string
        CategoryQuestion []common.MCategoryQuestion
    }
    type View struct {
        CacheV string
        BreadCrumb []BreadCrumb
        CategoryList []CategoryList
    }
    var view View
    view.CacheV = common.CACHE_V


    var categoryQuestion []common.MCategoryQuestion
    x := common.MCategoryQuestion{}
    x.QuestionID = 1
    x.CategoryID = 1
    categoryQuestion = append(categoryQuestion, x)
    categoryQuestion = append(categoryQuestion, x)
    var categoryList []CategoryList
    y := CategoryList{}
    y.Level = 2
    y.CategoryID = 2
    y.CategoryName = "name"
    y.CategoryQuestion = categoryQuestion
    categoryList = append(categoryList, y)
    categoryList = append(categoryList, y)
    view.CategoryList = categoryList
    fmt.Printf("categoryList %#v\n", categoryList)
    tpl := template.Must(template.ParseFiles("html/category.html"))
    tpl.Execute(w, view)
}