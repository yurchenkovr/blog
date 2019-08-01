package usecases

import (
	"blog/src/infrastructure/secure"
	"blog/src/models"
	"blog/src/repository/postgres"
	"blog/src/usecases/rbac"
	"errors"
	"github.com/labstack/echo"
	"log"
	"time"
)

type UserService interface {
	Create(CreateReqUser) error
	View(int) (*models.User, error)
	Delete(echo.Context, int) error
	List() ([]models.User, error)
	Update(echo.Context, int, models.User) error
	GetByUsername(string) (*models.User, error)
	Block(echo.Context, int) error
	Unblock(echo.Context, int) error
	Login(username, password string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(models.User) (string, error)
}

type userService struct {
	userRep        postgres.UserRepository
	tokenGenerator TokenGenerator
	rbac           rbac.RBAC
}

func NewUserService(userRep postgres.UserRepository, generator TokenGenerator, rbac rbac.RBAC) UserService {
	return &userService{userRep: userRep, tokenGenerator: generator, rbac: rbac}
}

type CreateReqUser struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	RoleID   models.AccessRole `json:"role_id"`
}

func (s userService) Login(username, password string) (string, error) {
	user, err := s.GetByUsername(username)
	if err != nil {
		log.Printf("error: %s", err)
		return "", err
	}

	if !secure.ComparePasswords(user.Password, []byte(password)) {
		return "", errors.New("password or login is incorrect")
	}

	return s.tokenGenerator.GenerateToken(*user)
}

func (s userService) Create(req CreateReqUser) error {
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
	if err := s.userRep.Create(user); err != nil {
		log.Printf("error SU, Reason: %v\n", err)
		return err
	}
	return nil
}

func (s userService) Update(c echo.Context, id int, req models.User) error {
	if !s.rbac.EnforceUser(c, id) {
		return errors.New("It`s not your user or you`re not an admin\n")
	}

	hash := secure.HashAndSalt([]byte(req.Password))

	updUser := models.User{
		Base: models.Base{
			UpdatedAt: time.Now(),
		},
		Username: req.Username,
		Password: hash,
	}
	if err := s.userRep.Update(id, updUser); err != nil {
		log.Printf("error UU, Reason: %v\n", err)
		return err
	}

	return nil
}

func (s userService) List() ([]models.User, error) {
	users, err := s.userRep.List()
	if err != nil {
		log.Printf("error GU, Reason: %v\n", err)
		return nil, err
	}

	return users, nil
}

func (s userService) Delete(c echo.Context, id int) error {
	if !s.rbac.EnforceUser(c, id) {
		return errors.New("It`s not your user or you`re not an admin\n")
	}

	if err := s.userRep.Delete(id); err != nil {
		log.Printf("error DU, Reason: %v\n", err)
		return err
	}

	return nil
}

func (s userService) View(id int) (*models.User, error) {
	user, err := s.userRep.View(id)
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

func (s userService) Block(c echo.Context, id int) error {
	if !s.rbac.IsAdmin(c) {
		return errors.New("You`re not an admin!\n")
	}

	if err := s.userRep.UpdateStatus(id, true); err != nil {
		log.Printf("Error Block: %v\n", err)
		return err
	}

	return nil
}

func (s userService) Unblock(c echo.Context, id int) error {
	if !s.rbac.IsAdmin(c) {
		return errors.New("You`re not an admin!\n")
	}

	if err := s.userRep.UpdateStatus(id, false); err != nil {
		log.Printf("Error Block: %v\n", err)
		return err
	}

	return nil
}
