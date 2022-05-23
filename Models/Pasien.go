package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Pasien struct {
	Id              int    `json:"id" gorm:"primary_key"`
	NoBpjs          string `json:"no_bpjs"`
	Nama            string `json:"nama"`
	Alamat          string `json:"alamat"`
	Provinsi        string `json:"provinsi"`
	Kota            string `json:"kota"`
	Kodepos         string `json:"kodepos"`
	RiwayatPenyakit string `json:"riwayat_penyakit"`
	TglLahir        string `json:"tgl_lahir"`
	FotoDiri        string `json:"foto_diri"`
	FotoKtp         string `json:"foto_ktp"`
	FotoKtpDiri     string `json:"foto_ktp_diri"`
	Hp              string `json:"hp"`
	Password        string `json:"password"`
}

func (b *Pasien) TableName() string {
	return "pasien"
}

func PasienGetAll(pasien *[]Pasien) (err error) {
	if err = Config.DB.Find(pasien).Error; err != nil {
		return err
	}
	return nil
}

func PasienCreate(pasien *Pasien) (err error) {
	if err = Config.DB.Create(pasien).Error; err != nil {
		return err
	}
	return nil
}

func PasienFindOne(pasien *Pasien, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(pasien).Error; err != nil {
		return err
	}
	return nil
}

func PasienFindByHp(pasien *Pasien, hp string) (err error) {
	if err = Config.DB.Where("hp = ?", hp).First(pasien).Error; err != nil {
		return err
	}
	return nil
}

func PasienUpdateOne(pasien *Pasien, id string) (err error) {
	fmt.Println(pasien)
	Config.DB.Save(pasien)
	return nil
}

func PasienDeleteOne(pasien *Pasien, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(pasien)
	return nil
}
