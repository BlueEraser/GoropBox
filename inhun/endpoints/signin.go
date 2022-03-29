package endpoints

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/inhun/GoropBox/internal/auth"

	"github.com/julienschmidt/httprouter"
)

func (e *Endpoints) CallbackGoogle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := r.URL.Query().Get("code")

	ctx := context.Background()
	tok, err := e.Oauth.Exchange(ctx, code)
	fmt.Println(tok)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(tok.Extra("id_token"))
	a, _ := auth.AuthGoogle(fmt.Sprintf("%v", tok.Extra("id_token")), e.Oauth.ClientID)
	fmt.Println(a)

	// fmt.Println(tok.Extra("id_token"))
	// client := e.Oauth.Client(ctx, tok)

}
