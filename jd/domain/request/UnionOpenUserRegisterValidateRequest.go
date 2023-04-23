package request

import (
	"encoding/json"
)

type UserStateRequest struct {
	UserId     string `json:"userId"`
	UserIdType int    `json:"userIdType"`
}

type UnionOpenUserRegisterValidateRawRequest struct {
	UserStateReq *UserStateRequest `json:"userStateReq"`
}

func NewUnionOpenUserRegisterValidateRowQueryRequest(userStateRequest *UserStateRequest) *UnionOpenUserRegisterValidateRawRequest {
	return &UnionOpenUserRegisterValidateRawRequest{
		UserStateReq: userStateRequest,
	}
}

func (req *UnionOpenUserRegisterValidateRawRequest) JsonParams() ([]byte, error) {
	paramJsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return paramJsonBytes, nil
}

func (req *UnionOpenUserRegisterValidateRawRequest) ResponseName() string {
	return "jd_union_open_user_register_validate"
}

func (req *UnionOpenUserRegisterValidateRawRequest) RequestMethod() string {
	return "jd.union.open.user.register.validate"
}
