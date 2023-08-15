package routes

import (
	"go_blog/api"
	"net/http"
)

// 跨域处理中间件
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置请求头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 调用下一个处理程序
		next.ServeHTTP(w, r)
	})
}

func InitRouter() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user/add", api.AddUser)
	mux.HandleFunc("/user/getAll", api.GetUsers)

	mux.HandleFunc("/category/add", api.AddCategory)
	mux.HandleFunc("/category/getAll", api.GetCategories)
	mux.HandleFunc("/category/update", api.UpdateCategory)
	mux.HandleFunc("/category/delete", api.DeleteCategory)

	mux.HandleFunc("/article/add", api.AddArticle)
	mux.HandleFunc("/article/getAll", api.GetArticleList)
	mux.HandleFunc("/article/delete", api.DeleteArticle)
	mux.HandleFunc("/article/update", api.UpdateArticle)
	mux.HandleFunc("/article/getById", api.GetArticleById)
	mux.HandleFunc("/article/search", api.SearchArticle)

	handler := corsMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.ListenAndServe()
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("hello world"))
// 	w.Header().Set("content-type", "application/json")

// 	data, _ := json.Marshal(IndexData)
// 	w.Write(data)
// }
