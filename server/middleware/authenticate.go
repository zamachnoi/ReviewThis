package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
	"github.com/zamachnoi/viewthis/util"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        jwtCookie, err := r.Cookie("_viewthis_jwt")
        if err != nil {
            log.Printf("No token found in cookie: %v", err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        token, claims, err := util.ParseJWTClaims(jwtCookie.Value)
        if err != nil {
            if errors.Is(err, jwt.ErrTokenExpired) {
                token, err = handleExpiredJWT(token, claims, w, r, next)
                if err != nil {
                    log.Printf("Error handling expired JWT: %v", err)
                    http.Error(w, "Unauthorized", http.StatusUnauthorized)
                    return
                }
            } else {
                log.Printf("Error parsing JWT claims: %v %T", err, err)
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
        }
        
        if !token.Valid {
            log.Printf("Token is not valid")
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        log.Printf("IN AUTH MIDDLEWARE")
        ctx := context.WithValue(r.Context(), util.UserKey, claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}


func handleRefreshToken(claims util.SessionJWTWithClaims) ( error) {
    user, err := data.GetUserByDiscordID(claims.DiscordID)
    if err != nil {
        return err
    }
    refreshExpiry := user.RefreshExpiry
    if time.Until(refreshExpiry) < time.Hour*72 {
        return refreshTokenInsertUser(claims.DiscordID, user.RefreshToken)
    } else if time.Until(refreshExpiry) < 0 {
        return errors.New("refresh token expired")
    }

    return nil
}

func refreshTokenInsertUser(discordId string, encrypedRefreshToken string) ( error){
    newTokens, err := util.GetNewToken(encrypedRefreshToken, "refresh_token")
    if err != nil {
        return err
    }

    err = UpdateUserWithDiscordData(newTokens.AccessToken, newTokens.RefreshToken)
    if err != nil {
        return err
    }

    return nil
}

func UpdateUserWithDiscordData(newAccessToken string, newRefreshToken string) (error) {
    updatedUser, err := util.GetDiscordUserData(newAccessToken, newRefreshToken)
    if err != nil {
        return err
    }
    _, err = data.UpdateUser(*updatedUser)
    return err
}

func UpdateDBUserWithDiscordData(newAccessToken string, newRefreshToken string) (*models.User, error) {
    newUser, err := util.GetDiscordUserData(newAccessToken, newRefreshToken)
    if err != nil {
        return nil, err
    }
    return data.UpdateUser(*newUser)

}

func handleExpiredJWT(token *jwt.Token, claims util.SessionJWTWithClaims, w http.ResponseWriter, r *http.Request, next http.Handler) (*jwt.Token, error) {
    err := handleRefreshToken(claims)
    if err != nil {
        log.Printf("Error handling refresh token: %v", err)
        return nil, err
    }
    
    newTokenString, err := util.GenerateSessionJWT(claims.SessionJWT)
    if err != nil {
        log.Printf("Error here.")
        return nil, err
    }

    util.SetJWTCookie(newTokenString, w)

    newTokenWithClaims, _, err := util.ParseJWTClaims(newTokenString)
    if err != nil {
        return nil, err
    }

    return newTokenWithClaims, nil
}