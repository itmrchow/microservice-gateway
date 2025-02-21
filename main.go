package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/itmrchow/microservice-gateway/delivery/handlers"
	mlog "github.com/itmrchow/microservice-gateway/entities/log"
)

func main() {
	initConfig()
	initBase()
	initLog()
	initRouter()

	time.Local = time.UTC
}

// initConfig 初始化config
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("config init error")
	}
}

// initBase 初始化系統基礎設定
func initBase() {
	// time zone
	time.Local = time.UTC
}

// initLog 初始化log
func initLog() {
	mlog.InitLog()
}

// initRouter 初始化路由
func initRouter() {
	var (
		// publicPort   = viper.GetString("http_public_port")
		internalPort = viper.GetString("http_internal_port")
	)

	mux := handlers.RegisterPublicHandlers()
	mlog.Info().Msg("http internal server listen in port " + internalPort)
	mlog.Fatal().AnErr("error", http.ListenAndServe(":"+internalPort, mux))
}
