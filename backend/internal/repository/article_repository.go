package repository

import (
	"intern-article-api/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository (db *gorm.DB) *ArticleRepository {
	db.AutoMigrate(&model.Article{})
	return &ArticleRepository {db :db}
}

func (r *ArticleRepository) SaveAll(articles []model.Article) error {
	return r.db.Save(&articles).Error
}

func (r *ArticleRepository) FindAll() ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *ArticleRepository) Delete (id int) error {
	err := r.db.Delete(&model.Article{}, id).Error
	return err
}

func (r *ArticleRepository) Save(article *model.Article) error {
	return r.db.Save(article).Error
}

func(r *ArticleRepository) Update(id int, title string, body string) error {
	article := model.Article(ID: id)

	result := r.db.Model(&article).Updates(model.Article {
		Title: title,
		Body: body,
	})

	return result.Error
}


