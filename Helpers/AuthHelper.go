package Helpers

import (
	"golang-klinik/Models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func AuthIsLoggedIn(c *gin.Context) bool {
	var store = sessions.NewCookieStore([]byte(LoadEnv("SESSION_KEY")))
	session, _ := store.Get(c.Request, "session")
	return session.Values["loggedin"] == "true"
}

func AuthPasienDetail(c *gin.Context) interface{} {
	var store = sessions.NewCookieStore([]byte(LoadEnv("SESSION_KEY")))
	session, _ := store.Get(c.Request, "session")

	if session.Values["pasien_id"] == nil {
		return nil
	}

	pasien_id := strconv.Itoa(session.Values["pasien_id"].(int))

	var pasien Models.Pasien
	Models.PasienFindOne(&pasien, pasien_id)
	return pasien
}
