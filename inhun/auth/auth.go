package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/inhun/GoropBox/config"

	"golang.org/x/oauth2"
)

func A(cfg config.Config) {
	var OauthConf = &oauth2.Config{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  cfg.Google.RedirectUrl,
		Endpoint: oauth2.Endpoint{
			TokenURL: cfg.Google.TokenUrl,
			AuthURL:  cfg.Google.AuthUrl,
		},
	}

	url := OauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("%v", url)

	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString((b))
	fmt.Println(state)

	fmt.Println(OauthConf.AuthCodeURL(state))
}
