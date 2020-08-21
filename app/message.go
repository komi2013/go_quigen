package app

import (
    "net/http"
    "html/template"
    "log"
    "time"
    "../common"
    "database/sql"
    _ "github.com/lib/pq"
    // "sort"
    "fmt"
    "strconv"
    "encoding/json"
    "html"
    "strings"
)

func Message(w http.ResponseWriter, r *http.Request) {
	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
  }
  defer db.Close()
  type Msg struct {
    Txt       string
    SenderFlg string
    CreatedAt time.Time
  }
  type MsgList struct {
    OpponentID     int
    OpponentImg    string
    OpponentColor  string
    Msg          []Msg
  }
  type View struct {
    CacheV       string
    CSRF         string
    Myphoto      string
    EtoColor     string
    MsgList      string
    Opponent     string
  }
  var view View
  view.CacheV = common.CACHE_V
  view.CSRF = common.MakeCSRF(w,r)
  u_id := common.GetUser(w,r)
  eto := common.Eto(u_id)
  view.Myphoto = eto[0]
  view.EtoColor = eto[1]
  view.Opponent = "0"
  uri := strings.Split(r.URL.Path, "/")
  fmt.Printf("url %#v\n", uri[2])

  // var msgList MsgList
  k1 := map[string][]map[string]string{}
  // arr := []map[string]string{}
  map1 := map[string]string{}
  map1["test1"] = "mada"
  map1["test2"] = "mada2"
  k1["1"] = append(k1["1"], map1)
  fmt.Printf("k1 %#v\n", k1)
  map1 = map[string]string{}
  map1["test1"] = "mada3"
  map1["test2"] = "mada4"
  k1["1"] = append(k1["1"], map1)
  // k1["1"] = append(k1["1"], map1)
  // k1["1"] = arr
  // k["SenderFlg"] = "message_content"
  // s := map[string]map[string]string{}
  // s["2"] = k
  // usrMap1 := map[string]map[string]map[string]string{}
  // usrMap1["1"] = s

  // fmt.Printf("usrMap1 %#v\n", usrMap1)
  // return
  t := time.Now()
  t1 := t.AddDate(0, -1, 0)
  time.Now().Format("2006-01-02 15:04:05")
  rows, err := db.Query("SELECT "+
    "sender,receiver,sender_img,message_content,message_id,created_at "+
    "FROM t_message                     "+
    "WHERE created_at > $1              "+
    "AND (sender = $2 OR receiver = $2) "+
    "ORDER BY created_at DESC ", 
    t1.Format("2006-01-02 15:04:05"), u_id)
  if err != nil {
      log.Print(err)
  }
  
  usrMap := map[int][]map[string]string{}
  // usrK := map[string][]string{}
  for rows.Next() {
      var sender int
      var receiver int
      var sender_img string
      var message_content string
      var message_id string
      var created_at string
      if err := rows.Scan(&sender,&receiver,&sender_img,&message_content,&message_id,&created_at); err != nil {
          log.Print(err)
      }
      d := map[string]string{}
      d["Txt"] = html.EscapeString(message_content)
      d["OpponentImg"] = sender_img
      d["CreatedAt"] = created_at
      if u_id == sender {
        d["SenderFlg"] = "1"
        usrMap[receiver] = append(usrMap[receiver], d)
      } else {
        d["SenderFlg"] = "0"
        usrMap[sender] = append(usrMap[sender], d)
      }
  }
  if uri[2] != "" {
    view.Opponent = uri[2]
    op, err := strconv.Atoi(uri[2])
    if err != nil {
      fmt.Printf("err %#v\n", err)
      return
    }
    if op < 0 && u_id < 0 {
      fmt.Printf("op < 0 && u_id < 0 %#v\n", op)
      return
    }
    usrMap[op] = append(usrMap[op], map[string]string{})
  }
  
  var msgList []MsgList
  for i, v := range usrMap {
    var msg []Msg
    l := MsgList{}
    for _, v2 := range v{
      m := Msg{}
      m.Txt = v2["Txt"]
      m.SenderFlg = v2["SenderFlg"]
      m.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", v2["CreatedAt"])
      msg = append([]Msg{m}, msg...)
    }
    l.Msg = msg
    if usrMap[i][0]["OpponentImg"] == "" {
      eto := common.Eto(i)
      l.OpponentImg   = eto[0]
      l.OpponentColor = eto[1]
    } else {
      l.OpponentImg   = usrMap[i][0]["OpponentImg"]
    }
    l.OpponentID = i
    msgList = append(msgList, l)
  }
  
  res, _ := json.Marshal(msgList)
  
  view.MsgList = string(res)
  // sort.Slice(usr, func(i, j int) bool { return usr[i].Sum > usr[j].Sum }) // DESC

  tpl := template.Must(template.ParseFiles("html/message.html"))
  tpl.Execute(w, view)

}
