package controllers

import (
	"bike_manage/config"
	"bike_manage/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CustomerInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User
	if err := config.DB.Preload("Roles.Permissions").First(&user, userID.(uint)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"status": "success", "userinfo": gin.H{
	//	"avatar": user.Avatar,
	//	"name":   user.Username,
	//	"role":   user.Roles,
	//}})
	roles := make([]gin.H, len(user.Roles))
	for i, role := range user.Roles {
		permissions := make([]gin.H, len(role.Permissions))
		for j, permission := range role.Permissions {
			permissions[j] = gin.H{
				//"id":   permission.ID,
				"name": permission.Name,
			}
		}
		roles[i] = gin.H{
			//"id":          role.ID,
			"name":        role.Name,
			"permissions": permissions,
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "userinfo": gin.H{
		"avatar": user.Avatar,
		"name":   user.Username,
		"role":   roles,
	}})
}
