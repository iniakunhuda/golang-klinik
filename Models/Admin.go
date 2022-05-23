package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	KlinikID int    `json:"klinik_id" gorm:"column:klinik_id"`
	Klinik   Klinik `json:"klinik"`
}

func (b *Admin) TableName() string {
	return "admin"
}

func AdminGetAll(admin *[]Admin) (err error) {
	if err = Config.DB.Preload("Klinik").Find(admin).Error; err != nil {
		return err
	}
	return nil
}

func AdminCreate(admin *Admin) (err error) {
	if err = Config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func AdminFindOne(admin *Admin, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Preload("Klinik").First(admin).Error; err != nil {
		return err
	}
	return nil
}

func AdminUpdateOne(admin *Admin, id string) (err error) {
	fmt.Println(admin)
	Config.DB.Save(admin)
	return nil
}

func AdminDeleteOne(admin *Admin, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(admin)
	return nil
}
