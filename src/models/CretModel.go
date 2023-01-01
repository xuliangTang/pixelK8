package models

type CertModel struct {
	CN        string `json:"域名 (CN)"`          //域名
	Algorithm string `json:"算法 (Algorithm)"`   //算法
	Issuer    string `json:"签发者 (Issuer)"`     //签发者
	BeginTime string `json:"生效时间 (BeginTime)"` //生效时间
	EndTime   string `json:"到期时间 (EndTime)"`   //到期时间
}
