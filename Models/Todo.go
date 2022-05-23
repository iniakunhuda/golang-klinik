package Models

import (
	"fmt"
	"golang-klinik/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAll(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func Create(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func FindOne(todo *Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOne(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func DeleteOne(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
