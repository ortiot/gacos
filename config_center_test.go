package gacos

import (
	"crypto/md5"
	"encoding/hex"
	"gacos/param"
	"testing"
)

func TestGetConfig(t *testing.T)  {
	g:=SingleGacos("http://127.0.0.1:8848")
	config,err:=g.GetConfig(&param.CfgParam{DataId:"springboot2-nacos-config",Group:"DEFAULT_GROUP"})
	if err!=nil {
		t.Log(err.Error())
		t.Fail()
	}
	t.Log(config)
	h:=md5.New()
	h.Write([]byte(config))

	t.Log(hex.EncodeToString(h.Sum(nil)))
}

func TestListenConfig(t *testing.T)  {
	g:=SingleGacos("http://127.0.0.1:8848")
	//g.ListenConfig(&param.ListenParam{CfgParam:param.CfgParam{DataId:"springboot2-nacos-config",Group:"DEFAULT_GROUP"}}, func(s string, err error) {
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(s)
	//})

	config,err:=g.listenConfig(&param.ListenParam{CfgParam:param.CfgParam{DataId:"springboot2-nacos-config",Group:"DEFAULT_GROUP"}})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	t.Log(config)
	//time.Sleep(3*time.Second)
}