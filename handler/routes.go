package handler

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"

	"github.com/edlorenzo/blog-api/utils"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	jwtMiddleware := jwt.New(
		jwt.Config{
			SigningKey: utils.JWTSecret,
			AuthScheme: "Token",
		})

	guestUsers := v1.Group("/users")
	guestUsers.Post("", h.SignUp)
	guestUsers.Post("/login", h.Login)
	user := v1.Group("/user", jwtMiddleware)
	user.Get("", h.CurrentUser)
	user.Put("", h.UpdateUser)

	profiles := v1.Group("/profiles", jwtMiddleware)
	profiles.Get("/:username", h.GetProfile)

	articlesJWTMiddleware := jwt.New(
		jwt.Config{
			SigningKey: utils.JWTSecret,
			AuthScheme: "Token",
			Filter: func(c *fiber.Ctx) bool {
				if c.Method() == "GET" && c.Path() != "/api/articles/feed" {
					return true
				}
				return false
			},
		})

	articles := v1.Group("/articles", articlesJWTMiddleware)
	articles.Get("/list", h.ArticleList)
	articles.Get("/:limit/:offset", h.Article)
	articles.Post("", h.CreateArticle)
	articles.Put("/:id", h.UpdateArticle)
	articles.Delete("/:id", h.DeleteArticle)
	articles.Get("/:id", h.GetArticle)
}
