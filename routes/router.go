package routes

import (
	"go_blog/api"
	"net/http"
)

func InitRouter() {

	http.HandleFunc("/user/add", api.AddUser)
	http.HandleFunc("/user/getAll", api.GetUsers)

	http.HandleFunc("/category/add", api.AddCategory)
	http.HandleFunc("/category/getAll", api.GetCategories)
	http.HandleFunc("/category/update", api.UpdateCategory)
	http.HandleFunc("/category/delete", api.DeleteCategory)

	http.HandleFunc("/article/add", api.AddArticle)
	http.HandleFunc("/article/getAll", api.GetArticleList)
	http.HandleFunc("/article/delete", api.DeleteArticle)
	http.HandleFunc("/article/update", api.UpdateArticle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("hello world"))
// 	w.Header().Set("content-type", "application/json")

// 	data, _ := json.Marshal(IndexData)
// 	w.Write(data)
// }
