package usecases

import (
	"blog/src/models"
	"blog/src/repository/postgres"
	"blog/src/usecases/rbac"
	"errors"
	"github.com/labstack/echo"
	"log"
	"time"
)

type ArticleService interface {
	Create(echo.Context, CreateReqArt) error
	View(int) (models.Article, error)
	Delete(echo.Context, int) error
	List() ([]models.Article, error)
	Update(echo.Context, int, UpdateReqArt) error
	GetByUsername(string) ([]models.Article, error)
}

type articleService struct {
	artRep postgres.ArticleRepository
	rbac   rbac.RBAC
}

func NewArtService(artRep postgres.ArticleRepository, rbac rbac.RBAC) ArticleService {
	return &articleService{artRep: artRep, rbac: rbac}
}

type UpdateReqArt struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type CreateReqArt struct {
	Title    string `json:"title"`
	Username string `json:"username"`
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
}

func (s articleService) GetByUsername(username string) ([]models.Article, error) {
	article, err := s.artRep.GetByUsername(username)
	if err != nil {
		log.Printf("error GU, Reason: %v\n", err)
		return article, err
	}

	return article, nil
}

func (s articleService) Create(c echo.Context, req CreateReqArt) error {
	if s.rbac.IsBlocked(c) {
		return errors.New("Your user is blocked for posting articles\n")
	}

	article := models.Article{
		Base: models.Base{
			CreatedAt: time.Now(),
		},
		Title:    req.Title,
		Username: req.Username,
		Content:  req.Content,
		UserID:   req.UserID,
	}
	if err := s.artRep.Create(article); err != nil {
		log.Printf("error SA, Reason: %v\n", err)
		return err
	}

	return nil
}

func (s articleService) Update(c echo.Context, id int, req UpdateReqArt) error {
	if !s.rbac.EnforceUser(c, id) {
		return errors.New("It`s not your article or you`re not an admin\n")
	}

	updArticle := models.Article{
		Base: models.Base{
			UpdatedAt: time.Now(),
		},
		Title:   req.Title,
		Content: req.Content,
	}
	if err := s.artRep.Update(id, updArticle); err != nil {
		log.Printf("error UA, Reason: %v\n", err)
		return err
	}

	return nil
}

func (s articleService) List() ([]models.Article, error) {
	articles, err := s.artRep.List()
	if err != nil {
		log.Printf("error GA, Reason: %v\n", err)
		return nil, err
	}

	return articles, nil
}

func (s articleService) Delete(c echo.Context, id int) error {
	if !s.rbac.EnforceUser(c, id) {
		return errors.New("It`s not your user or you`re not an admin\n")
	}

	if err := s.artRep.Delete(id); err != nil {
		log.Printf("error DA, Reason: %v\n", err)
		return err
	}

	return nil
}

func (s articleService) View(id int) (models.Article, error) {
	article, err := s.artRep.View(id)
	if err != nil {
		log.Printf("error GIA, Reason: %v\n", err)
		return article, err
	}

	return article, nil
}
