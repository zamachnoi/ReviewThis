package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string
    DiscordID string `gorm:"uniqueIndex"`
    Avatar string
    RefreshToken string
    RefreshExpiry time.Time
    AccessExpiry time.Time
    Queues   []Queue `gorm:"foreignKey:UserID"`
}

type UserSessionData struct {
    DbID uint `json:"db_id"`
    Username string `json:"username"`
    Avatar string `json:"avatar"`
    RefreshExpiry time.Time `json:"refresh_expiry"`
}

type DiscordUser struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Avatar   string `json:"avatar"`
}

type Queue struct {
    gorm.Model
    Name        string
    Description string
    Type        string `sql:"type:enum('soundcloud');default:'soundcloud'"`
    UserID      uint
    IsPrivate   bool
    Submissions []Submission `gorm:"foreignKey:QueueID;onDelete:CASCADE"`
}

type Submission struct {
    gorm.Model
    Content     string
    UserID      uint
    QueueID     uint
    IsPrivate   bool
    Feedbacks   []Feedback `gorm:"foreignKey:SubmissionID;onDelete:CASCADE"`
}

type Feedback struct {
    gorm.Model
    Content      string
    UserID       uint
    Submission   Submission
    SubmissionID uint
}

type DiscordTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
}