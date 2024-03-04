package repositories

import (
	"github.com/josebruno2020/go-devbook-api/src/db"
	"github.com/josebruno2020/go-devbook-api/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := db.Connect()
	return &UserRepository{db}
}

func (c UserRepository) List() ([]models.User, error) {
	var users []models.User
	result := c.db.Find(&users)
	return users, result.Error
}

func (c UserRepository) Create(user models.User) (uint, error) {
	result := c.db.Create(&user)
	return user.ID, result.Error
}

func (c UserRepository) GetById(id int) (models.User, error) {
	var user models.User
	result := c.db.First(&user, id)
	return user, result.Error
}

func (c UserRepository) UpdateById(user models.User) error {
	result := c.db.Save(&user)
	return result.Error
}

func (c UserRepository) DeleteById(id int) error {
	var user models.User
	result := c.db.Delete(&user, id)
	return result.Error
}
