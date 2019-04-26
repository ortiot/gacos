package param

type CfgParam struct {
	DataId string
	Group string
	Tenant string
}

type ListenParam struct {
	ContentMD5 string
	CfgParam
}