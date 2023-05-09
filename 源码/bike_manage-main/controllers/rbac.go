package controllers

import (
	"bike_manage/config"
	models "bike_manage/models"
	RBACValidator "bike_manage/validate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
	"strconv"
)

// CreateRole 创建角色并分配权限
func CreateRole(c *gin.Context) {
	var req RBACValidator.RoleCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{Name: req.Name}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(req.PermissionIDs) > 0 {
		var permissions []models.Permission
		if err := config.DB.Where("id IN (?)", req.PermissionIDs).Find(&permissions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := config.DB.Model(&role).Association("Permissions").Append(permissions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"role": role})
}
func AssignPermissionToRole(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	permissionID, err := strconv.Atoi(c.Param("permission_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	req := RBACValidator.AssignPermission{RoleID: uint(roleID), PermissionID: uint(permissionID)}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := config.DB.Preload("Permissions").First(&role, req.RoleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, permission := range role.Permissions {
		if permission.ID == req.PermissionID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role already has this permission"})
			return
		}
	}

	var permission models.Permission

	if err := config.DB.First(&role, req.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found"})
		return
	}

	if err := config.DB.First(&permission, req.PermissionID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission not found"})
		return
	}

	config.DB.Model(&role).Association("Permissions").Append(&permission)

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Permission assigned to role"})
}

// AssignRoleToUser 为用户分配角色
func AssignRoleToUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	req := RBACValidator.AssignRole{UserID: uint(userID), RoleID: uint(roleID)}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := config.DB.Preload("Roles").First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, role := range user.Roles {
		if role.ID == req.RoleID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already has this role"})
			return
		}
	}
	var role models.Role
	if err := config.DB.First(&role, req.RoleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Roles = append(user.Roles, &role)

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// ListAllPermissions 查询所有权限
func ListAllPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := config.DB.Find(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

// ListAllRolesWithPermissions 查询所有角色及其权限
func ListAllRolesWithPermissions(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	var role models.Role
	var searchRole models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if role.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}
	if err := config.DB.Where("id = (?)", role.ID).First(&searchRole).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role Not Found"})
		return
	}
	if searchRole.Name == "SuperAdmin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SuperAdmin Can't Rewrite!"})
		return
	}

	if err := config.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	id := c.Param("id")

	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if role.Name == "SuperAdmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "SuperAdmin role cannot be deleted"})
		return
	}

	if err := config.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

func UserList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	users, totalCount := getUsers(config.DB, page, pageSize)
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"users":       users,
		"total_pages": totalPages,
	})
}

func getUsers(db *gorm.DB, page int, pageSize int) ([]models.User, int64) {
	var users []models.User
	var totalCount int64

	db.Preload("Roles").Scopes(UserPaginate(page, pageSize)).Find(&users)
	db.Model(&models.User{}).Count(&totalCount)

	return users, totalCount
}
func UserPaginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
