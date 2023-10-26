package initialize

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/docs"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	exampleRouter := router.RouterGroupApp.FfmpegRouter
	//if global.BaileysConfig.System.RunMode == "docker" {
	//
	//} else if global.BaileysConfig.System.RunMode == "k8s" {
	//	k8sRouter := router.RouterGroupApp.K8s
	//}

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GlobalConfig.System.RouterPrefix
	Router.GET(global.GlobalConfig.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GlobalLog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GlobalConfig.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		exampleRouter.InitRuntimeRouter(PublicGroup)

	}

	global.GlobalLog.Info("router register success")
	return Router
}
