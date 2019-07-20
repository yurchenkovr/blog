package postgres

import (
	"blog/src/models"
	"github.com/go-pg/pg"
	"log"
)

type ArticleRepository interface {
	SaveArticle(models.Article) error
	GetByIDArticle(int) (models.Article, error)
	DeleteArticle(int) error
	GetAllArticles() ([]models.Article, error)
	UpdateArticle(int, models.Article) error
}

func NewArticleRepository(db *pg.DB) ArticleRepository {
	return &articleRepository{db: db}
}

type articleRepository struct {
	db *pg.DB
}

func (a *articleRepository) UpdateArticle(id int, art models.Article) error {
	var updatedArt models.Article
	if _, err := a.db.Model(&updatedArt).Set("title = ?, author = ?, content = ?", art.Title, art.Author, art.Content).
		Where("id = ?", id).Update(); err != nil {
		log.Printf("Error while updating, Reason: %v\n", err)
		return err
	}
	return nil
}
func (a *articleRepository) SaveArticle(article models.Article) error {
	if err := a.db.Insert(&article); err != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", err)
		return err
	}
	return nil
}

func (a *articleRepository) GetByIDArticle(id int) (models.Article, error) {
	var article models.Article
	if err := a.db.Model(&article).Where("id = ?", id).First(); err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (a *articleRepository) DeleteArticle(id int) error {
	var article models.Article
	if _, err := a.db.Model(&article).Where("id = ?", id).Delete(); err != nil {
		log.Printf("Error while deleting, Reason: %v\n", err)
		return err
	}
	return nil
}
func (a *articleRepository) GetAllArticles() ([]models.Article, error) {
	var article []models.Article
	if err := a.db.Model(&article).Select(); err != nil {
		log.Printf("error while trying to Select All Articles, Reason: %v\n", err)
		return nil, err
	}
	return article, nil
}
