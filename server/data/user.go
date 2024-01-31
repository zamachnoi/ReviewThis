// data/user.go
package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := db.GetDB().First(&user, id).Error; err != nil {
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