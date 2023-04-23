package internal

type IRequest interface {
	JsonParams() ([]byte, error)
	ResponseName() string
	RequestMethod() string
}

type SystemRequest struct {
	Method      string `json:"method" url:"method"` //API接口名称 jd.union.open.order.query
	AppKey      string `json:"app_key"`             //联盟分配给应用的appkey，可在应用查看中获取appkey
	AccessToken string `json:"access_token"`        //根据API属性标签，如果需要授权，则此参数必传;如果不需要授权，则此参数不需要传
	Timestamp   string `json:"timestamp"`           //时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8。API服务端允许客户端请求最大时间误差为10分钟
	Format      string `json:"format"`              //json
	Version     string `json:"v"`                   //API协议版本，请根据API具体版本号传入此参数，一般为1.0
	SignMethod  string `json:"sign_method"`         //签名的摘要算法，暂时只支持md5
}
