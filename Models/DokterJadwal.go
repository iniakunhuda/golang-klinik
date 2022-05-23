package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type DokterJadwal struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Hari       string `json:"hari"`
	JamMulai   int    `json:"jam_mulai"`
	JamSelesai int    `json:"jam_selesai"`
	IsActive   int    `json:"is_active"`
	DokterID   int    `json:"dokter_id" gorm:"column:dokter_id"`
	Dokter     Dokter `json:"dokter"`
	KlinikID   int    `json:"klinik_id" gorm:"column:klinik_id"`
	Klinik     Klinik `json:"klinik"`
	PoliID     int    `json:"poli_id" gorm:"column:poli_id"`
	Poli       Poli   `json:"poli"`
}

func (b *DokterJadwal) TableName() string {
	return "dokter_jadwal"
}

func DokterJadwalGetAll(dokter_jadwal *[]DokterJadwal) (err error) {
	if err = Config.DB.Preload("Klinik").Preload("Dokter").Preload("Poli").Find(dokter_jadwal).Error; err != nil {
		return err
	}
	return nil
}

func DokterJadwalCreate(dokter_jadwal *DokterJadwal) (err error) {
	if err = Config.DB.Create(dokter_jadwal).Error; err != nil {
		return err
	}
	return nil
}

func DokterJadwalFindOne(dokter_jadwal *DokterJadwal, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(dokter_jadwal).Error; err != nil {
		return err
	}
	return nil
}

func DokterJadwalUpdateOne(dokter_jadwal *DokterJadwal, id string) (err error) {
	fmt.Println(dokter_jadwal)
	Config.DB.Save(dokter_jadwal)
	return nil
}

func DokterJadwalDeleteOne(dokter_jadwal *DokterJadwal, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(dokter_jadwal)
	return nil
}
