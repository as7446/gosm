package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Configuration struct {
	Server struct {
		HttpAuth struct {
			Enable   bool   `yaml:"enable"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
		Addr              string `yaml:"addr"`
		LogOutfile        string `yaml:"logOutfile"`
		Mode              string `yaml:"mode"`
		DBHost            string `yaml:"db_host"`
		DBPort            int    `yaml:"db_port"`
		DBName            string `yaml:"db_name"`
		DBUser            string `yaml:"db_user"`
		DBPassword        string `yaml:"db_password"`
		ConnMaxIdle       int    `yaml:"conn_max_idle"`
		ConnMaxConnection int    `yaml:"conn_max_connection"`
	}
	Client struct {
		ServerUrl string `yaml:"serverUrl"`
	}
}

func IsExits(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func IsFile(path string) bool {
	return !IsDir(path)
}
func OpenConfig(filename string) (c Configuration) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		data = []byte("")
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}
	return
}
func ReadConf(filename string) (c Configuration, err error) {
	c.Server.Addr = ":9000"
	c.Client.ServerUrl = "http://localhost:9000"
	c.Server.Mode = "dev"
	if !IsExits(filename) {
		us, _ := user.Current()
		defaultConfFile = filepath.Join(us.HomeDir, defaultConfFile)
		if IsExits(defaultConfFile) && IsFile(defaultConfFile) {
			c = OpenConfig(defaultConfFile)
		} else {
			cfgDir := filepath.Join(us.HomeDir, ".gosm")
			if !IsDir(cfgDir) {
				os.MkdirAll(cfgDir, 0755)
			}
			data, _ := yaml.Marshal(c)
			ioutil.WriteFile(defaultConfFile, data, 0644)
		}
	} else {
		c = OpenConfig(filename)
	}
	return
}
