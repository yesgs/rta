package response

type UserRegisterValidateResponse struct {
	UserResp UserResp `json:"userResp"`
}
type UserResp struct {
	JdUser int `json:"jdUser"`
}

type UnionOpenUserRegisterValidateRowResponse struct {
	Code      int                           `json:"code"`
	Message   string                        `json:"message"`
	RequestId string                        `json:"requestId"`
	Data      *UserRegisterValidateResponse `json:"data"`
}
