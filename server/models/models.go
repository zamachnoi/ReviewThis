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
    Queues   []Queue `gorm:"foreignKey:UserID"`
}


type DiscordUser struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Avatar   string `json:"avatar"`
    RefreshToken string `json:"refresh_token"`
    RefreshExpiry time.Time `json:"refresh_expiry"`
}


type Queue struct {
    gorm.Model
    Name        string
    Description string
    Type        string
    UserID      uint
    Submissions []Submission `gorm:"foreignKey:QueueID;onDelete:CASCADE"`
}

type Submission struct {
    gorm.Model
    Content     string
    UserID      uint
    QueueID     uint
    Feedbacks   []Feedback `gorm:"foreignKey:SubmissionID;onDelete:CASCADE"`
}

type Feedback struct {
    gorm.Model
    Content     string
    UserID      uint
    Submission  Submission
    SubmissionID uint
}