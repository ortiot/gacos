package gacos

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

func (g *gacos) GetConfig(param *CfgParam) (config string, err error) {
	if param == nil || param.Group == "" || param.DataId == "" {
		return "", errors.New("dataId 和 group 不能为null")
	}
	var tenantParam = ""
	if param.Tenant != "" {
		tenantParam = "&tenant=" + param.Tenant+"&show=all"
	}
	if resp, err := http.Get(g.endPoint + "/nacos/v1/cs/configs?dataId=" + param.DataId + "&group=" + param.Group+tenantParam); err != nil {
		return "", err
	} else {
		if resp.StatusCode != 200 {
			return "", errors.New(errStatusTostring(resp.StatusCode))
		}
		defer resp.Body.Close()
		b := make([]byte, 1<<12)
		n, err := resp.Body.Read(b)
		if err != nil && err != io.EOF {
			return "", err
		}
		content:=string(b[:n])
		if param.Tenant!="" {
			res:=&cfgResp{}
			if err:=json.Unmarshal(b[:n],res);err!=nil{
				return "",err
			}
			if res.Content=="" {
				return "",nil
			}
			content=res.Content
		}
		hash := md5.New()
		hash.Write([]byte(content))
		g.cacheMd5 = hex.EncodeToString(hash.Sum(nil))
		return content, nil
	}
}

func (g *gacos) ListenConfig(param *CfgParam, f func(isupdate bool, err error)) {
	go func() {
		tick := time.Tick(30 * time.Second)
		for range tick {
			f(g.listenConfig(param))
		}
	}()
}

func (g *gacos) listenConfig(param *CfgParam) (isupdate bool, err error) {
	if param == nil || param.Group == "" || param.DataId == "" {
		return false, errors.New("dataId 和 group 不能为null")
	}
	client := http.Client{}
	bstr := param.DataId + string(rune(2)) + param.Group + string(rune(2)) + g.cacheMd5 + string(rune(2)) +param.Tenant+ string(rune(1))
	bstr = "Listening-Configs=" + bstr
	rm := make(map[string]string)
	rm["Listening-Configs"] = bstr
	req, err := http.NewRequest(http.MethodPost, g.endPoint+"/nacos/v1/cs/configs/listener", strings.NewReader(bstr))
	if err != nil {
		return false, err
	}
	req.Header.Add("Long-Pulling-Timeout", "30000")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if resp, err := client.Do(req); err != nil {
		return false, err
	} else {
		if resp.StatusCode != 200 {
			return false, errors.New("errcode is " + resp.Status)
		}
		defer resp.Body.Close()
		b := make([]byte, 2<<10)
		n, err := resp.Body.Read(b)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			return false, nil
		}
		return true, nil
	}
}
