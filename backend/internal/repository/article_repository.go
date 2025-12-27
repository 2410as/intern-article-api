package repository

import (
	"github.com/2410as/api-1/internal/model](http://github.com/2410as/api-1/internal/model)"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ArtilceRepository struct {
	db *gorm.DB
}

func NewArticleRepository (dbFile string) (*ArticleRepository, error) {
	db, err := gorm.Open(sqlite.Open (dbfile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Article{})
	return &ArticleRepository {db :db}, nil
}

func (r *NewArticleRepository) SaveAll() ([]model.Article) error {
	return.r.db.Save(&article).Error
}

func (r *ArticleRepository) FindAll() ([]model.Article) error {
	var article []model.Artile
	err :=r.db.Find(&articles).Error
	return article 
}

func (r *ArticleRepository) Delete (id int) error {
	errn:= r.db.delete(&model.Article{}, id).Error
	return err
}

