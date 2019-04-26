package gacos

import (
	"errors"
	"fmt"
	"gacos/param"
	"io"
	"net/http"
	"strings"
	"time"
)

func (g *gacos) GetConfig(param *param.CfgParam) (config string, err error) {
	if param == nil || param.Group == "" || param.DataId == "" {
		return "", errors.New("dataId 和 group 不能为null")
	}
	var tenantParam = ""
	if param.Tenant != "" {
		tenantParam = "&tenant=" + param.Tenant
	}
	if resp, err := http.Get(g.endPoint + "/nacos/v1/cs/configs?dataId=" + param.DataId + "&group=" + param.Group + tenantParam); err != nil {
		return "", err
	} else {
		if resp.StatusCode != 200 {
			return "", errors.New("errcode is " + resp.Status)
		}
		defer resp.Body.Close()
		b := make([]byte, 2<<10)
		n, err := resp.Body.Read(b)
		if err != nil && err != io.EOF {
			return "", err
		}
		return string(b[:n]), nil
	}
}

func (g *gacos)ListenConfig(param *param.ListenParam,f func(config string,err error)) {
	go func() {
		tick:=time.Tick(30*time.Second)
		for range tick{
			f(g.listenConfig(param))
		}
	}()
}

func (g *gacos) listenConfig(param *param.ListenParam) (config string, err error) {
	if param == nil || param.Group == "" || param.DataId == "" {
		return "", errors.New("dataId 和 group 不能为null")
	}
	client := http.Client{}
	//bs := "Listening-Configs=" + param.DataId + "^"
	md5:=make([]byte,0)
	tentant:=make([]byte,0)
	b2:=[]byte{25,30,32}
	b1:=[]byte{25,30,31}
	if param.ContentMD5!="" {
		md5=append(md5,[]byte(param.ContentMD5)...)
	}
	if param.Tenant!="" {
		tentant=append(tentant,[]byte(param.Tenant)...)
	}
	bs := bytesCombine([]byte("Listening-Configs=" + param.DataId),b2,[]byte(param.Group),b2,md5,b2,tentant,b1)


	bstr := param.DataId+string(rune(2))+param.Group+string(rune(2))+param.Tenant+string(rune(1))
	bstr="Listening-Configs="+bstr

	//b1:=[]byte(string(rune(2)))
	//b2:=[]byte(string(rune(1)))
	fmt.Println(string(bs))
	fmt.Println(b1)
	fmt.Println(b2)

	rm := make(map[string]string)
	rm["Listening-Configs"]=bstr

	req, err := http.NewRequest(http.MethodPost, g.endPoint+"/nacos/v1/cs/configs/listener", strings.NewReader(bstr))
	if err != nil {
		return "", err
	}
	req.Header.Add("Long-Pulling-Timeout","30000")
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")

	if resp, err := client.Do(req); err != nil {
		return "", err
	} else {
		if resp.StatusCode != 200 {
			return "", errors.New("errcode is " + resp.Status)
		}
		defer resp.Body.Close()
		b := make([]byte, 2<<10)
		n, err := resp.Body.Read(b)
		if err != nil && err != io.EOF {
			return "", err
		}
		return string(b[:n]), nil
	}
}
