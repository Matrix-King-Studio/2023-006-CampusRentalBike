package controllers

import (
	"bike_manage/config"
	"bike_manage/middlewares"
	"bike_manage/models"
	AuthValidator "bike_manage/validate"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// WeChatLoginResponseSuccess 申请token
type WeChatLoginResponseSuccess struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

// WeChatLoginResponseFail 接受微信用户错误信息
type WeChatLoginResponseFail struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// WeChatUserInfo 接受微信用户信息
type WeChatUserInfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid"`
}

// Register 注册
func Register(c *gin.Context) {
	var req AuthValidator.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	user.Username = req.Username

	if err := user.SetPassword(req.Password); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to set password"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to create user", "info": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login 登录
func Login(c *gin.Context) {
	var req AuthValidator.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or password incorrect"})
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or password incorrect"})
		return
	}

	token, err := middlewares.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "token": token})
}

// WeChatLogin 微信登陆
func WeChatLogin(c *gin.Context) {
	var queryParams AuthValidator.WeChatLogin
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "请求参数有误！"})
		return
	}
	// 创建 HTTP 客户端
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/sns/oauth2/access_token", nil)

	query := req.URL.Query()
	query.Add("appid", config.AppID)
	query.Add("secret", config.AppSecret)
	query.Add("code", queryParams.Code)
	query.Add("grant_type", "authorization_code")
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "detail": "request error"})
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var resultF WeChatLoginResponseFail
		if err := json.Unmarshal(body, &resultF); err == nil && resultF.ErrCode != 0 {
			// 错误响应
			c.JSON(http.StatusForbidden, gin.H{"error_code": resultF.ErrCode, "error_message": resultF.ErrMsg})
		} else {
			// 成功响应
			var resultS WeChatLoginResponseSuccess

			if err := json.Unmarshal(body, &resultS); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误！"})
			}

			// 请求用户信息
			req, err = http.NewRequest("GET", "https://api.weixin.qq.com/sns/userinfo", nil)
			if err != nil {
				// 处理错误
				c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误！"})
				return
			}

			query = req.URL.Query()
			query.Add("access_token", resultS.AccessToken)
			query.Add("openid", resultS.OpenId)
			req.URL.RawQuery = query.Encode()

			resp, err = client.Do(req)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "detail": "request error"})
				return
			}
			defer resp.Body.Close()

			body, err = io.ReadAll(resp.Body)
			if err != nil {
				// 处理错误
				c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误！"})
				return
			}

			if resp.StatusCode == http.StatusOK {
				var userInfo WeChatUserInfo
				if err := json.Unmarshal(body, &userInfo); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误！"})
					return
				}
				// 获得用户信息后
				var user models.User
				if err := config.DB.Where("unionid = (?)", userInfo.UnionID).First(&user).Error; err != nil {
					user.Username = userInfo.Nickname
					user.Unionid = userInfo.UnionID
					user.Avatar = userInfo.HeadImgURL
					config.DB.Create(&user)
				}

				user.Avatar = userInfo.HeadImgURL

				if err := config.DB.Save(&user).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "detail": "Update False!"})
					return
				}

				token, err := middlewares.GenerateToken(user)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"status": "success", "token": token, "info": gin.H{"name": user.Username, "avatar": user.Avatar}})
			}
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误！"})
	}

}
