// data/user.go
package data

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/zamachnoi/viewthis/lib"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)
// user data
func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := lib.GetDB().Preload("Queues").First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}
func DeleteUser(user models.User) error {
	if err := lib.GetDB().Delete(&user).Error; err != nil {
		return err
	}

	return lib.DeleteCache("user:" + user.DiscordID)
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := lib.GetDB().Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}
func CreateUser(user models.User) (*models.User, error) {
    if err := lib.GetDB().Create(&user).Error; err != nil {
        return nil, err
    }

    if err := SetUserSessionDataInCache(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

func UpdateUser(user models.User) (*models.User, error) {
    if err := lib.GetDB().Save(&user).Error; err != nil {
        return nil, err
    }

    if err := SetUserSessionDataInCache(&user); err != nil {
        return nil, err
    }

    return &user, nil
}

func GetUserSessionData(discordId string) (*models.UserSessionData, error) {
    // Try to get the user session data from the cache first
    userSessionDataJson, err := lib.GetCache("user:" + discordId)
    if err == nil {
        var userSessionData models.UserSessionData
        json.Unmarshal([]byte(userSessionDataJson), &userSessionData)
        return &userSessionData, nil
    }

    // If the user session data is not in the cache, get the User from the database
    var user models.User
    if err := lib.GetDB().Where("discord_id = ?", discordId).First(&user).Error; err != nil {
        return &models.UserSessionData{}, err
    }

    if err := SetUserSessionDataInCache(&user); err != nil {
        return nil, err
    }

    return UserToSessionData(&user), nil
}

func UserToSessionData(user *models.User) *models.UserSessionData {
    return &models.UserSessionData{
        DbID: user.ID,
        Username: user.Username,
        Avatar: user.Avatar,
        RefreshExpiry: user.RefreshExpiry,
    }
}

func GetUserByDiscordID(discordId string) (*models.User, error) {
    var user models.User
    result := lib.GetDB().Where("discord_id = ?", discordId).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return &models.User{}, nil
        }
        log.Printf("Error getting user by discord ID: %v", result.Error)
        return nil, result.Error
    }
    return &user, nil
}

func SetUserSessionDataInCache(user *models.User) error {
    // Convert the User to UserSessionData
    userSessionData := UserToSessionData(user)

    // Store the user session data in the cache for future use
    userSessionDataBytes, err := json.Marshal(userSessionData)
    if err != nil {
        return err
    }

    return lib.SetCache("user:"+user.DiscordID, string(userSessionDataBytes), time.Hour)
}
