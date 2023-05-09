package middlewares

import (
	"bike_manage/config"
	"bike_manage/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"unionid":  user.Unionid,
		"exp":      time.Now().Add(time.Hour * time.Duration(config.Expires)).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// JWTAuth Middleware for JWT authentication
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is invalid"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				now := time.Now().Unix()
				if int64(exp) < now {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
				c.Abort()
				return
			}

			if _, ok := claims["user_id"].(float64); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
				c.Abort()
				return
			}

			if _, ok := claims["username"].(string); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
				c.Abort()
				return
			}

			if _, ok := claims["unionid"].(string); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token error"})
				c.Abort()
				return
			}

			c.Set("user_id", uint(claims["user_id"].(float64)))
			c.Set("username", claims["username"].(string))
			c.Set("unionid", claims["unionid"].(string))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
