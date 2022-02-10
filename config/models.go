package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

//file, err := ioutil.TempFile("", "webd")
//if err != nil {
//log.Fatal(err)
//}
//log.Println(file.Name())
//defer func(name string) {
//	err := os.Remove(name)
//	if err != nil {
//		log.Fatal(err)
//	}
//}(file.Name())

func GetUserHome() string {
	var homeDir, _ = os.UserHomeDir()
	return homeDir
}

func GetCurrentDir() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return getwd
}

func GetSystemRoot() string {
	return viper.GetString("system_root")
}

func GetListenAddress() string {
	return viper.GetString("listen_address")
}

func GetServerHeader() string {
	return viper.GetString("server_header")
}

func GeScripts() map[string]string {
	return viper.GetStringMapString("scripts")
}

func init() {
	viper.AddConfigPath("/etc/webhookd")
	viper.AddConfigPath(path.Join(GetUserHome(), ".webhookd"))
	viper.AddConfigPath(GetCurrentDir())

	viper.SetConfigName("webhookd_config")
	viper.SetDefault("system_root", ".")
	viper.SetDefault("listen_address", ":3002")
	viper.SetDefault("server_header", "Server")
	viper.SetDefault("scripts", map[string]string{})
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("Fatal error config file: %w \n", err))

	}
}
