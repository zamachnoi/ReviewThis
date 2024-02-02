package handlers

import (
	"log"
	"net/http"
	"os"

	dbData "github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/util"
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

    jwt, err := util.GenerateDiscordIDJWT(newUserInfo.DiscordID)
    if err != nil {
        log.Printf("Error generating JWT: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    handleRedirect(w, r, jwt)
}

func handleRedirect(w http.ResponseWriter, r *http.Request, jwt string) {
    util.SetJWTCookie(jwt, w)

    //todo FIX THIS
    http.Redirect(w, r, os.Getenv("CLIENT_REDIRECT_URL"), http.StatusFound)
}
