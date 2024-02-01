package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
	"github.com/zamachnoi/viewthis/util"
)

type RequestContext struct {
    User  *models.User
    Mutex sync.RWMutex
}

type ContextKey string

func (c ContextKey) String() string {
    return string(c)
}

var requestContextKey = ContextKey("requestContext")

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        reqContext := &RequestContext{}

        tokenCookie, err := r.Cookie("token")
        if err != nil {
            log.Printf("No token found in cookie: %v", err)
            reqContext.Mutex.Lock()
            reqContext.User = nil
            reqContext.Mutex.Unlock()
            ctx := context.WithValue(r.Context(), requestContextKey, reqContext)
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }

        claims := util.DiscordClaims{}

        token, err := jwt.ParseWithClaims(tokenCookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            log.Printf("Error parsing token: %v", err)
            reqContext.Mutex.Lock()
            reqContext.User = nil
            reqContext.Mutex.Unlock()
            ctx := context.WithValue(r.Context(), requestContextKey, reqContext)
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }

        user, err := data.GetUserByDiscordID(claims.DiscordID)
        if err != nil {
            log.Printf("Error getting user by discord ID: %v", err)
            reqContext.Mutex.Lock()
            reqContext.User = nil
            reqContext.Mutex.Unlock()
            ctx := context.WithValue(r.Context(), requestContextKey, reqContext)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }

        reqContext.Mutex.Lock()
        reqContext.User = user
        reqContext.Mutex.Unlock()
        ctx := context.WithValue(r.Context(), requestContextKey, reqContext)

        reqContext.Mutex.RLock()
        v := reqContext.User
        reqContext.Mutex.RUnlock()

        if v == nil {
            log.Printf("Error getting user from context")
        }
        log.Printf("User from context: %v", v)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}