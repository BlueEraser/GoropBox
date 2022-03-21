package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"phobyjun/config"
)

var (
	store = sessions.NewCookieStore([]byte(config.Cfg.Session.SECRET))
)

const (
	sessionValue = "user_sess"
)

func Init(e *echo.Echo) {
	e.Use(session.Middleware(store))
}

func Get(c echo.Context) *sessions.Session {
	sess, _ := session.Get(sessionValue, c)
	return sess
}

func Save(c echo.Context, email string) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	sess.Values["email"] = email
	return saveSession(c, sess)
}

func Delete(c echo.Context) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	sess.Values["email"] = nil
	return saveSession(c, sess)
}

func saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return nil
}
