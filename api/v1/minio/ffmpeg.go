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

	c.Header("x-amz-request-route", eventReq.GetObjectContext.OutputRoute)
	c.Header("x-amz-request-token", eventReq.GetObjectContext.OutputToken)

	go ffmpegService.ConvertVideo(eventReq)

	response.OkWithDetailed("", "发起转码成功", c)
}
