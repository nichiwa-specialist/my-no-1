package gaehelloworld

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
	"time"
//	"log"
	"appengine"
	"appengine/datastore"
	"io/ioutil"
)


// ヘッダー情報構造体 for json
type List struct {
	Key *datastore.Key	`json:"key"`			// 後でデータベースのキーに変更 *datastore.Key
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Adr string			`json:"adr"`			// 住所
	Date time.Time		`json:"date"`			// 後でデータベースの日付に変更 time.Time 
}

// 詳細情報構造体 for json
type Detail struct {
//	Key int				`json:"key"`			// 後でデータベースのキーに変更 *datastore.Key
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Adr string			`json:"adr"`			// 住所
	Date time.Time		`json:"date"`			// 後でデータベースの日付に変更 time.Time
	Detail string		`json:"detail"`			// 本文
}

// Post情報構造体 for json
type PostData struct {
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Detail string		`json:"detail"`			// 本文
}

type Message struct {
	Message string		`json:"Error"`
}

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/article/", article)
}

func getRemarkID(path string) (int, bool) {
    split := strings.Split(path, "/")
    length := len(split)
	id, err := strconv.Atoi(split[length-1])
    return id, (err == nil)
}

func article(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    
    switch r.Method {
    case "POST":		// POST の仮に GET でテスト

		var p_data PostData

		fmt.Fprint(w, r.Body)
		
		body, err := ioutil.ReadAll(r.Body)
   		if err != nil {
   			fmt.Fprint(w, err)
   		}
		
		err = json.Unmarshal(body, &p_data)
   		if err != nil {
   			fmt.Fprint(w, err)
   		}




   		d := Detail{
   			Title: p_data.Title,
   			Lat: p_data.Lat,
   			Lng: p_data.Lng,
   			Adr: "大阪",
   			Date: time.Now(),
   			Detail: p_data.Detail,
   		}
   		
   		key := datastore.NewIncompleteKey(c, "Detail", nil)
   		key, err = datastore.Put(c, key, &d)
   		if err != nil {
   			fmt.Fprint(w, err)
   		}

	case "GET":
		id, foundID := getRemarkID(r.URL.Path)
		
//		var query *datastore.Query
		
		if  foundID {
			m := List{Title: "世界一うまいラーメン" + string(int(id)), Lat: 35.0394195, Lng: 135.7915279, Adr: "大阪" }
			out, _ := json.Marshal(m)
			fmt.Fprint(w, string(out))

//			m := Message{string(id)}
//			out, _ := json.Marshal(m)
//			fmt.Fprint(w, string(out))
		}else{
//			m := List{Title: "世界一うまいラーメン", Lat: 35.0394195, Lng: 135.7915279, Adr: "大阪" }
//			out, _ := json.Marshal(m)
//			fmt.Fprint(w, string(out))
			
//			query = datastore.NewQuery(REMARK_KIND).Order("-Time")
			q := datastore.NewQuery("Detail").Order("-Date")
			list := make([]List, 0, 200)
			if _, err := q.GetAll(c, &list); err != nil {
				
			}
			out, _ := json.Marshal(list)
			fmt.Fprint(w, string(out))
		}




	default:
		m := Message{"Resource not found"}
		out, _ := json.Marshal(m)
		fmt.Fprint(w, string(out))
    }
}

func root(w http.ResponseWriter, r *http.Request) {
	m := Message{"Resource not found"}
	out, _ := json.Marshal(m)
	fmt.Fprint(w, string(out))
}
