package usecases

import (
	"blog/src/models"
	"blog/src/repository/postgres"
	"log"
)

type ArticleService interface {
	SaveArticle(CreateReqArt) error
	GetByIDArticle(int) (models.Article, error)
	DeleteArticle(int) error
	GetAllArticles() ([]models.Article, error)
	UpdateArticle(int, UpdateReqArt) error
}

type articleService struct {
	artRep postgres.ArticleRepository
}

func NewArtService(artRep postgres.ArticleRepository) ArticleService {
	return &articleService{artRep: artRep}
}

type UpdateReqArt struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
type CreateReqArt struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (s articleService) SaveArticle(req CreateReqArt) error {
	article := models.Article{
		ID:      req.ID,
		Title:   req.Title,
		Author:  req.Title,
		Content: req.Content,
	}
	if err := s.artRep.SaveArticle(article); err != nil {
		log.Printf("error SA, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s articleService) UpdateArticle(id int, req UpdateReqArt) error {
	updArticle := models.Article{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
	}
	if err := s.artRep.UpdateArticle(id, updArticle); err != nil {
		log.Printf("error UA, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s articleService) GetAllArticles() ([]models.Article, error) {
	articles, err := s.artRep.GetAllArticles()
	if err != nil {
		log.Printf("error GA, Reason: %v\n", err)
		return nil, err
	}
	return articles, nil
}
func (s articleService) DeleteArticle(id int) error {
	if err := s.artRep.DeleteArticle(id); err != nil {
		log.Printf("error DA, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s articleService) GetByIDArticle(id int) (models.Article, error) {
	article, err := s.artRep.GetByIDArticle(id)
	if err != nil {
		log.Printf("error GIA, Reason: %v\n", err)
		return article, err
	}
	return article, nil
}
