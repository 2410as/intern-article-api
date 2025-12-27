package service

 import (
	"encoding/json"
	"net/http"
	"github.com/2410as/api-1/internal/model"
 )

type ArticleService struct{
	repo *repository.ArticleRepository
}



func (s *ArticleService) CreateArticle(title string, body string) error {
	article := model.Article{
		Title: title,
		Body: body,
	}
	return s.repo.Save(&article)
}

func (s *ArticleService) ImportExternalArticle(url string) error {
	articles,err := s.FetchExternalArticles(url)
	if err != nil {
		return err
	}

	return s.repo.SaveAll(articles)
}



func (s *ArticleService) GetArticles() ([]model.Article error) {
	articles,err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return articles,nil
}