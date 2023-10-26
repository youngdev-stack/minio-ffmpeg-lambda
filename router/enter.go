package router

import "github.com/youngdev-stack/minio-ffmpeg-lambda/router/minio"

type RouterGroup struct {
	FfmpegRouter minio.FfmpegRouter
}

var RouterGroupApp = new(RouterGroup)
