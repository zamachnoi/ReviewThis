package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	dbData "github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/util"
)

type CookieValueResponse struct {
    Status string `json:"status"`
    JWT    string `json:"jwt,omitempty"` // omit if empty to avoid confusion in error responses
}

func DiscordAuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DISCORD_OAUTH_URL")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func DiscordAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// get code from url query
	code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    log.Printf("Code: %s", code)
    // get access tokens from discord
    discordTokenBody, err := util.GetNewToken(code, "authorization_code")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Discord Token Body: %v", discordTokenBody)
    // get user data from discord with token
    newUserInfo, err := util.GetDiscordUserData(discordTokenBody.AccessToken, discordTokenBody.RefreshToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("New User Info: %v", newUserInfo)
	// update the user in the database
    if _, err := dbData.UpdateUser(*newUserInfo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return // returns if error
    }

    // create a SessionJWT
    sessionJWT := util.SessionJWT{
        DiscordID: newUserInfo.DiscordID,
        Avatar:    newUserInfo.Avatar,
        Username:  newUserInfo.Username,
        DBID:      newUserInfo.ID,
    }

    jwt, err := util.GenerateSessionJWT(sessionJWT)
    if err != nil {
        log.Printf("Error generating JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    handleRedirect(w, r, jwt)
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookies := r.Cookies()

    // Print all cookies
    for _, cookie := range cookies {
        log.Printf("Received cookie: %s = %s\n", cookie.Name, cookie.Value)
    }

    // Get a specific cookie
    jwtCookie, err := r.Cookie("_viewthis_jwt")
    if err != nil {
        log.Println("Did not receive _viewthis_jwt cookie")
    } else {
        log.Printf("Received _viewthis_jwt cookie: %s\n", jwtCookie.Value)
    }

    cookie, err := r.Cookie("_viewthis_jwt")
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        if err == http.ErrNoCookie {
            // No cookie found, return an appropriate response
            w.WriteHeader(http.StatusNotFound)
            json.NewEncoder(w).Encode(CookieValueResponse{Status: "error", JWT: "Cookie not found"})
        } else {
            // Handle other errors
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(CookieValueResponse{Status: "error", JWT: err.Error()})
        }
        return
    }
    
    // Cookie found, encode and send the cookie value in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(CookieValueResponse{Status: "success", JWT: cookie.Value})
}

func handleRedirect(w http.ResponseWriter, r *http.Request, jwt string) {
    util.SetJWTCookie(jwt, w)

    //todo FIX THIS
    http.Redirect(w, r, os.Getenv("CLIENT_REDIRECT_URL"), http.StatusFound)
}

func DiscordAuthLogout(w http.ResponseWriter, r *http.Request) {
    util.ExpireCookie("_viewthis_jwt", w)
    http.Redirect(w, r, os.Getenv("CLIENT_REDIRECT_URL"), http.StatusFound)
}
