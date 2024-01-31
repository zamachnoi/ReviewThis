package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

func GetUserByDiscordID(id string) (*models.User, error) {
	var user models.User
	if err := db.GetDB().Where("discord_id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return &models.User{}, err
	}
	return &user, nil
}

