package types

// JPConfig ...
type JPConfig struct {
	AppID          string // appId 通信ID，在ADMIN设置
	URL            string // url 瑶池URL
	JadePubKey     string // jadePubKey 验签瑶池公钥
	EccKey         string // eccKey ECC通信私钥
	EccKeyEncoder  string // eccKeyEncoder ECC通信私钥编码方式，可以是HEX或BASE64
	AuthKey        string // authKey 授权币种私钥
	AuthKeyEncoder string // authKeyEncoder 授权币种私钥编码方式，可以是HEX或BASE64
	Language       string // language 报错语言
}
