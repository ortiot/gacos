package gacos

type CfgParam struct {
	DataId string
	Group  string
	Tenant string
}

type cfgResp struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
