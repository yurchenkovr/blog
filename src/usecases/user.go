package usecases

import (
	"blog/src/infrastructure/secure"
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
	GetByUsername(string) (*models.User, error)
	BlockUser(int) error
	UnblockUser(int) error
	//IsAdmin(int) bool
	//DeleteAnyUser(int, int) error
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
	hash := secure.HashAndSalt([]byte(req.Password))

	user := models.User{
		Base: models.Base{
			CreatedAt: time.Now(),
		},
		Username: req.Username,
		Password: hash,
		RoleID:   req.RoleID,
		Blocked:  false,
	}
	if err := s.userRep.SaveUser(user); err != nil {
		log.Printf("error SU, Reason: %v\n", err)
		return err
	}
	return nil
}
func (s userService) UpdateUser(id int, req models.User) error {
	hash := secure.HashAndSalt([]byte(req.Password))

	updUser := models.User{
		Base: models.Base{
			UpdatedAt: time.Now(),
		},
		Username: req.Username,
		Password: hash,
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
func (s userService) GetByUsername(username string) (*models.User, error) {
	user, err := s.userRep.GetByUsername(username)
	if err != nil {
		log.Printf("error GBU, Reason: %v\n", err)
		return user, err
	}
	return user, nil
}
func (s userService) BlockUser(id int) error {
	if err := s.userRep.UpdateStatus(id, true); err != nil {
		log.Printf("Error Block: %v\n", err)
		return err
	}
	return nil
}
func (s userService) UnblockUser(id int) error {
	if err := s.userRep.UpdateStatus(id, false); err != nil {
		log.Printf("Error Block: %v\n", err)
		return err
	}
	return nil
}

/*
func (s userService) DeleteAnyUser(clamsID, id int) error {
	if s.IsAdmin(clamsID) == true {
		if err := s.DeleteUser(id); err != nil {
			return err
		}
		return nil
	}
	return errors.New("you`re not an Admin")
}*/
/*func (s userService) IsAdmin(id int) bool {
	user, err := s.GetByIDUser(id)
	if err != nil {
		fmt.Println("something go bad..when checking IsAdmin")
		return false
	}
	if user.Role.AccessLevel == models.AdminRole {
		return true
	}
	return false
}
*/
