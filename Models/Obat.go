package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Obat struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Nama     string `json:"nama"`
	Harga    int    `json:"harga"`
	Penyakit string `json:"penyakit"`
	Stok     int    `json:"stok"`
}

func (b *Obat) TableName() string {
	return "obat"
}

func ObatGetAll(obat *[]Obat) (err error) {
	if err = Config.DB.Find(obat).Error; err != nil {
		return err
	}
	return nil
}

func ObatCreate(obat *Obat) (err error) {
	if err = Config.DB.Create(obat).Error; err != nil {
		return err
	}
	return nil
}

func ObatFindOne(obat *Obat, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(obat).Error; err != nil {
		return err
	}
	return nil
}

func ObatUpdateOne(obat *Obat, id string) (err error) {
	fmt.Println(obat)
	Config.DB.Save(obat)
	return nil
}

func ObatDeleteOne(obat *Obat, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(obat)
	return nil
}
