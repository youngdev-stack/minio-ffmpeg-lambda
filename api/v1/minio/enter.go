package minio

import "github.com/youngdev-stack/minio-ffmpeg-lambda/service"

type ApiGroup struct {
	FfmpegAPI
}

var (
	ffmpegService = service.ServiceGroupApp.MinioServiceGroup.FfmpegService
)
