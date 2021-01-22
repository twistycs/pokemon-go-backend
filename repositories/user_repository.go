package repositories

import (
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"gorm.io/gorm"
)

type userRepository struct {
	connect *gorm.DB
}

func UserRepositories(connect *gorm.DB) imp.UserRepository {
	return &userRepository{connect}
}

func (repo *userRepository) GetAllUser(u *[]models.User) (err error) {
	if err = repo.connect.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUserByUserName(u *models.User, userId string) (err error) {
	if err = repo.connect.Where("LOWER(username = ?)", userId).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) InsertUser(u *models.User) (err error) {
	if err = repo.connect.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) UpdateUser(u *models.User) (err error) {
	if err = repo.connect.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) DeleteUserById(id string) (err error) {
	if err = repo.connect.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
