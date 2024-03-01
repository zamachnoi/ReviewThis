package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	dbData "github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/util"
)

type CookieValueResponse struct {
    Status string `json:"status"`
    JWT    string `json:"jwt,omitempty"` // omit if empty to avoid confusion in error responses
}

// redirect to callback
func DiscordAuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DISCORD_OAUTH_URL")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func DiscordBotAddHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := util.GetJWTValue(r)
    _, claims, err := util.ParseJWTClaims(tokenString)
    if err != nil {
        log.Printf("Error parsing JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    premium, err := dbData.GetPremiumUser(int(claims.DBID))
    if err != nil || !premium {
        http.Error(w, "You are not a premium user", http.StatusUnauthorized)
        return
    }
    url := os.Getenv("DISCORD_OAUTH_BOT_URL")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// callback from discord
func DiscordAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// get code from url query
	code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    // get access tokens from discord
    discordTokenBody, err := util.GetNewToken(code, "authorization_code", os.Getenv("DISCORD_REDIRECT_URI"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // get user data from discord with token
    newUserInfo, err := util.GetDiscordUserData(discordTokenBody.AccessToken, discordTokenBody.RefreshToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

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
        Premium:   true,
    }

    // generate a JWT with user data
    jwt, err := util.GenerateSessionJWT(sessionJWT)
    if err != nil {
        log.Printf("Error generating JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    handleRedirect(w, r, jwt, os.Getenv("CLIENT_REDIRECT_URL"))
}

func DiscordBotCallbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    log.Println("Code: ", code)
    tokenString := util.GetJWTValue(r)
    log.Println("Token: ", tokenString)
    _, claims, err := util.ParseJWTClaims(tokenString)
    if err != nil {
        log.Printf("Error parsing JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    premium, err := dbData.GetPremiumUser(int(claims.DBID))
    if err != nil || !premium {
        log.Printf("Not Premium: %v", premium)
        http.Error(w, "You are not a premium user", http.StatusUnauthorized)
        return
    }

    discordTokenBody, err := util.GetNewToken(code, "authorization_code", os.Getenv("DISCORD_BOT_REDIRECT_URI"))
    if err != nil {
        log.Printf("Error getting new token: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Printf("Discord Token: %v", discordTokenBody)

    guildId := discordTokenBody.GuildResponse.GuildID
    if guildId == "" {
        log.Printf("Guild ID not found")
        http.Error(w, "Guild ID not found", http.StatusBadRequest)
        return
    }

    name := discordTokenBody.GuildResponse.Name
    if name == "" {
        log.Printf("Guild name not found")
        http.Error(w, "Guild name not found", http.StatusBadRequest)
        return
    }

    err = dbData.CreateGuild(guildId, claims.DBID, name)
    if err != nil {
        log.Printf("Error creating guild: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    redirectUrl := os.Getenv("CLIENT_REDIRECT_URL")
    redirectUrl += "/users/" + strconv.Itoa(int(claims.DBID)) + "?bot_success=true"
    handleRedirect(w, r, tokenString, redirectUrl)
}

// get cookie/jwt
func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
    cookies := r.Cookies()

    // Print all cookies
    for _, cookie := range cookies {
        log.Printf("Received cookie: %s = %s\n", cookie.Name, cookie.Value)
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

func handleRedirect(w http.ResponseWriter, r *http.Request, jwt string, redirectURL string) {
    util.SetJWTCookie(jwt, w)

    
    http.Redirect(w, r, redirectURL, http.StatusFound)
}

func DiscordAuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
    util.ExpireCookie("_viewthis_jwt", w)
    http.Redirect(w, r, os.Getenv("CLIENT_REDIRECT_URL"), http.StatusFound)
}
