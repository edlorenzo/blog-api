package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/edlorenzo/blog-api/model"
)

func New() *gorm.DB {
	dsn := "./database/blog.db"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	/*
	 *db, err := gorm.Open(postgres.New(postgres.Config{
	 *  DSN: dsn,
	 *  //PreferSimpleProtocol: true, // disables implicit prepared statement usage
	 *}), &gorm.Config{})
	 */

	//db, err := gorm.Open("postgresql", "postgresql://blog@/blog?host=/tmp")
	//db, err := gorm.Open("sqlite3", "./database/blog.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func TestDB() *gorm.DB {
	dsn := "./../database/blog_test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../database/blog_test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Article{},
	)
	if err != nil {
		return
	}
}
