package core

import (
	"fmt"
	"time"

	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GlobalConfig.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GlobalConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GlobalLog.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 baileys
	当前版本:v2.5.5
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.GlobalLog.Error(s.ListenAndServe().Error())
}
