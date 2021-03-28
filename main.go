package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gosturcture/repository"
	"gosturcture/rest"
)

func initConf() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	env := viper.GetString("use")
	viper.SetConfigName(env)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDatasource() {
	host := viper.GetString("datasource.pg.host")
	port := viper.GetInt("datasource.pg.port")
	db := viper.GetString("datasource.pg.db")
	user := viper.GetString("datasource.pg.user")
	pass := viper.GetString("datasource.pg.pass")
	conInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		db,
	)
	pg,err := sql.Open("postgres",conInfo)
	if err != nil {
		panic(err)
	}
	repository.Pg = pg
}

func main() {
	initConf()
	initDatasource()
	host := viper.GetString("app.host")
	port := viper.GetInt("app.port")
	app := fiber.New()
	addr := fmt.Sprintf("%s:%d",host,port)
	rest.Router(app)
	if err := app.Listen(addr); err != nil {
		fmt.Println(err)
	}
}
