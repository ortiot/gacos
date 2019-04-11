package configs

import (
	"flag"
	"github.com/MWangxj/gocfg"
	"github.com/MWangxj/logger"
	"os"
	"reflect"
)


var (
	GCfg config
	cfg  profile
)

type profile struct {
	Profile string `yaml:"profile"`
}

type config struct {
	//这里写你的配置结构体
}

// init
func init() {
	// 初始化 接受命令行参数激活配置文件和端口使用
	env := flag.String("env", "", "USAGE [-env = dev or prod] to choose profile")
	port := flag.Int("port", 0, "USAGE [-port = 8081 ] to user port od use default 8080")
	flag.Parse()
	// 读取激活的配置文件 默认在app.yml文件中，命令行参数优先级大于此文件中的配置文件激活选择权重
	if err := gocfg.LoadYml("./configs/app.yaml", &cfg); err != nil || cfg.Profile == "" {
		os.Exit(101)
		return
	}

	if *env == "" {
		env = &cfg.Profile
	}
	// 加载配置文件，并反射到结构体
	if err := gocfg.LoadYml("./configs/app-"+*env+".yaml", &GCfg); err != nil {
		os.Exit(102)
		return
	}
	GCfg.Runmode = *env

	if *port != 8080 && *port != 0 {
		GCfg.App.Port = *port
	}
	GCfg.debubPrint()
}

// debubPrint
func (c *config) debubPrint() {
	t := reflect.ValueOf(c).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		logger.Info(t.Type().Field(i).Name, f.Interface())
	}
}