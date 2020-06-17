package product

import (
	"github.com/gin-gonic/gin"
)

// product 路由注册
func RegisterRoutes(router *gin.Engine) {
	ProductRouter := router.Group("/product")
	{
		ProductRouter.GET("/list", List)         // 产品列表
		ProductRouter.POST("", Add)              // 新增产品
		ProductRouter.PUT("/:id", Edit)          // 更新产品
		ProductRouter.DELETE("/:id", Delete)     // 删除产品
		ProductRouter.GET("/detail/:id", Detail) // 获取产品详情
	}
}
