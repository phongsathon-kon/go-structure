package a

import "github.com/spf13/viper"

var Conf *viper.Viper

func GetAppName() string {
	return Conf.GetString("app.name")
}
