package postrges

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/gommon/log"
	"blog/api/models"
)

type ArticleRepository interface {
	Save(models.Article) error
	GetByID(int) (models.Article, error)
}

func NewArticleRepository(db *pg.DB) ArticleRepository {
	return &articleRepository{db: db}
}

type articleRepository struct {
	db *pg.DB
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
	err := a.db.Model(&article).Where("id = ?", id).First()
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func New() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "vitmantekoR_2408",
		Database: "simpleblog",
		Addr:     "localhost:5432",
	})
	defer db.Close()
	if db == nil {
		log.Printf("Failed to connect to database!\n")
	}
	log.Printf("Connection to database successful.\n")
	CreateTableArticle(db)

	/*artRep := articleRepository{
		db:db,
	}
	art1, err := artRep.GetByID(1)
	if err!= nil {
		fmt.Println("s")
	}
	article := models.Article{
		ID:3,
		Title:"Third",
		Author:"Mr.Max",
		Content:"eee",
	}
	artRep.Save(article)
	fmt.Println(art1)
	*/
	//GetAllArticles(db)
	return db
}

func CreateTableArticle(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&models.Article{}, opts)
	if err != nil {
		log.Printf("error while creating table Article, Reason: %v\n", err)
		return err
	}
	return nil
} /*
func GetAllArticles(db *pg.DB) error {
	var article []models.Article
	err := db.Select(&article)
	if err != nil {
		log.Printf("error while trying to Select All Articles, Reason: %v\n", err)
		return err
	}
	fmt.Println(article)
	return nil
}
*/
