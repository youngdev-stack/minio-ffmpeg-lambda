package main

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/core"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GlobalViper = core.Viper() // 初始化Viper
	global.GlobalLog = core.Zap()     // 初始化zap日志库
	zap.ReplaceGlobals(global.GlobalLog)
	//global.GlobalDB = initialize.GormMysql() // gorm连接数据库
	// initialize.Timer()

	//if global.GlobalDB != nil {
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.GlobalDB.DB()
	//	defer db.Close()
	//}

	//if global.BaileysConfig.System.RunMode == "docker" {
	//	global.BaileysDockerClient = initialize.DockerClient()
	//} else if global.BaileysConfig.System.RunMode == "k8s" {
	//	global.BaileysK8sConfig = initialize.NewK8sConfig()
	//	global.BaileysK8sClient = initialize.NewDynamicClient(global.BaileysK8sConfig)
	//}

	core.RunWindowsServer()
}
