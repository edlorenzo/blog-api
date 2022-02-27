package handler

import (
	"github.com/edlorenzo/blog-api/article"
	"github.com/edlorenzo/blog-api/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
	validator    *Validator
}

func NewHandler(us user.Store, as article.Store) *Handler {
	v := NewValidator()
	return &Handler{
		userStore:    us,
		articleStore: as,
		validator:    v,
	}
}
