package service

 import (
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
