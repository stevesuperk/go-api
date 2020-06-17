package route

import (
	"github.com/gin-gonic/gin"
	"go-api/app/product"
	"go-api/app/test"
	"go-api/route/middleware/exception"
	"go-api/route/middleware/jaeger"
	"go-api/route/middleware/limiter"
	"go-api/route/middleware/logger"
	"go-api/util/response"
)

func InitRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), limiter.SetUp(5), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	//公共模块路由注册
	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	// 测试链路追踪
	//engine.GET("/jaeger_test", jaeger_conn.JaegerTest)

	//@todo 记录请求超时的路由

	//模块路由注册
	product.RegisterRoutes(engine) //产品模块路由注册
	test.RegisterRoutes(engine)    //测试模块路由注册
}
