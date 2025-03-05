package svc

import (
	"fmt"

	accountV1 "github.com/itmrchow/microservice-proto/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// NewAccount 初始化
func NewAccountUserSvcV1() (accountV1.UserServiceClient, error) {
	var err error
	if _accountConn == nil {
		_accountConn, err = grpc.NewClient(accountLocation, accountOptions...)
		if err != nil {
			return nil, err
		}
	}
	//check state: if(state == TransientFailure/Shutdown)-> re conn
	if _accountConn.GetState() != connectivity.Idle && _accountConn.GetState() != connectivity.Ready {
		//close old conn
		_accountConn.Close()
		_accountConn, err = grpc.NewClient(accountLocation, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {

			return nil, err
		}
	}

	return accountV1.NewUserServiceClient(_accountConn), nil
}
