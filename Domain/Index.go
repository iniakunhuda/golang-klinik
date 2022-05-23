package Domain

import (
	"fmt"
	"golang-klinik/Helpers"
	"golang-klinik/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func ViewAccountRegister(c *gin.Context) {
	// check auth
	if Helpers.AuthIsLoggedIn(c) {
		c.Redirect(http.StatusFound, "/")
	}
	var data = map[string]interface{}{
		"title": "Login",
	}
	c.HTML(http.StatusOK, "register.html", data)
}

func RegisterNewAccount(c *gin.Context) {
	name := c.PostForm("nama")
	tgl_lahir := c.PostForm("tgl_lahir")
	bpjs := c.PostForm("bpjs")
	provinsi := c.PostForm("provinsi")
	kota := c.PostForm("kota")
	kodepos := c.PostForm("kodepos")
	alamat := c.PostForm("alamat")
	riwayat_penyakit := c.PostForm("riwayat")
	hp := c.PostForm("hp")
	password := c.PostForm("password")

	// create user
	pasien := Models.Pasien{
		Nama:            name,
		TglLahir:        tgl_lahir,
		NoBpjs:          bpjs,
		Provinsi:        provinsi,
		Kota:            kota,
		Kodepos:         kodepos,
		Alamat:          alamat,
		RiwayatPenyakit: riwayat_penyakit,
		Hp:              hp,
		Password:        Helpers.Encrypt(password, Helpers.GetKey()),
	}
	err := Models.PasienCreate(&pasien)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func ViewAccountLogin(c *gin.Context) {
	// check auth
	if Helpers.AuthIsLoggedIn(c) {
		c.Redirect(http.StatusFound, "/")
	} else {
		var data = map[string]interface{}{
			"title": "Login",
		}
		c.HTML(http.StatusOK, "login.html", data)
	}
}

func AccountLogin(c *gin.Context) {
	sess_key := Helpers.LoadEnv("SESSION_KEY")
	var store = sessions.NewCookieStore([]byte(sess_key))
	session, _ := store.Get(c.Request, "session")

	if session.Values["loggedin"] == "true" {
		c.Redirect(http.StatusFound, "/")
	} else {
		// check db for user
		var pasien Models.Pasien

		hp := c.PostForm("hp")
		pass := c.PostForm("pass")

		err := Models.PasienFindByHp(&pasien, hp)
		if err != nil {
			c.Redirect(http.StatusFound, "/account/login")
		} else {
			// check password
			if Helpers.Decrypt(pasien.Password, Helpers.GetKey()) == pass {
				session.Values["loggedin"] = "true"
				session.Values["pasien_id"] = pasien.Id
				session.Values["hp"] = pasien.Hp
				session.Save(c.Request, c.Writer)
				c.Redirect(http.StatusFound, "/")
			} else {
				c.Redirect(http.StatusFound, "/account/login")
			}
		}
	}
}

func AccountLogout(c *gin.Context) {
	sess_key := Helpers.LoadEnv("SESSION_KEY")
	var store = sessions.NewCookieStore([]byte(sess_key))
	session, _ := store.Get(c.Request, "session")
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		fmt.Println("failed to delete session", err)
	}

	c.Redirect(http.StatusFound, "/")
}

func ViewIndex(c *gin.Context) {
	var klinik []Models.Klinik
	var poli, pilihan_polis []Models.Poli
	var filter Models.FilterPoli

	err := Models.KlinikGetAll(&klinik)

	klinik_id, _ := strconv.Atoi(c.Query("klinik_id"))
	poli_id, _ := strconv.Atoi(c.Query("poli_id"))

	filter = Models.FilterPoli{
		Klinik: (klinik_id),
		Poli:   (poli_id),
	}
	poli, _ = Models.PoliGetAntrian(&poli, &filter)
	Models.PoliGetAll(&pilihan_polis)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		var data = map[string]interface{}{
			"title":         "Homepage",
			"kliniks":       klinik,
			"pilihan_polis": pilihan_polis,
			"polis":         poli,
			"filter":        filter,
			"pasien_login":  Helpers.AuthPasienDetail(c),
		}

		c.HTML(http.StatusOK, "index.html", data)
	}
}

func ViewRegister(c *gin.Context) {
	var klinik Models.Klinik
	var poli Models.Poli
	var antrian []Models.Poli

	if Helpers.AuthPasienDetail(c) == nil {
		c.Redirect(http.StatusFound, "/account/login")
	}

	klinik_id, _ := c.Params.Get("klinik_id")
	poli_id, _ := c.Params.Get("poli_id")

	k, _ := strconv.Atoi(klinik_id)
	p, _ := strconv.Atoi(poli_id)

	filter := Models.FilterPoli{
		Klinik: k,
		Poli:   p,
	}

	antrian, _ = Models.PoliGetAntrian(&antrian, &filter)
	err := Models.KlinikFindOne(&klinik, klinik_id)
	err = Models.PoliFindOne(&poli, poli_id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var data = map[string]interface{}{
			"title":        "Antri Klinik",
			"klinik":       klinik,
			"antrian":      antrian,
			"poli":         poli,
			"pasien_login": Helpers.AuthPasienDetail(c),
		}

		c.HTML(http.StatusOK, "antri.html", data)
	}
}

func CreateRegister(c *gin.Context) {
	name := c.PostForm("nama")
	tgl_lahir := c.PostForm("tgl_lahir")
	bpjs := c.PostForm("bpjs")
	provinsi := c.PostForm("provinsi")
	kota := c.PostForm("kota")
	kodepos := c.PostForm("kodepos")
	alamat := c.PostForm("alamat")
	riwayat_penyakit := c.PostForm("riwayat")

	klinik_id, _ := strconv.Atoi(c.PostForm("klinik_id"))
	poli_id, _ := strconv.Atoi(c.PostForm("poli_id"))
	no_urutan, _ := strconv.Atoi(c.PostForm("antrian"))

	currentTime := time.Now()
	tgl := currentTime.Day()
	bln := int(currentTime.Month())
	thn := currentTime.Year()

	var pasienId int

	fmt.Println(c.PostForm("pasien_id"))
	if c.PostForm("pasien_id") == "" {
		// create user
		pasien := Models.Pasien{
			Nama:            name,
			TglLahir:        tgl_lahir,
			NoBpjs:          bpjs,
			Provinsi:        provinsi,
			Kota:            kota,
			Kodepos:         kodepos,
			Alamat:          alamat,
			RiwayatPenyakit: riwayat_penyakit,
		}
		err := Models.PasienCreate(&pasien)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		pasienId = pasien.Id
	} else {
		// update user
		pasienId, _ = strconv.Atoi(c.PostForm("pasien_id"))
		pasien := Models.Pasien{
			Id:              pasienId,
			Nama:            name,
			TglLahir:        tgl_lahir,
			NoBpjs:          bpjs,
			Provinsi:        provinsi,
			Kota:            kota,
			Kodepos:         kodepos,
			Alamat:          alamat,
			RiwayatPenyakit: riwayat_penyakit,
		}
		err := Models.PasienUpdateOne(&pasien, c.PostForm("pasien_id"))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
	}

	// register ke antrian
	antrian := Models.Antrian{
		KlinikID: klinik_id,
		PoliID:   poli_id,
		PasienID: pasienId,
		Tanggal:  tgl,
		Bulan:    bln,
		Tahun:    thn,
		Urutan:   no_urutan,
	}
	err := Models.AntrianCreate(&antrian)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusFound, "/")
	}

}
