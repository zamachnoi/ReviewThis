package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	dbData "github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
	"github.com/zamachnoi/viewthis/util"
)

type DiscordResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
}
const (
    DiscordTokenURL = "https://discord.com/api/v10/oauth2/token"
    DiscordUserURL  = "https://discord.com/api/users/@me"
)


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

    data := url.Values{
        "client_id":     {os.Getenv("DISCORD_CLIENT_ID")},
        "client_secret": {os.Getenv("DISCORD_CLIENT_SECRET")},
        "grant_type":    {"authorization_code"},
        "code":          {code},
        "redirect_uri":  {os.Getenv("DISCORD_REDIRECT_URI")},
        "scope":         {"identify"},
    }
    
    resp, err := http.PostForm(DiscordTokenURL, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    defer resp.Body.Close()
    
    var discordAuthResponse DiscordResponse
    if err := json.NewDecoder(resp.Body).Decode(&discordAuthResponse); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
        
    // set up request to get user data
    discordRequest, err := http.NewRequest("GET", DiscordUserURL, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // set authorization header
    discordRequest.Header.Set("Authorization", "Bearer "+ discordAuthResponse.AccessToken)
    
    // get user data
    client := &http.Client{}
    discordDataResponse, err := client.Do(discordRequest)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer discordDataResponse.Body.Close()

	// encode to struct
    var discordUser models.DiscordUser
    if err := json.NewDecoder(discordDataResponse.Body).Decode(&discordUser); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// check if user exists in database
    newUserInfo, err := util.EncodeDiscordUserInfo(discordUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// update the user in the database
    if _, err := dbData.UpdateUser(*newUserInfo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return // returns if error
    }

    jwt, err := util.GenerateDiscordJWT(newUserInfo.DiscordID)
    if err != nil {
        log.Printf("Error generating JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Generated JWT: %s", jwt)

    expiry := util.GetJWTExpiry()
    log.Printf("JWT Expiry: %s", expiry)

    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    jwt,
        Expires:  expiry,
        Path:    "/", // set to root so it's accessible from all pages

    })

    // redirect to frontend
    log.Printf("redirecting to %s", os.Getenv("CLIENT_REDIRECT_URL"))
    http.Redirect(w, r, os.Getenv("CLIENT_REDIRECT_URL"), http.StatusFound)
}