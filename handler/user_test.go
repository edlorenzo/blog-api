package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/edlorenzo/blog-api/utils"
)

func TestSignUpCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"username":"unittestuser1", "email":"unittestuser1@gmail.com","password":"abc123"}}`
	)
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set("Content-type", "application/json")
	h.Register(e)
	resp, _ := e.Test(req, -1)
	if assert.Equal(t, http.StatusCreated, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginCaseSuccess(t *testing.T) {
	var (
		reqJSON = `{"user":{"email":"unittestuser1@gmail.com","password":"abc123"}}`
	)
	req := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set("Content-type", "application/json")
	h.Register(e)
	resp, _ := e.Test(req, -1)
	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
	}
}

func TestLoginCaseFailed(t *testing.T) {
	var (
		reqJSON = `{"user":{"email":"wrongunittestuser1@gmail.com","password":"wrongabc123"}}`
	)
	req := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set("Content-type", "application/json")
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

func TestCurrentUserCaseSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}
func TestCurrentUserCaseInvalid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(100)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestUpdateUserEmail(t *testing.T) {
	var (
		user1UpdateReq = `{"user":{"email":"unittestuser1@gmail.com","password":"updatedpassword","role_type":7}}`
	)
	req := httptest.NewRequest(http.MethodPut, "/api/user", strings.NewReader(user1UpdateReq))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestUpdateUserMultipleField(t *testing.T) {
	var (
		user1UpdateReq = `{"user":{"email":"unittestuser1@gmail.com","password":"updatedpassword","role_type":7}}`
	)
	req := httptest.NewRequest(http.MethodPut, "/api/user", strings.NewReader(user1UpdateReq))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestGetProfileCaseSuccess(t *testing.T) {
	username := "unittestuser1"
	url := "/api/profiles/" + username
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)

	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, resp.StatusCode) {
		body, _ := ioutil.ReadAll(resp.Body)
		m := responseMap(body, "profile")
		assert.Equal(t, "unittestuser1", m["username"])
		assert.Equal(t, "unittestuser1@gmail.com", m["email"])
	}
}

func TestGetProfileCaseNotFound(t *testing.T) {
	username := "wrongunittestuser1"
	url := "/api/profiles/" + username
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)

	resp, err := e.Test(req, -1)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
