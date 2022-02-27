package store

import (
	"errors"

	"gorm.io/gorm"

	"github.com/edlorenzo/blog-api/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByEmail(e string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Updates(u).Error
}
