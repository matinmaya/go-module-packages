package user_test

import (
	"github.com/gin-gonic/gin"
	"github.com/matinmaya/go-module-packages/routehook"
	"gorm.io/gorm"
)

type UserModule struct{}

func init() {
	routehook.RegisterModule(&UserModule{})
}

func (m *UserModule) RegisterRoutes(rg *gin.RouterGroup, tx *gorm.DB) {
	rg.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from UserModule",
		})
	})
}
