package Models

import (
	"fmt"
	"golang-klinik/Config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type RiwayatPeriksa struct {
	Id                 int                  `json:"id" gorm:"primary_key"`
	HasilPeriksa       string               `json:"hasil_periksa"`
	RujukanKe          string               `json:"rujukan_ke"`
	TglPeriksa         time.Time            `json:"tgl_periksa"`
	TotalBayar         int                  `json:"total_bayar"`
	IsBPJS             int                  `json:"is_bpjs"`
	Status             string               `json:"status"`
	KlinikID           int                  `json:"klinik_id" gorm:"column:klinik_id"`
	Klinik             Klinik               `json:"klinik"`
	PasienID           int                  `json:"pasien_id" gorm:"column:pasien_id"`
	Pasien             Pasien               `json:"pasien"`
	PoliID             int                  `json:"poli_id" gorm:"column:poli_id"`
	Poli               Poli                 `json:"poli"`
	DokterID           int                  `json:"dokter_id" gorm:"column:dokter_id"`
	Dokter             Dokter               `json:"dokter"`
	RiwayatPeriksaObat []RiwayatPeriksaObat `json:"riwayat_obat" gorm:"ForeignKey:RiwayatPeriksaID"`
}

func (b *RiwayatPeriksa) TableName() string {
	return "riwayat_periksa"
}

func RiwayatPeriksaGetAll(riwayat_periksa *[]RiwayatPeriksa) (err error) {
	if err = Config.DB.Preload("Klinik").Preload("Pasien").Preload("Poli").Preload("Dokter").Preload("RiwayatPeriksaObat").Preload("RiwayatPeriksaObat.Obat").Find(riwayat_periksa).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaCreate(riwayat_periksa *RiwayatPeriksa) (err error) {
	if err = Config.DB.Create(riwayat_periksa).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaFindOne(riwayat_periksa *RiwayatPeriksa, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Preload("Klinik").Preload("Pasien").Preload("Poli").Preload("Dokter").Preload("RiwayatPeriksaObat").Preload("RiwayatPeriksaObat.Obat").First(riwayat_periksa).Error; err != nil {
		return err
	}
	return nil
}

func RiwayatPeriksaUpdateOne(riwayat_periksa *RiwayatPeriksa, id string) (err error) {
	fmt.Println(riwayat_periksa)
	Config.DB.Save(riwayat_periksa)
	return nil
}

func RiwayatPeriksaDeleteOne(riwayat_periksa *RiwayatPeriksa, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(riwayat_periksa)
	return nil
}
