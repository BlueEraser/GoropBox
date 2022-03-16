package endpoints

import (
	"net/http"

	ihttp "github.com/inhun/GoropBox/internal/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

func (e *Endpoints) Oauth2Google(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	url := e.Oauth.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

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
