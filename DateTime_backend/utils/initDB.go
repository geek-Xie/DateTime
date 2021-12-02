package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func InitDB() *gorm.DB  {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password,host,port,database,charset)
	db, error := gorm.Open(driverName, args)
	if error != nil {
		panic("fail to connect database!! error: " + error.Error())
	}
	return db
}
