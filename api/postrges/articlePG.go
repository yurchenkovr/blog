package postrges

import (
	"blog/api/models"
	"github.com/go-pg/pg"
	"log"
)

type ArticleRepository interface {
	Save(models.Article) error
	GetByID(int) (models.Article, error)
	Delete(int) error
	GetAllArticles() ([]models.Article, error)
	Update(int, UpdateArt) error
}

func NewArticleRepository(db *pg.DB) ArticleRepository {
	return &articleRepository{db: db}
}

type articleRepository struct {
	db *pg.DB
}
type UpdateArt struct {
	Title   string
	Author  string
	Content string
}

func (a *articleRepository) Update(id int, art UpdateArt) error {
	var updatedArt models.Article
	if _, err := a.db.Model(&updatedArt).Set("title = ?, author = ?, content = ?", art.Title, art.Author, art.Content).
		Where("id = ?", id).Update(); err != nil {
		log.Printf("Error while updating, Reason: %v\n", err)
		return err
	}
	return nil
}
func (a *articleRepository) Save(article models.Article) error {
	if err := a.db.Insert(&article); err != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", err)
		return err
	}
	return nil
}

func (a *articleRepository) GetByID(id int) (models.Article, error) {
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
func (a *articleRepository) GetAllArticles() ([]models.Article, error) {
	var article []models.Article
	if err := a.db.Model(&article).Select(); err != nil {
		log.Printf("error while trying to Select All Articles, Reason: %v\n", err)
		return nil, err
	}
	return article, nil
}
