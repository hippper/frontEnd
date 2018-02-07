package main

import (
	"log"
	"runtime"
	"sync"

	"github.com/hippper/frontEnd/define"
	"github.com/hippper/frontEnd/server"
	"github.com/hippper/frontEnd/utils"

	logger "github.com/shengkehua/xlog4go"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//recover work
	defer func() {
		err := recover()
		if nil != err {
			stackInfo := utils.GetStackInfo()
			logger.Fatal("panic stackinfo: %s", stackInfo)
		}
	}()

	//init config
	if err := define.InitConfig(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	//init log
	if err := define.InitLogger(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	defer logger.Close()

	//init db

	//start server
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		webServer := server.NewServer()
		err := webServer.StartServer()
		if err != nil {
			logger.Error("Start server failed, err: %v", err)
		}
	}()

	waitGroup.Wait()

	return
}
