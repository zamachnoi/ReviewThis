// data/user.go
package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models" // import models
)

func GetUser(id uint) (*models.User, error) {
    var user models.User // use models.User
    if err := db.GetDB().First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}