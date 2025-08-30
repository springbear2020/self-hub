package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	holder(publicGroup, privateGroup)

	{
		taskRouter := router.RouterGroupApp.Task
		taskRouter.InitDailyTaskRouter(privateGroup, publicGroup)
		taskRouter.InitDailyTaskCompletionRouter(privateGroup, publicGroup)
	}
	{
		miserRouter := router.RouterGroupApp.Miser
		miserRouter.InitMiserCategoryRouter(privateGroup, publicGroup)
		miserRouter.InitMiserTransactionRouter(privateGroup, publicGroup)
		miserRouter.InitMiserTransactionItemRouter(privateGroup, publicGroup)
		miserRouter.InitMiserStatRouter(privateGroup, publicGroup)
	}
}
