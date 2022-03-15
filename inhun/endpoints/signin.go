package endpoints

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (e *Endpoints) Signin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	code := r.URL.Query().Get("code")
	fmt.Println(code)

	ctx := context.Background()
	tok, err := e.Oauth.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tok)
	client := e.Oauth.Client(ctx, tok)
	fmt.Println(client)
}
