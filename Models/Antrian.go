package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Antrian struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Tanggal  int    `json:"tanggal"`
	Bulan    int    `json:"bulan"`
	Tahun    int    `json:"tahun"`
	Urutan   int    `json:"urutan"`
	KlinikID int    `json:"klinik_id" gorm:"column:klinik_id"`
	Klinik   Klinik `json:"klinik"`
	PasienID int    `json:"pasien_id" gorm:"column:pasien_id"`
	Pasien   Pasien `json:"pasien"`
	PoliID   int    `json:"poli_id" gorm:"column:poli_id"`
	Poli     Poli   `json:"poli"`
}

func (b *Antrian) TableName() string {
	return "antrian"
}

func SetAntrian(antrian *Antrian) (err error) {
	Config.DB.Create(antrian)
	return nil
}

func AntrianGetAll(antrian *[]Antrian) (err error) {
	if err = Config.DB.Preload("Klinik").Preload("Pasien").Preload("Poli").Find(antrian).Error; err != nil {
		return err
	}
	return nil
}

func AntrianCreate(antrian *Antrian) (err error) {
	if err = Config.DB.Create(antrian).Error; err != nil {
		return err
	}
	return nil
}

func AntrianFindOne(antrian *Antrian, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Preload("Klinik").Preload("Pasien").Preload("Poli").First(antrian).Error; err != nil {
		return err
	}
	return nil
}

func AntrianUpdateOne(antrian *Antrian, id string) (err error) {
	fmt.Println(antrian)
	Config.DB.Save(antrian)
	return nil
}

func AntrianDeleteOne(antrian *Antrian, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(antrian)
	return nil
}
