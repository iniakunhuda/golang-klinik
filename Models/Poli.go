package Models

import (
	"fmt"
	"golang-klinik/Config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Poli struct {
	Id           int    `json:"id" gorm:"primary_key"`
	Nama         string `json:"nama"`
	NamaKlinik   string `json:"nama_klinik"`
	AlamatKlinik string `json:"alamat_klinik"`
	KlinikID     int    `json:"klinik_id" gorm:"column:klinik_id"`
	Klinik       Klinik `json:"klinik"`
	TotalAntrian int    `json:"total_antrian"`
}

type FilterPoli struct {
	Klinik int
	Poli   int
}

func (b *Poli) TableName() string {
	return "poli"
}

func PoliGetAntrian(poli *[]Poli, filter *FilterPoli) (pl []Poli, err error) {

	sql := "SELECT p.*, k.nama as nama_klinik, k.alamat as alamat_klinik, IFNULL(a.total, 0) as total_antrian FROM poli p JOIN klinik k ON p.klinik_id = k.id	LEFT JOIN (SELECT a.poli_id, count(1) total FROM antrian a GROUP BY a.poli_id) as a ON a.poli_id = p.id"

	if filter.Klinik != 0 || filter.Poli != 0 {

		if filter.Klinik != 0 && filter.Poli != 0 {
			sql += " WHERE p.klinik_id = '" + strconv.Itoa(filter.Klinik) + "' AND p.id = '" + strconv.Itoa(filter.Poli) + "'"
		} else if filter.Klinik != 0 {
			sql += " WHERE p.klinik_id = '" + strconv.Itoa(filter.Klinik) + "'"
		} else if filter.Poli != 0 {
			sql += " WHERE p.id = '" + strconv.Itoa(filter.Poli) + "'"
		}
	}

	sql += "  ORDER BY 1"

	var polis []Poli
	rows, err := Config.DB.Raw(sql).Rows()

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	for rows.Next() {
		var p Poli
		err = Config.DB.ScanRows(rows, &p)
		if err != nil {
			fmt.Printf("Unable to scan the row. %v", err)
		}

		polis = append(polis, p)
	}

	// return empty user on error
	return polis, err
}

func PoliGetAll(poli *[]Poli) (err error) {
	if err = Config.DB.Preload("Klinik").Find(poli).Error; err != nil {
		return err
	}
	return nil
}

func PoliCreate(poli *Poli) (err error) {
	if err = Config.DB.Create(poli).Error; err != nil {
		return err
	}
	return nil
}

func PoliFindOne(poli *Poli, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Preload("Klinik").First(poli).Error; err != nil {
		return err
	}
	return nil
}

func PoliUpdateOne(poli *Poli, id string) (err error) {
	fmt.Println(poli)
	Config.DB.Save(poli)
	return nil
}

func PoliDeleteOne(poli *Poli, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(poli)
	return nil
}
