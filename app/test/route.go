package test

import (
	"github.com/gin-gonic/gin"
)

// test 测试加密性能
func RegisterRoutes(router *gin.Engine) {
	TestRouter := router.Group("/test")
	{
		TestRouter.GET("/md5", Md5Test) // 测试 MD5 组合 的性能
		TestRouter.GET("/aes", AesTest) // 测试 AES 的性能
		TestRouter.GET("/rsa", RsaTest) // 测试 RSA 的性能
	}
}
