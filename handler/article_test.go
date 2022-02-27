package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/edlorenzo/blog-api/utils"
)

func TestCreateArticleSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"article":{"author":"unittest author","content":"unittest content","title":"unittest title"}}`
	)
	req := httptest.NewRequest(http.MethodPost, "/api/articles", strings.NewReader(reqJSON))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, _ := e.Test(req, -1)
	if assert.Equal(t, http.StatusCreated, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)

		m := responseMap(body, "article")
		data := m["data"]
		md, _ := data.(map[string]interface{})

		assert.Equal(t, "unittest author", md["author"])
		assert.Equal(t, "unittest content", md["content"])
		assert.Equal(t, "unittest title", md["title"])
	}
}

func TestGetArticleList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/articles/list", nil)
	req.Header.Set("Content-type", "application/json")
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "article")
		assert.Equal(t, "success", m["message"])
	}
}

func TestGetArticleListLimitOffset(t *testing.T) {
	limit := 5
	offset := 2
	url := fmt.Sprintf("/api/articles/%d/%d", limit, offset)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-type", "application/json")
	h.Register(e)

	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "article")
		assert.Equal(t, "success", m["message"])
	}
}

func TestPutArticle(t *testing.T) {
	var (
		reqUpdate = `{"article":{"author":"unittest updated author","content":"unittest updated content","title":"unittest updated title"}}`
	)
	articleID := 1
	url := fmt.Sprintf("/api/articles/%d", articleID)
	req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(reqUpdate))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)

		m := responseMap(body, "article")
		data := m["data"]
		md, _ := data.(map[string]interface{})

		assert.Equal(t, "unittest updated author", md["author"])
		assert.Equal(t, "unittest updated content", md["content"])
		assert.Equal(t, "unittest updated title", md["title"])
	}
}

func TestDeleteArticle(t *testing.T) {
	articleID := 1
	url := fmt.Sprintf("/api/articles/%d", articleID)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "article")
		assert.Equal(t, "success", m["message"])
	}
}
