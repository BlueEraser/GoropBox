package endpoints

import (
	"fmt"
	"net/http"

	ihttp "github.com/inhun/GoropBox/internal/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

func (e *Endpoints) Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	url := e.Oauth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("%v\n", url)

	ihttp.ResponseOK(w, "success", url)
	return

	/*
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
	*/
}
