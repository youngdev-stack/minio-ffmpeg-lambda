package minio

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/youngdev-stack/minio-ffmpeg-lambda/api/v1"
)

type FfmpegRouter struct{}

func (e *FfmpegRouter) InitRuntimeRouter(Router *gin.RouterGroup) {
	minioRouter := Router.Group("/")
	MinioService := v1.ApiGroupApp.MinioGroup.FfmpegAPI
	{
		minioRouter.POST("", MinioService.CovertVideo)
	}
}
