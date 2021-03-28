package main

import (
	"fmt"
	"gosturcture/services/a"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf-uat")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	a.Conf = viper.GetViper()
	fmt.Println(a.GetAppName())
	name := viper.GetString("app.name")
	version := viper.GetInt("app.version")
	nums := viper.GetIntSlice("app.numbers")
	ss := viper.GetStringSlice("app.strings")
	fmt.Println(name, version, nums, ss)
}
