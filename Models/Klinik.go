package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Klinik struct {
	Id           int    `json:"id" gorm:"primary_key"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	Telp         string `json:"telp"`
	TahunBerdiri string `json:"tahun_berdiri"`
	Polis        []Poli `json:"polis" gorm:"ForeignKey:KlinikID"`
}

func (b *Klinik) TableName() string {
	return "klinik"
}

func KlinikGetAll(klinik *[]Klinik) (err error) {
	if err = Config.DB.Preload("Polis").Find(klinik).Error; err != nil {
		return err
	}
	return nil
}

func KlinikCreate(klinik *Klinik) (err error) {
	if err = Config.DB.Create(klinik).Error; err != nil {
		return err
	}
	return nil
}

func KlinikFindOne(klinik *Klinik, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(klinik).Error; err != nil {
		return err
	}
	return nil
}

func KlinikUpdateOne(klinik *Klinik, id string) (err error) {
	fmt.Println(klinik)
	Config.DB.Save(klinik)
	return nil
}

func KlinikDeleteOne(klinik *Klinik, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(klinik)
	return nil
}
