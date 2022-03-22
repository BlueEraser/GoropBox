package auth

import (
	"context"

	iconfig "github.com/inhun/GoropBox/config"

	"google.golang.org/api/idtoken"

	"github.com/mitchellh/mapstructure"

	"golang.org/x/oauth2"
)

type TokenInfo struct {
	Sub           string
	Email         string
	AtHash        string `mapstructure:"at_hash"`
	Aud           string
	EmailVerified bool `mapstructure:"email_verified"`
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Local         string
	Iss           string
	Azp           string
	Iat           int64
	Exp           int64
}

func LoadAuthConfig(cfg iconfig.GoogleConfig) *oauth2.Config {
	OauthConf := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  cfg.RedirectUrl,
		Endpoint: oauth2.Endpoint{
			TokenURL: cfg.TokenUrl,
			AuthURL:  cfg.AuthUrl,
		},
	}

	return OauthConf
}

func AuthGoogle(idToken string, ClientId string) (*TokenInfo, error) {
	ctx := context.Background()

	v, err := idtoken.Validate(ctx, idToken, ClientId)
	if err != nil {
		return nil, err
	}

	var tokenInfo TokenInfo
	err = mapstructure.Decode(v.Claims, &tokenInfo)
	if err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}
