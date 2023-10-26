package minio

import (
	"github.com/gin-gonic/gin"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/models/common/response"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/models/minio/request"
)

type FfmpegAPI struct{}

// CovertVideo
// @Tags      RuntimeApi
// @Summary   获取当前运行模式
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=string,msg=string}  "创建基础api"
// @Router    /baileys/minio/getRuntime [get]
func (e FfmpegAPI) CovertVideo(c *gin.Context) {
	var eventReq request.EventReq

	if err := c.ShouldBind(&eventReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ffmpegService.ConvertVideo(eventReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData("转码成功", c)
}
