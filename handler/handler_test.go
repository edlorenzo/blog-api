package handler

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"

	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/edlorenzo/blog-api/article"
	"github.com/edlorenzo/blog-api/db"
	"github.com/edlorenzo/blog-api/model"
	"github.com/edlorenzo/blog-api/router"
	"github.com/edlorenzo/blog-api/store"
	"github.com/edlorenzo/blog-api/user"
)

var (
	d  *gorm.DB
	us user.Store
	as article.Store
	h  *Handler
	e  *fiber.App
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func authHeader(token string) string {
	return "Token " + token
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)
	us = store.NewUserStore(d)
	as = store.NewArticleStore(d)

	h = NewHandler(us, as)
	e = router.New()
}
func tearDown() {
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}
func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	u1 := model.User{
		Username: "unittesthandler1",
		Email:    "unittesthandler1@blog.com",
	}
	u1.Password, _ = u1.HashPassword("abc123")
	if err := us.Create(&u1); err != nil {
		return err
	}

	err := us.Create(&u1)
	if err != nil {
		return err
	}

	return nil
}
