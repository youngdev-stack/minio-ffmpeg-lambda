package service

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/service/minio"
)

type ServiceGroup struct {
	MinioServiceGroup minio.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
