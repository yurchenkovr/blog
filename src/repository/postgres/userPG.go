package postgres

import (
	"blog/src/models"
	"github.com/go-pg/pg"
	"log"
	"time"
)

type UserRepository interface {
	SaveUser(models.User) error
	GetByIDUser(int) (*models.User, error)
	DeleteUser(int) error
	GetAllUsers() ([]models.User, error)
	UpdateUser(int, UpdateUser) error
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *pg.DB
}
type UpdateUser struct {
	Username  string
	Password  string
	UpdatedAt time.Time
}

func (a *userRepository) UpdateUser(id int, art UpdateUser) error {
	var updatedUser models.User
	if _, err := a.db.Model(&updatedUser).Set("username = ?, password = ?, updated_at = ?", art.Username, art.Password, art.UpdatedAt).
		Where("id = ?", id).Update(); err != nil {
		log.Printf("Error while updating User, Reason: %v\n", err)
		return err
	}
	return nil
}
func (a *userRepository) SaveUser(user models.User) error {
	if err := a.db.Insert(&user); err != nil {
		log.Printf("Error while inserting new User into DB, Reason: %v\n", err)
		return err
	}
	return nil
}

/*func (a *userRepository) GetByIDUser(id int) (models.User, error) {
	var user models.User
	if err := a.db.Model(&user).Where("id = ?", id).First(); err != nil {
		log.Printf("Error while GetByIDUser, Reason: %v\n", err)
		return models.User{}, err
	}
	return user, nil
}*/
func (a *userRepository) GetByIDUser(id int) (*models.User, error) {
	var user = new(models.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name"
	FROM "users" AS "user" JOIN "roles" AS "role" ON "role"."id" = "user"."role_id"
	WHERE ("user"."id" = ?)`
	_, err := a.db.QueryOne(user, sql, id)
	if err != nil {
		log.Printf("ERROR while GetByIDUser, Reason: %v\n", err)
		return nil, err
	}
	return user, nil
}
func (a *userRepository) DeleteUser(id int) error {
	var user models.User
	if _, err := a.db.Model(&user).Where("id = ?", id).Delete(); err != nil {
		log.Printf("Error while deleting User, Reason: %v\n", err)
		return err
	}
	return nil
}
func (a *userRepository) GetAllUsers() ([]models.User, error) {
	var user []models.User
	if err := a.db.Model(&user).Select(); err != nil {
		log.Printf("Error while trying to Select All Users, Reason: %v\n", err)
		return nil, err
	}
	return user, nil
}
