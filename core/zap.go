package core

import (
	"fmt"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/core/internal"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GlobalConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GlobalConfig.Zap.Director)
		_ = os.Mkdir(global.GlobalConfig.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GlobalConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
