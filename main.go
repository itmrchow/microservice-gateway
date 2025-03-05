package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/itmrchow/microservice-gateway/delivery/handlers"
	mlog "github.com/itmrchow/microservice-gateway/infrastructure/log"
	"github.com/itmrchow/microservice-gateway/infrastructure/svc"
)

func main() {
	initConfig()
	initBase()
	initLog()
	initSubService()
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

	log.Info().Msgf("config init success")
}

// initBase 初始化系統基礎設定
func initBase() {
	// time zone
	time.Local = time.UTC

	log.Info().Msgf("base init success")
}

// initLog 初始化log
func initLog() {
	mlog.InitLog(mlog.LogSettingInfo{
		LogLevelStr: viper.GetString("log_level"),
		Output:      viper.GetString("log_output"),
		File:        viper.GetString("log_file"),
		Dir:         viper.GetString("log_dir"),
		ServerName:  viper.GetString("server_name"),
	})

	mlog.Info().Msgf("log init success")
}

// initSubService 初始化子服務
func initSubService() {
	svc.InitAccLocation(viper.GetString("service_address.account_grpc"))
	mlog.Info().Msgf("sub service init success")
}

// initRouter 初始化路由
func initRouter() {
	var (
		// publicPort   = viper.GetString("http_public_port")
		internalPort = viper.GetString("http_internal_port")
	)

	mux := handlers.RegisterPublicHandlers()
	mlog.Info().Msgf("http internal server listen in port " + internalPort)
	mlog.Fatal().AnErr("error", http.ListenAndServe(":"+internalPort, mux))
}
