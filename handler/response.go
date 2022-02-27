package handler

import (
	"time"

	_ "github.com/gofiber/fiber/v2"

	"github.com/edlorenzo/blog-api/model"
	"github.com/edlorenzo/blog-api/user"
	"github.com/edlorenzo/blog-api/utils"
)

type userResponse struct {
	User struct {
		ID        uint      `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		RoleType  uint      `json:"role_type"`
		Creator   uint      `json:"creator"`
		Modifier  uint      `json:"modifier"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.RoleType = u.RoleType
	r.User.Creator = u.Creator
	r.User.Modifier = u.Modifier
	r.User.Token = utils.GenerateJWT(u.ID)
	r.User.CreatedAt = u.CreatedAt
	r.User.UpdatedAt = u.UpdatedAt

	return r
}

type profileResponse struct {
	Profile struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		RoleType uint   `json:"role_type"`
	} `json:"profile"`
}

func newProfileResponse(u *model.User) *profileResponse {
	r := new(profileResponse)
	r.Profile.Username = u.Username
	r.Profile.Email = u.Email
	r.Profile.RoleType = u.RoleType
	return r
}

type singleArticleResponse struct {
	Article *articleResponse `json:"article"`
}

type articleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id       uint   `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		Author   string `json:"author"`
		Creator  uint   `json:"creator"`
		Modifier uint   `json:"modifier"`
		User     struct {
			ID        uint
			CreatedAt time.Time
			UpdatedAt time.Time
		} `json:"user"`
	} `json:"data"`
}

type articleDataListResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []*DataList `json:"data"`
}

type articleStatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type articleDeleteResponse struct {
	Article articleStatusResponse `json:"article"`
}

type articleListResponse struct {
	Article articleDataListResponse `json:"article"`
}

type DataList struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Creator  uint   `json:"creator"`
	Modifier uint   `json:"modifier"`
	User     struct {
		ID        uint
		CreatedAt time.Time
		UpdatedAt time.Time
	} `json:"user"`
}

func newArticleResponse(userID uint, status int, message string, a *model.Article) *singleArticleResponse {
	ar := new(articleResponse)
	ar.Status = status
	ar.Message = message
	if a != nil {
		ar.Data.Id = a.ID
		ar.Data.Title = a.Title
		ar.Data.Content = a.Content
		ar.Data.Author = a.Author
		ar.Data.Creator = a.Creator
		ar.Data.Modifier = a.Modifier
		ar.Data.User.ID = a.User.ID
		ar.Data.User.CreatedAt = a.User.CreatedAt
		ar.Data.User.UpdatedAt = a.User.UpdatedAt
	}

	return &singleArticleResponse{ar}
}

func newArticleListResponse(us user.Store, articles []model.Article, status int, message string) *articleListResponse {
	r := new(articleListResponse)
	r.Article.Status = status
	r.Article.Message = message
	r.Article.Data = make([]*DataList, 0)
	for _, a := range articles {
		ar := new(DataList)
		ar.ID = a.ID
		ar.Title = a.Title
		ar.Content = a.Title
		ar.Author = a.Author
		ar.Creator = a.Creator
		ar.Modifier = a.Modifier

		usr, err := us.GetByID(a.UserID)
		if err == nil {
			ar.User.ID = usr.Model.ID
			ar.User.CreatedAt = usr.Model.CreatedAt
			ar.User.UpdatedAt = usr.Model.UpdatedAt
		}
		r.Article.Data = append(r.Article.Data, ar)
	}

	return r
}

func newArticleDeleteResponse(status int, message string) *articleDeleteResponse {
	ar := new(articleDeleteResponse)
	ar.Article.Status = status
	ar.Article.Message = message

	return ar
}
