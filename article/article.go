package article

import "github.com/edlorenzo/blog-api/model"

type Store interface {
	CreateArticle(article *model.Article) error
	GetArticleByIDs(ids *model.Article) (*model.Article, error)
	UpdateArticle(*model.Article) error
	DeleteArticle(article *model.Article) error
	ListLimitOffset(offset, limit int) ([]model.Article, int64, error)
	List() ([]model.Article, error)
}
