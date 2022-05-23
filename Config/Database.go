package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func BuildConfig() *DBConfig {
	config := DBConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root_user",
		Password: "oEqMRGWilr9oNUgu",
		Database: "sister_klinik",
	}
	return &config
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}
