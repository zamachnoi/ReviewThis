package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/zamachnoi/viewthis/data"
)

type DiscordResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
}

type DiscordUser struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Avatar   string `json:"avatar"`
}

func DiscordAuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DISCORD_OAUTH_URL")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

const (
    DiscordTokenURL = "https://discord.com/api/oauth2/token"
    DiscordUserURL  = "https://discord.com/api/users/@me"
)

func DiscordAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// get code from url query
	code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

	// get access token from discord
    auth, err := http.PostForm(DiscordTokenURL, url.Values{
        "client_id":     {os.Getenv("DISCORD_CLIENT_ID")},
        "client_secret": {os.Getenv("DISCORD_CLIENT_SECRET")},
        "grant_type":    {"authorization_code"},
        "code":          {code},
        "redirect_uri":  {os.Getenv("DISCORD_REDIRECT_URI")},
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer auth.Body.Close()

    var discordAuthResponse DiscordResponse
    if err := json.NewDecoder(auth.Body).Decode(&discordAuthResponse); err != nil {
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
    discordRequest.Header.Set("Authorization", "Bearer "+discordAuthResponse.AccessToken)

	// get user data
    client := &http.Client{}
    discordDataResponse, err := client.Do(discordRequest)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer discordDataResponse.Body.Close()

	// encode to struct
    var discordUser DiscordUser
    if err := json.NewDecoder(discordDataResponse.Body).Decode(&discordUser); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// check if user exists in database
    newUserInfo, err := data.GetUserByDiscordID(discordUser.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// set the user's username and avatar
    newUserInfo.Username = discordUser.Username
    newUserInfo.Avatar = discordUser.Avatar
    newUserInfo.DiscordID = discordUser.ID

	// update the user in the database
    if _, err := data.UpdateUser(*newUserInfo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUserInfo)
}
