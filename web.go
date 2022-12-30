package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosm/config"
	_ "gosm/config"
	"gosm/router"
)

type JSONResponse struct {
	Status int         `json:"status"`
	Value  interface{} `json:"value"`
}

func formatter(params gin.LogFormatterParams) string {
	return ""
}
func Run() {

	r := gin.New()
	gin.LoggerWithConfig(gin.LoggerConfig{Formatter: formatter})
	router.InitRouter(r)
}
func main() {
	//Run()
	conf, err := config.ReadConf(".configs.yml")
	if err != nil {
		return
	}
	fmt.Println(conf)
	fmt.Println(conf.Server.DBPort)
}
