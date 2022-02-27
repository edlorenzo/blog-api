package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/edlorenzo/blog-api/model"
)

// CreateArticle godoc
// @Summary Create an article
// @Description Create an article. Auth is require
// @ID create-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param article body articleCreateRequest true "Article to create"
// @Success 201 {object} singleArticleResponse
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles [post]
func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	var a model.Article
	req := &articleCreateRequest{}
	if err := req.bind(c, &a, h.validator); err != nil {
		status := http.StatusBadRequest
		message := fmt.Sprintf("error trying to parse the body: %v", err.Error())
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, &a))
	}

	usr := userIDFromToken(c)
	a.UserID, a.Creator, a.Modifier = usr, usr, usr
	a.User.ID = a.UserID

	err := h.articleStore.CreateArticle(&a)
	if err != nil {
		status := http.StatusUnprocessableEntity
		message := fmt.Sprintf("error creating the article: %v", err.Error())
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, &a))
	}

	status := http.StatusCreated
	message := "success"
	return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, &a))
}

// UpdateArticle godoc
// @Summary Update an article
// @Description Update an article. Auth is required
// @ID update-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the article to update"
// @Param article body articleUpdateRequest true "Article to update"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{id} [put]
func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	var article = model.Article{}
	id, _ := c.ParamsInt("id")
	article.ID = uint(id)

	a, err := h.articleStore.GetArticleByIDs(&article)
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error retreiving article info: %v", err.Error())
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
	}

	if a == nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("error no record found")
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
	}
	req := &articleUpdateRequest{}
	req.populate(a)
	if err := req.bind(c, a, h.validator); err != nil {
		status := http.StatusUnprocessableEntity
		message := fmt.Sprintf("error trying to parse the body: %v", err.Error())
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
	}

	usr := userIDFromToken(c)
	a.UserID, a.Modifier = usr, usr
	if err = h.articleStore.UpdateArticle(a); err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error updating the article: %v", err.Error())
		return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
	}

	status := http.StatusOK
	message := "success"
	return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article. Auth is required
// @ID delete-article
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the article to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /articles/{id} [delete]
func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
	var article = model.Article{}
	id, _ := c.ParamsInt("id")
	article.ID = uint(id)

	a, err := h.articleStore.GetArticleByIDs(&article)
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error retreiving article info: %v", err.Error())
		return c.Status(status).JSON(newArticleDeleteResponse(status, message))
	}
	if a == nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("error no record found")
		return c.Status(status).JSON(newArticleDeleteResponse(status, message))
	}
	err = h.articleStore.DeleteArticle(a)
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error deleting the article: %v", err.Error())
		return c.Status(status).JSON(newArticleDeleteResponse(status, message))
	}

	status := http.StatusOK
	message := "success"
	return c.Status(status).JSON(newArticleDeleteResponse(status, message))
}

// GetArticle godoc
// @Summary Get an article
// @Description Get an article. Auth not required
// @ID get-article-by-id
// @Tags article
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the article to get"
// @Success 200 {object} singleArticleResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /articles/{id} [get]
func (h *Handler) GetArticle(c *fiber.Ctx) error {
	var article = model.Article{}
	id, _ := c.ParamsInt("id")
	article.ID = uint(id)

	a, err := h.articleStore.GetArticleByIDs(&article)
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error retreiving article info: %v", err.Error())
		return c.Status(status).JSON(map[string]interface{}{"message": message})
	}
	if a == nil {
		status := http.StatusNotFound
		message := fmt.Sprintf("error no record found")
		return c.Status(status).JSON(map[string]interface{}{"message": message})
	}

	status := http.StatusOK
	message := "success"
	return c.Status(status).JSON(newArticleResponse(userIDFromToken(c), status, message, a))
}

// Article godoc
// @Summary Get all articles with Limit and Offset.
// @Description Get most recent article globally. Auth is optional
// @ID get-article
// @Tags article
// @Accept json
// @Produce json
// @Param limit path int true "Limit number of articles returned (default is 20)"
// @Param offset path int true "Offset/skip number of articles (default is 0)"
// @Success 200 {object} articleDataListResponse
// @Failure 500 {object} utils.Error
// @Router /articles/{limit}/{offset} [get]
func (h *Handler) Article(c *fiber.Ctx) error {
	var (
		articles []model.Article
	)

	offset, err := c.ParamsInt("offset")
	if err != nil {
		offset = 0
	}

	limit, err := c.ParamsInt("limit")
	if err != nil {
		limit = 20
	}

	articles, _, err = h.articleStore.ListLimitOffset(offset, limit)
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error retreiving article info: %v", err.Error())
		return c.Status(status).JSON(map[string]interface{}{"message": message})
	}

	status := c.Response().StatusCode()
	message := "success"
	return c.Status(status).JSON(newArticleListResponse(h.userStore, articles, status, message))
}

// ArticleList godoc
// @Summary Get all articles
// @Description Get article list. Auth not required
// @ID get-article-lists
// @Tags article
// @Accept  json
// @Produce  json
// @Success 200 {object} articleDataListResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /articles/list [get]
func (h *Handler) ArticleList(c *fiber.Ctx) error {
	var (
		articles []model.Article
	)

	articles, err := h.articleStore.List()
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("error retreiving article info: %v", err.Error())
		return c.Status(status).JSON(map[string]interface{}{"message": message})
	}

	status := c.Response().StatusCode()
	message := "success"
	return c.Status(status).JSON(newArticleListResponse(h.userStore, articles, status, message))
}
