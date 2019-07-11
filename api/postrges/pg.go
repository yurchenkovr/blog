package postrges

import (
	"blog/api/models"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/gommon/log"
)

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

	return db
}
func GetAllArticles(artRep articleRepository) {
	articles, err := artRep.GetAllArticles()
	if err != nil {
		log.Printf("Error while tryinf to GetAllArticles. Reason: %v\n", err)
	}
	fmt.Println(articles)
}
func GetByIDArticle(artRep articleRepository) {
	art1, err := artRep.GetByID(1)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(art1)
}
func UpdateArticle(artRep articleRepository) {
	updArt := UpdateArt{
		Title:   "Luck",
		Author:  "James",
		Content: "Something",
	}

	artRep.Update(1, updArt)
}
func DeleteArticle(artRep articleRepository) {
	artRep.Delete(3)
}
func SaveArticle(artRep articleRepository) {
	article := models.Article{
		ID:      3,
		Title:   "Third",
		Author:  "Mr.Max",
		Content: "eee",
	}
	artRep.Save(article)
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
}
