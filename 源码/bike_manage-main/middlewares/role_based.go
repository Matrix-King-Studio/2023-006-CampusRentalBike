package middlewares

import (
	"bike_manage/config"
	"bike_manage/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireSuperAdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.Preload("Roles").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		hasSuperAdminRole := false
		for _, role := range user.Roles {
			if role.Name == "SuperAdmin" {
				hasSuperAdminRole = true
				break
			}
		}

		if !hasSuperAdminRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
