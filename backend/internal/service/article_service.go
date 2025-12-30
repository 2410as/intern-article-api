package service

 import (
	"fmt"
	"encoding/json"
	"net/http"
	"intern-article-api/internal/model"
	"intern-article-api/internal/repository"
 )

type ArticleService struct{
	repo *repository.ArticleRepository
	externalURL string
}



func (s *ArticleService) CreateArticle(title string, body string) error {
	article := model.Article{
		Title: title,
		Body: body,
	}

	return s.repo.Save(&article)
}

func (s *ArticleService) ImportExternalArticle() error {
	resp, err := http.Get(s.externalURL)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var articles []model.Article

	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return err
	}

	return s.repo.SaveAll(articles)
}



func (s *ArticleService) GetArticles() ([]model.Article, error) {
	articles,err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return articles,nil
}

func NewArticleService(repo *repository.ArticleRepository, url string) *ArticleService {
	return &ArticleService {
	repo:		 repo,
	externalURL: url,
	}
}

func (s *ArticleService) DeleteArticle(id int) error {
    return s.repo.Delete(id)
}

func (s *ArticleService) TogglePin(id int) error {
    // 1. 全記事を持ってきて対象を探す（本当はFindByIDがあると便利ですが、今は既存のrepoを活用します）
    articles, _ := s.repo.FindAll()
    for _, art := range articles {
        if art.ID == id {
            // 2. 状態を反転させる
            art.IsPinned = !art.IsPinned
            // 3. 上書き保存
            return s.repo.Save(&art)
        }
    }
    return fmt.Errorf("記事が見つかりません")
}