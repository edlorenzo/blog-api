package store

import (
	"errors"

	"gorm.io/gorm"

	"github.com/edlorenzo/blog-api/model"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (as *ArticleStore) CreateArticle(a *model.Article) error {
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(a.ID).Preload("User").First(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ArticleStore) UpdateArticle(a *model.Article) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Updates(a).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(a.ID).Preload("User").First(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *ArticleStore) DeleteArticle(a *model.Article) error {
	return as.db.Delete(a).Error
}

func (as *ArticleStore) GetArticleByIDs(application *model.Article) (*model.Article, error) {
	var m model.Article
	err := as.db.Where(application.ID).Preload("User").First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &m, err
}

func (as *ArticleStore) ListLimitOffset(offset, limit int) ([]model.Article, int64, error) {
	var (
		articles []model.Article
		count    int64
	)
	as.db.Model(&articles).Count(&count)
	as.db.Preload("User").
		Offset(offset).
		Limit(limit).
		Order("created_at desc").Find(&articles)
	return articles, count, nil
}

func (as *ArticleStore) List() ([]model.Article, error) {
	var (
		articles []model.Article
	)
	as.db.Model(&articles)
	as.db.Preload("User").
		Order("created_at desc").Find(&articles)
	return articles, nil
}
