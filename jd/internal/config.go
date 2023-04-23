package internal

type SystemConfig struct {
	Method      string `json:"method" url:"method"`
	AppKey      string `json:"app_key" url:"app_key"`
	AccessToken string `json:"access_token,omitempty" url:"access_token"`
	Timestamp   string `json:"timestamp" url:"timestamp"`
	Format      string `json:"format" url:"format"`
	Version     string `json:"v" url:"v"`
	SignMethod  string `json:"sign_method" url:"sign_method"`
	Sign        string `json:"sign" url:"sign"`
	ParamJson   string `json:"360buyParamJson"  url:"360buy_param_json"`
}
