package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/edlorenzo/blog-api/model"
)

type userUpdateRequest struct {
	User struct {
		Username  string    `json:"-"`
		Email     string    `json:"email" validate:"email"`
		Password  string    `json:"password"`
		RoleType  uint      `json:"role_type"`
		Modifier  uint      `json:"-"`
		UpdatedAt time.Time `json:"-"`
	} `json:"user"`
}

func newUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}
func (r *userUpdateRequest) populate(u *model.User) {
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Password = u.Password
	r.User.RoleType = u.RoleType
	r.User.Modifier = u.ID
	r.User.UpdatedAt = u.UpdatedAt
}

func (r *userUpdateRequest) bind(c *fiber.Ctx, u *model.User, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email

	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h

	}
	return nil
}

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		RoleType uint   `json:"role_type"`
		Creator  uint   `json:"-"`
		Modifier uint   `json:"-"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c *fiber.Ctx, u *model.User, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c *fiber.Ctx, v *Validator) error {

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

type articleCreateRequest struct {
	Article struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Author   string `json:"author"`
		Creator  uint   `json:"-"`
		Modifier uint   `json:"-"`
	} `json:"article"`
}

func (r *articleCreateRequest) bind(c *fiber.Ctx, a *model.Article, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	a.Title = r.Article.Title
	a.Content = r.Article.Content
	a.Author = r.Article.Author

	return nil
}

type articleUpdateRequest struct {
	Article struct {
		ID       uint   `json:"-"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		Author   string `json:"author"`
		Modifier uint   `json:"-"`
		UserID   uint   `json:"-"`
	} `json:"article"`
}

func (r *articleUpdateRequest) populate(a *model.Article) {
	r.Article.Title = a.Title
	r.Article.Content = a.Content
	r.Article.Author = a.Author
	r.Article.Modifier = a.Modifier
	r.Article.UserID = a.UserID
}

func (r *articleUpdateRequest) bind(c *fiber.Ctx, a *model.Article, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	a.Title = r.Article.Title
	a.Content = r.Article.Content
	a.Author = r.Article.Author
	a.Modifier = r.Article.Modifier
	a.UserID = r.Article.UserID

	return nil
}
