package imp

import "github.com/twistycs/pokemon-go-backend/models"

type UserService interface {
	GetAllUser() (userList []models.User, err error)
	GetUserByUserName(userId string) (user models.User, err error)
	InsertUser(u *models.User) (err error)
	UpdateUser(u *models.User, id string) (err error)
	DeleteUserById(id string) (err error)
}

type UserRepository interface {
	GetAllUser(u *[]models.User) (err error)
	GetUserByUserName(u *models.User, userId string) (err error)
	InsertUser(u *models.User) (err error)
	UpdateUser(u *models.User) (err error)
	DeleteUserById(id string) (err error)
}
