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
	//.encファイルを確認
	godotenv.Load()
	url := os.Getenv("EXTERNAL_API_URL")
	if url == "" {
	//中身が空ならその場で停止
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

	router.Use(func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/articles", func (w http.ResponseWriter, r *http.Request) {

		articles, _ := svc.GetArticles()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)

	})

	//外部データ取得
	router.Post("/articles/import", func(w http.ResponseWriter, r *http.Request) {

		err := svc.ImportExternalArticle()	
		if err != nil {
			http.Error(w, "保存に失敗しました", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("外部データも取り込みに成功しました"))
	})

	router.Post("/articles", func(w http.ResponseWriter, r *http.Request) {
    		var input struct {
			Title string `json:"title"`
			Body string `json:"body"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return
		}
		svc.CreateArticle(input.Title, input.Body)
		w.WriteHeader(http.Status)
		w.Write([]byte("自分の記事を保存しました"))
	})

	router.Delete("/articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		fmt.Sscanf(idStr, "%d", &id)
	
	err := svc.DeleteArticle(id)
    
	if err != nil {
        http.Error(w, "削除に失敗しました", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
})

	router.Patch("/articles/{id}/pin", func(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    var id int
    fmt.Sscanf(idStr, "%d", &id)

    err := svc.TogglePin(id)
    if err != nil {
        http.Error(w, "ピン留めの更新に失敗しました", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("ピン留めを更新しました"))
})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080",router)

}

	router.Put("/artiles/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		fmt.Ssacnf(idStr, "%d", &id)
		var input struct {
            Title string `json:"title"`
            Body  string `json:"body"`
        }
        json.NewDecoder(r.Body).Decode(&input)

        // 3. シェフに更新を依頼
        err := svc.UpdateArticle(id, input.Title, input.Body)
        if err != nil {
            http.Error(w, "更新できませんでした", 500)
            return
        }

        w.Write([]byte("更新しました"))
    })