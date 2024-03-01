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
type GuildResponse struct {
    GuildID string `json:"id"`
    Name    string `json:"name"`
}

type DiscordTokenResponse struct {
	AccessToken     string  `json:"access_token"`
	TokenType       string  `json:"token_type"`    
	ExpiresIn       int32   `json:"expires_in"`
	RefreshToken    string  `json:"refresh_token"`
	Scope           string  `json:"scope"`
    GuildResponse   GuildResponse `json:"guild"`
}

type User struct {
    gorm.Model
    Username      string    `json:"username"`
    DiscordID     string    `gorm:"uniqueIndex" json:"discord_id"`
    Avatar        string    `json:"avatar"`
    RefreshToken  string    `json:"refresh_token"`
    RefreshExpiry time.Time `json:"refresh_expiry"`
    Premium       bool      `json:"premium"`
    Queues        []Queue   `gorm:"foreignKey:UserID" json:"queues"`
    Guilds        []Guild   `gorm:"many2many:guild_users;" json:"guilds"`
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
    Name      string     `json:"name"`
    Content   string     `json:"content"`
    UserID    uint       `json:"user_id"`
    Username  string     `json:"username"`
    Avatar    string     `json:"avatar"`
    QueueID   uint       `json:"queue_id"`
    Private   bool       `json:"private"`
    DiscordID     string    `gorm:"uniqueIndex" json:"discord_id"`
    User         User       `gorm:"foreignKey:UserID"` // Add this line to establish the foreign key relationship
}

type Feedback struct {
    gorm.Model
    Content      string     `json:"content"`
    UserID       uint       `json:"user_id"`
    Submission   Submission `json:"submission"`
    SubmissionID uint       `json:"submission_id"`
    User         User       `gorm:"foreignKey:UserID"` // Add this line to establish the foreign key relationship

}

type Guild struct {
    gorm.Model
    Name     string `json:"name"`
    GuildID   string `gorm:"uniqueIndex" json:"guild_id"`
    OwnerID   uint   `json:"owner_id" gorm:"foreignKey:ID"` // Owner's User ID, references the User model's ID
    Owner     User   `gorm:"foreignKey:OwnerID"` // This line ensures OwnerID is treated as a foreign key
    Users     []User `gorm:"many2many:guild_users;" json:"users"`
}

type GuildUser struct {
    UserID     uint `gorm:"primaryKey" json:"user_id"`
    GuildID    uint `gorm:"primaryKey" json:"guild_id"` 
    Authorized bool `json:"authorized" gorm:"default:false"`
    User       User `gorm:"foreignKey:UserID;references:ID"`
    Guild      Guild `gorm:"foreignKey:GuildID;references:ID"`
}