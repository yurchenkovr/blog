package postgres

import (
	"blog/src/models"
	"github.com/go-pg/pg"
	"log"
)

type ArticleRepository interface {
	Create(models.Article) (models.Article, error)
	View(int) (models.Article, error)
	Delete(int) error
	List() ([]models.Article, error)
	Update(int, models.Article) error
	GetByUsername(string) ([]models.Article, error)
}

func NewArticleRepository(db *pg.DB) ArticleRepository {
	return &articleRepository{db: db}
}

type articleRepository struct {
	db *pg.DB
}

func (a *articleRepository) GetByUsername(username string) ([]models.Article, error) {
	var article []models.Article

	if err := a.db.Model(&article).Where("username = ?", username).Select(); err != nil {
		return []models.Article{}, err
	}
	return article, nil
}

func (a *articleRepository) Update(id int, art models.Article) error {
	var updatedArt models.Article

	if _, err := a.db.Model(&updatedArt).Set("title = ?,content = ?, updated_at = ?", art.Title, art.Content, art.UpdatedAt).
		Where("id = ?", id).Update(); err != nil {
		log.Printf("Error while updating, Reason: %v\n", err)
		return err
	}
	return nil
}

func (a *articleRepository) Create(article models.Article) (models.Article, error) {
	if err := a.db.Insert(&article); err != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", err)
		return models.Article{}, err
	}
	return article, nil
}

func (a *articleRepository) View(id int) (models.Article, error) {
	var article models.Article

	if err := a.db.Model(&article).Where("id = ?", id).First(); err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (a *articleRepository) Delete(id int) error {
	var article models.Article

	if _, err := a.db.Model(&article).Where("id = ?", id).Delete(); err != nil {
		log.Printf("Error while deleting, Reason: %v\n", err)
		return err
	}
	return nil
}

func (a *articleRepository) List() ([]models.Article, error) {
	var article []models.Article

	if err := a.db.Model(&article).Select(); err != nil {
		log.Printf("error while trying to Select All Articles, Reason: %v\n", err)
		return nil, err
	}
	return article, nil
}
