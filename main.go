package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"my/blogs/common"
	"os"
)

func main() {
	InitConfig() //项目一开始需要读取配置
	db := common.InitDB()
	defer db.Clauses() //todo db.close
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}
