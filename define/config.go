package define

import (
	"flag"

	"github.com/BurntSushi/toml"
	logger "github.com/shengkehua/xlog4go"
)

var (
	FrontEndServerCfg FrontEndServerConfig
)

func InitConfig() error {
	configFile := flag.String("config", "conf/frontEnd.toml", "config file name")
	flag.Parse()

	_, err := toml.DecodeFile(*configFile, &FrontEndServerCfg)
	if nil != err {
		return err
	}
	return nil
}

func InitLogger() error {
	err := logger.SetupLogWithConf(FrontEndServerCfg.LogCfg.FileName)
	return err
}
