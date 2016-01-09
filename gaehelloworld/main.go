package gaehelloworld

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strconv"
	"fmt"
)
//	"appengine"


// ヘッダー情報構造体 for json
type List struct {
	Key int				`json:"key"`			// 後でデータベースのキーに変更 *datastore.Key
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Adr string			`json:"adr"`			// 住所
	Date string			`json:"date"`			// 後でデータベースの日付に変更 time.Time 
}

// 詳細情報構造体 for json
type Detail struct {
	Key int				`json:"key"`			// 後でデータベースのキーに変更 *datastore.Key
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Adr string			`json:"adr"`			// 住所
	Date string			`json:"date"`			// 後でデータベースの日付に変更 time.Time
	Detail string		`json:"detail"`			// 本文
}

// Post情報構造体 for json
type PostData struct {
	Title string		`json:"title"`			// タイトル
	Lat float32			`json:"lat"`			// 緯度
	Lng float32			`json:"lng"`			// 経度
	Detail string		`json:"detail"`			// 本文
}



func init() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	
	api.Use(&rest.CorsMiddleware{
        OriginValidator: func(origin string, request *rest.Request) bool {
            return true
        },
    })
	
	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!2"})
		}),
		rest.Get("/article", func(w rest.ResponseWriter, req *rest.Request) {
		
			var l []List

			l = append(l, List{Title: "世界一うまいラーメン", Lat: 35.0394195, Lng: 135.7915279, Adr: "大阪" } )
			l = append(l, List{Title: "Test2", Adr: "兵庫" } )
		
			w.WriteJson(l)
		}),
		rest.Get("/article/:key", func(w rest.ResponseWriter, req *rest.Request) {

			var l Detail
//			var intKey int

//			var strKey string
//			strKey = req.PathParam("key")
			intKey, err := strconv.ParseInt( req.PathParam("key"), 10 , 0 )

			if err != nil{
			  fmt.Println(err)
			}
			
			l.Key = int(intKey)
			l.Title = "世界一うまいラーメン2"
			l.Lat = 35.0394195
			l.Lng = 135.7915279
			l.Detail = "詳細"

			w.WriteJson(l)
		}),

//		rest.Get("/article_post", func(w rest.ResponseWriter, req *rest.Request) {
//			
//			post := PostData{}
//			
//			err := req.DecodeJsonPayload(&post)
//			if err != nil {
//				rest.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			
//			c := appengine.NewContext(req)
//			err := db.Save(post.Title, post.Lat, post.Lng, post.Detail, c)
//			
//			
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//		}),


		)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/", api.MakeHandler())
}
