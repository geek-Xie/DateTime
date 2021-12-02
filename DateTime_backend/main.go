package main

import (
	"DateTime_backend/route"
	"DateTime_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initConfig()
	r := gin.Default()
	r = route.CollectRoutes(r)
	db := utils.InitDB()
	defer db.Close()
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}
}

func initConfig()  {
	 workDir, _ := os.Getwd()
	 viper.SetConfigName("application")
	 viper.SetConfigType("yml")
	 viper.AddConfigPath(workDir + "/config")
	 err := viper.ReadInConfig()
	 if err != nil {
	 	panic(err)
	 }
}
