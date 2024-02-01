package util

import (
	// import go jwt library
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
)

// Create Struct to get Subject from the token
type DiscordClaims struct {
	DiscordID string `json:"discord_id"`
	jwt.RegisteredClaims
}


func GenerateDiscordJWT(subject string) (string, error) {

	claims := DiscordClaims{
		subject,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(GetJWTExpiry()),
		},
	}

	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func GetJWTExpiry() time.Time {
	return time.Now().Add(time.Hour * 72)
}

func EncryptRefreshToken(token string) (string, error) {
    block, err := aes.NewCipher([]byte(os.Getenv("AES_ENCRYPTION_KEY")))
    if err != nil {
        return "", err
    }

    ciphertext := make([]byte, aes.BlockSize+len(token))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(token))

    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptRefreshToken(encryptedToken string) (string, error) {
    block, err := aes.NewCipher([]byte(os.Getenv("AES_ENCRYPTION_KEY")))
    if err != nil {
        return "", err
    }

    decodedToken, err := base64.URLEncoding.DecodeString(encryptedToken)
    if err != nil {
        return "", err
    }

    if len(decodedToken) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv, ciphertext := decodedToken[:aes.BlockSize], decodedToken[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    return string(ciphertext), nil
}

func EncodeDiscordUserInfo(discordUser models.DiscordUser) (*models.User, error) {
    newUserInfo, err := data.GetUserByDiscordID(discordUser.ID)
    if err != nil {
        return nil, err
    }
    if newUserInfo == nil {
        log.Printf("User not found, creating new user")
        newUserInfo = &models.User{}
    }

    // set the user's username and avatar
    newUserInfo.Username = discordUser.Username
    newUserInfo.Avatar = discordUser.Avatar
    newUserInfo.DiscordID = discordUser.ID
    newUserInfo.RefreshToken, err = EncryptRefreshToken(discordUser.RefreshToken)
    if err != nil {
        return nil, err
    }
    newUserInfo.RefreshExpiry = discordUser.RefreshExpiry

    return newUserInfo, nil
}