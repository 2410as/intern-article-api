package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
    "intern-article-api/internal/repository" 
    "intern-article-api/internal/service"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func main () {
	godotenv.Load()
	url := os.Getenv("EXTERNAL_API_URL")
	if url == "" {
    panic("EXTERNAL_API_URL が設定されていません！")
}
	
	db, err := gorm.Open(sqlite.Open("articles.db"), &gorm.Config{})
	if err != nil {
		panic("DBに失敗しました")
	}
	
	//実体化
	repo := repository.NewArticleRepository(db)
	svc := service.NewArticleService(repo, url)

	router := chi.NewRouter()

	router.Get("/articles", func (w http.ResponseWriter, r *http.Request) {
		articles, err := svc.GetArticles()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)

	})

	//外部データ取得
	router.Post("/articles", func(w http.ResponseWriter, r *http.Request) {
		err := svc.ImportExternalArticle()
		if err !=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("外部データの取り込みに成功しました"))
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080",router)

}