package hello

import (
    "fmt"
    "net/http"
    "encoding/json"
)


// ヘッダー情報構造体 for json
type List struct {
	Key int				`json:"key"`			// 後でデータベースのキーに変更 *datastore.Key
	Title string		`json:"title"`
	Position string		`json:"position"`		// 緯度経度
	Adr string			`json:"adr"`
	Date string			`json:"date"`			// 後でデータベースの日付に変更 time.Time 
}

// ヘッダー一覧用構造体 for json
type Listslice struct {
	Lists []List
}



func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/list", list)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world! 4")
}

// とりあえず、固定で結果を返せるようにしておく。正式実装次第、差し替える
func list(w http.ResponseWriter, r *http.Request) {

	var l Listslice
	l.Lists = append(l.Lists, List{Title: "世界一うまいラーメン", Position: "35.0394195,135.7915279", Adr: "大阪" } )
	l.Lists = append(l.Lists, List{Title: "Test2", Adr: "兵庫" } )
	j, err := json.Marshal(l)
	if err != nil {
		fmt.Fprint(w, "json err:", err)
	}

    fmt.Fprint(w, string(j))
}
