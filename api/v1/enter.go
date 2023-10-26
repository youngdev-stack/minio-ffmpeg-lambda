package v1

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/api/v1/minio"
)

type ApiGroup struct {
	MinioGroup minio.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
