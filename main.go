package main

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/edlorenzo/blog-api/db"
	_ "github.com/edlorenzo/blog-api/docs"
	"github.com/edlorenzo/blog-api/handler"
	"github.com/edlorenzo/blog-api/router"
	"github.com/edlorenzo/blog-api/store"
)

// @description Blog API
// @title Blog API

// @BasePath /api

// @schemes http https
// @produce article/json
// @consumes article/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()
	r.Get("/swagger/*", swagger.HandlerDefault)
	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)

	h := handler.NewHandler(us, as)
	h.Register(r)
	err := r.Listen(":5007")
	if err != nil {
		fmt.Printf("%v", err)
	}
}
