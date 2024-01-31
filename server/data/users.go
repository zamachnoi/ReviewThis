// data/user.go
package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := db.GetDB().Preload("Queues").First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func CreateUser(user models.User) (*models.User, error) {
    if err := db.GetDB().Create(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(user models.User) (*models.User, error) {
	if err := db.GetDB().Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(user models.User) error {
	if err := db.GetDB().Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}
