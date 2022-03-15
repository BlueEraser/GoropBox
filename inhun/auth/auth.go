package auth

import (
	"context"
	"fmt"
	"log"

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

	url := OauthConf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	fmt.Printf("%v\n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tok, err := OauthConf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tok)
	client := OauthConf.Client(ctx, tok)

	fmt.Println(client)

}
