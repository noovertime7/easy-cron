package main

import (
	"es-cron/config"
	"es-cron/logger"
	"es-cron/pkg/core"
	"es-cron/pkg/utils"
	"os"
)

func main() {
	//注册logger
	utils.Must(logger.InitLogger())
	utils.Must(config.InitConfig())

	for _, obj := range config.SysConfig.Cron {
		handler := core.NewCronHandler(obj.Name, obj.Cron, obj.Shell)
		utils.Must(handler.Start())
	}

	<-utils.SetupSignalHandler()
	os.Exit(0)
}
