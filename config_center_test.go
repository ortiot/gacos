package gacos

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gacos/param"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	g := SingleGacos("http://192.168.48.121:8848")
	config, err := g.GetConfig(&param.CfgParam{DataId: "springboot2-nacos-config", Group: "DEFAULT_GROUP"})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	t.Log(config)
	h := md5.New()
	h.Write([]byte(config))

	t.Log(hex.EncodeToString(h.Sum(nil)))
}

func TestListenConfig(t *testing.T) {
	g := SingleGacos("http://127.0.0.1:8848")
	p:=&param.CfgParam{DataId:"springboot2-nacos-config",Group:"DEFAULT_GROUP"}
	g.ListenConfig(p, func(isupdate bool, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(isupdate)
		if isupdate {
			cfg,err:=g.GetConfig(p)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(cfg)
		}
	})
	time.Sleep(30*time.Second)
}
