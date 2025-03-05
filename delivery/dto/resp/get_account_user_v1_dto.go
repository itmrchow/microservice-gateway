package resp

import (
	accountV1 "github.com/itmrchow/microservice-proto/account/v1"
)

type GetAccountUserV1Resp struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (r *GetAccountUserV1Resp) FromProto(proto *accountV1.GetUserResponse) {
	r.Id = proto.Id
	r.Username = proto.Username
}
