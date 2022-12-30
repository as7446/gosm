package config

import (
	"fmt"
	"github.com/snail007/go-activerecord/mysql"
	"gosm/model"
	"os"
)

var (
	defaultConfFile = "~/.gosm/config.yml"
	configFile      = "config.yml"
	Conf            Configuration
)

func init() {
	initConfig()
	initDB()
}

func initConfig() {
	c, err := ReadConf(configFile)
	if err != nil {
		return
	}
	Conf = c
}

func initDB() {
	dbHost := Conf.Server.DBHost
	dbPort := Conf.Server.DBPort
	dbName := Conf.Server.DBName
	dbUser := Conf.Server.DBUser
	dbPassword := Conf.Server.DBPassword
	maxIdle := Conf.Server.ConnMaxIdle
	maxConn := Conf.Server.ConnMaxConnection
	model.G = mysql.NewDBGroup("default")
	if dbPort == 0 {
		dbPort = 3306
	}
	cfg := mysql.NewDBConfigWith(dbHost, dbPort, dbName, dbUser, dbPassword)
	if maxIdle > 0 && maxConn > 0 {
		cfg.MaxIdleConns = maxConn
		cfg.MaxIdleConns = maxIdle
	}
	err := model.G.Regist("default", cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
