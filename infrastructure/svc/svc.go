package svc

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	mlog "github.com/itmrchow/microservice-gateway/infrastructure/log"
)

var (
	_accountConn *grpc.ClientConn

	accountLocation string
	accountOptions  []grpc.DialOption
)

func InitAccLocation(url string) {
	accountLocation = url

	jsonConfig, err := os.ReadFile("grpc-service-config.json")
	if err != nil {
		mlog.Fatal().Err(err).Msg("read grpc-service-config.json error")
	}

	grpc.WithDisableRetry()

	accountOptions = append(accountOptions, grpc.WithDefaultServiceConfig(string(jsonConfig)), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
