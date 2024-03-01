package util

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/zamachnoi/viewthis/models"
)

const (
    DiscordTokenURL = "https://discord.com/api/v10/oauth2/token"
    DiscordUserURL  = "https://discord.com/api/users/@me"
)

func GetNewToken(code string, grantType string, redirectUri string) (*models.DiscordTokenResponse, error){
    data := url.Values{
        "client_id":     {os.Getenv("DISCORD_CLIENT_ID")},
        "client_secret": {os.Getenv("DISCORD_CLIENT_SECRET")},
        "grant_type":    {grantType},
        "code":          {code},
        "redirect_uri":  {redirectUri},
        "scope":         {"identify"},
    }

     // get access tokens from discord
    discordTokenReq, err := http.PostForm(DiscordTokenURL, data)
    if err != nil {
        return nil, err
    }

    var discordTokenBody models.DiscordTokenResponse
    if err := json.NewDecoder(discordTokenReq.Body).Decode(&discordTokenBody); err != nil {
         return nil, err
    }
    log.Printf("DiscordTokenReq: %v", discordTokenReq.Body)

    return &discordTokenBody, nil
    
}

func GetDiscordUserData(accessToken string, refreshToken string) (*models.User, error) {
        // set up request to get user data
        discordUserReq, err := http.NewRequest("GET", DiscordUserURL, nil)
        if err != nil {
            return nil, err
        }
        
        // set authorization header
        discordUserReq.Header.Set("Authorization", "Bearer "+ accessToken)
    
        // get user data
        client := &http.Client{}
        discordDataResponse, err := client.Do(discordUserReq)
        if err != nil {
            return nil, err
        }
        defer discordDataResponse.Body.Close()
    
        // encode to struct
        var discordUser models.DiscordUser
        if err := json.NewDecoder(discordDataResponse.Body).Decode(&discordUser); err != nil {
            return nil, err
        }
    
        // encode to db model
        newUserInfo, err := EncodeDiscordUserInfo(discordUser, refreshToken)
        if err != nil {
            return nil, err
        }

        return newUserInfo, nil
}