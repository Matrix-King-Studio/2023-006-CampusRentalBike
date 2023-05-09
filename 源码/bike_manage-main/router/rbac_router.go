package router

import (
	"bike_manage/controllers"
	"bike_manage/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRBAC(router *gin.Engine) {
	rbac := router.Group("/rbac")
	rbac.Use(middlewares.JWTAuth())
	//rbac.Use(middlewares.RequireSuperAdminRole())
	{
		rbac.POST("/roles", controllers.CreateRole)
		rbac.GET("/roles/:role_id/permissions/:permission_id", controllers.AssignPermissionToRole)
		rbac.GET("/users/:user_id/roles/:role_id", controllers.AssignRoleToUser)
		rbac.GET("/permissions", controllers.ListAllPermissions)
		rbac.GET("/roles", controllers.ListAllRolesWithPermissions)
		rbac.PUT("/role", controllers.UpdateRole)
		rbac.DELETE("/role/:id", controllers.DeleteRole)
		rbac.GET("/user", controllers.UserList)
	}
}
