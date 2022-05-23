package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Dokter struct {
	Id        int    `json:"id" gorm:"primary_key"`
	NoStr     string `json:"no_str"`
	Nama      string `json:"nama"`
	Spesialis string `json:"spesialis"`
}

func (b *Dokter) TableName() string {
	return "dokter"
}

func DokterGetAll(dokter *[]Dokter) (err error) {
	if err = Config.DB.Find(dokter).Error; err != nil {
		return err
	}
	return nil
}

func DokterCreate(dokter *Dokter) (err error) {
	if err = Config.DB.Create(dokter).Error; err != nil {
		return err
	}
	return nil
}

func DokterFindOne(dokter *Dokter, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(dokter).Error; err != nil {
		return err
	}
	return nil
}

func DokterUpdateOne(dokter *Dokter, id string) (err error) {
	fmt.Println(dokter)
	Config.DB.Save(dokter)
	return nil
}

func DokterDeleteOne(dokter *Dokter, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(dokter)
	return nil
}
