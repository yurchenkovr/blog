package postgres

import (
	"blog/src/models"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/gommon/log"
	"time"
)

func New() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "vitmantekoR_2408",
		Database: "simpleblog",
		Addr:     "localhost:5432",
	})
	if db == nil {
		log.Printf("Failed to connect to database!\n")
	}
	fmt.Printf("Connection to database successful.\n")

	CreateTableArticle(db)

	return db
}
func SaveUserTest(userRep userRepository) {
	newUser := models.User{
		Base: models.Base{
			ID:        6,
			CreatedAt: time.Now(),
			//UpdatedAt: nil,
		},
		Username: "mantekor",
		Password: "ttt333",
		Role:     new(models.Role),
		RoleID:   100,
	}
	err := userRep.SaveUser(newUser)
	if err != nil {
		log.Printf("Error testing saving User, Reason: %v\n", err)
	}
	fmt.Println("Saving user successfully")
}

func DeleteUserTest(userRep userRepository) {
	err := userRep.DeleteUser(5)
	if err != nil {
		log.Printf("Error while testing deleting User, Reason: %v\n", err)
	}
	fmt.Println("Deleting User successfully")
}
func UpdateUserTest(userRep userRepository) {
	updUser := UpdateUser{
		Username:  "Jamie",
		Password:  "zxc123",
		UpdatedAt: time.Now(),
	}
	err := userRep.UpdateUser(2, updUser)
	if err != nil {
		log.Printf("Error while Testing UpdateUSer, Reason: %v\n", err)
	}
	fmt.Println("Update Successfully")
}
func CreateTableArticle(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&models.Article{}, opts)
	if err != nil {
		log.Printf("Error while creating table Article, Reason: %v\n", err)
	}
}
