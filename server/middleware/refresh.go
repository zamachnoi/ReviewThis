package middleware

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zamachnoi/viewthis/util"
)

func RefreshJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get _viewthis_jwt 
		jwtToken, err := r.Cookie("_viewthis_jwt")
		if err != nil {
			next.ServeHTTP(w, r) // Call the next middleware or handler
			return
		}    
		token, claims, err := util.ParseJWTClaims(jwtToken.Value)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				_, _ = handleExpiredJWT(token, claims, w, r, next)
			}
		}
		next.ServeHTTP(w, r) // Call the next middleware or handler
	})
}