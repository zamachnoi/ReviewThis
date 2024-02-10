package models

import (
	"time"

	"gorm.io/gorm"
)
type UserSessionData struct {
    DbID uint `json:"db_id"`
    DiscordID string `json:"discord_id"`
    Username string `json:"username"`
    Avatar string `json:"avatar"`
    RefreshExpiry time.Time `json:"refresh_expiry"`
}

type DiscordUser struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Avatar   string `json:"avatar"`
}

type DiscordTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
}

type User struct {
    gorm.Model
    Username      string    `json:"username"`
    DiscordID     string    `gorm:"uniqueIndex" json:"discord_id"`
    Avatar        string    `json:"avatar"`
    RefreshToken  string    `json:"refresh_token"`
    RefreshExpiry time.Time `json:"refresh_expiry"`
    AccessExpiry  time.Time `json:"access_expiry"`
    Queues        []Queue   `gorm:"foreignKey:UserID" json:"queues"`
}

type Queue struct {
    gorm.Model
    Name        string       `json:"name"`
    Description string       `json:"description"`
    Type        string       `sql:"type:enum('soundcloud');default:'soundcloud'" json:"type"`
    DiscordID   string       `json:"discord_id"`
    UserID      uint         `json:"user_id"`
    Username    string       `json:"username"`
    Avatar      string       `json:"avatar"`
    Private     bool         `json:"private"`
    Submissions []Submission `gorm:"foreignKey:QueueID;onDelete:CASCADE" json:"submissions"`
}

type Submission struct {
    gorm.Model
    Content   string     `json:"content"`
    UserID    uint       `json:"user_id"`
    Username  string     `json:"username"`
    Avatar    string     `json:"avatar"`
    QueueID   uint       `json:"queue_id"`
    Private   bool       `json:"private"`
    Feedbacks []Feedback `gorm:"foreignKey:SubmissionID;onDelete:CASCADE" json:"feedbacks"`
}

type Feedback struct {
    gorm.Model
    Content      string     `json:"content"`
    UserID       uint       `json:"user_id"`
    Submission   Submission `json:"submission"`
    SubmissionID uint       `json:"submission_id"`
}