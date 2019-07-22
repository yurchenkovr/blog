package usecases

import (
	"blog/src/models"
	"blog/src/repository/postgres"
	"log"
	"time"
)

type UserService interface {
	SaveUser(CreateReqUser) error
	GetByIDUser(int) (*models.User, error)
	DeleteUser(int) error
	GetAllUsers() ([]models.User, error)
	UpdateUser(int, models.User) error
}
type userService struct {
	userRep postgres.UserRepository
}

func NewUserService(userRep postgres.UserRepository) UserService {
	return &userService{userRep: userRep}
}

type CreateReqUser struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	RoleID   models.AccessRole `json:"role_id"`
}

func (s userService) SaveUser(req CreateReqUser) error {
	user := models.User{
		Base: models.Base{
			CreatedAt: time.Now(),
		},
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.RoleID,
	}
	if err := s.userRep.SaveUser(user); err != nil {
		log.Printf("error SU, Reason: %v\n", err)
	}
	return nil
}
func (s userService) UpdateUser(id int, req models.User) error {
	updUser := models.User{
		Base: models.Base{
			UpdatedAt: time.Now(),
		},
		Username: req.Username,
		Password: req.Password,
	}
	if err := s.userRep.UpdateUser(id, updUser); err != nil {
		log.Printf("error UU, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s userService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRep.GetAllUsers()
	if err != nil {
		log.Printf("error GU, Reason: %v\n", err)
		return nil, err
	}
	return users, nil
}
func (s userService) DeleteUser(id int) error {
	if err := s.userRep.DeleteUser(id); err != nil {
		log.Printf("error DU, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s userService) GetByIDUser(id int) (*models.User, error) {
	user, err := s.userRep.GetByIDUser(id)
	if err != nil {
		log.Printf("error GIU, Reason: %v\n", err)
		return user, err
	}
	return user, nil
}
