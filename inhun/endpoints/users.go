package endpoints

import (
	"fmt"
	"net/http"

	ihttp "github.com/inhun/GoropBox/internal/http"
	"github.com/inhun/GoropBox/models"

	"github.com/julienschmidt/httprouter"
)

func (e *Endpoints) GetUserList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User
	e.DB.Find(&user)
	fmt.Println(user)

	ihttp.ResponseOK(w, "success", user)
	return
}
