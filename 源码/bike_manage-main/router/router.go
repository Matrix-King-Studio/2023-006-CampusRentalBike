package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	InitAuth(router)
	InitCustomerRouter(router)
	InitRBAC(router)
	InitBikeRouter(router)
	InitFenceRouter(router)
	InitIncomeRouter(router)
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	InitUser(router)
	InitPriceRouter(router)
}
