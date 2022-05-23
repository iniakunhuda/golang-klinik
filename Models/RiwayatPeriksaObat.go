package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type RiwayatPeriksaObat struct {
	Id               int  `json:"id" gorm:"primary_key"`
	HargaSatuan      int  `json:"harga_satuan"`
	Quantity         int  `json:"quantity"`
	Subtotal         int  `json:"subtotal"`
	RiwayatPeriksaID int  `json:"riwayat_periksa_id" gorm:"column:riwayat_periksa_id"`
	ObatID           int  `json:"obat_id" gorm:"column:obat_id"`
	Obat             Obat `json:"obat"`
}

func (b *RiwayatPeriksaObat) TableName() string {
	return "riwayat_periksa_obat"
}

func RiwayatPeriksaObatGetAll(riwayat_periksa_obat *[]RiwayatPeriksaObat) (err error) {
	if err = Config.DB.Find(riwayat_periksa_obat).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaObatCreate(riwayat_periksa_obat *RiwayatPeriksaObat) (err error) {
	if err = Config.DB.Create(riwayat_periksa_obat).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaObatFindOne(riwayat_periksa_obat *RiwayatPeriksaObat, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(riwayat_periksa_obat).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaObatUpdateOne(riwayat_periksa_obat *RiwayatPeriksaObat, id string) (err error) {
	fmt.Println(riwayat_periksa_obat)
	Config.DB.Save(riwayat_periksa_obat)
	return nil
}

func RiwayatPeriksaObatDeleteOne(riwayat_periksa_obat *RiwayatPeriksaObat, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(riwayat_periksa_obat)
	return nil
}
