package auth

import (
	"context"

	"google.golang.org/api/idtoken"

	"github.com/mitchellh/mapstructure"
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
