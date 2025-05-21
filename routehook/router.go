package routehook

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRouterModule interface {
	RegisterRoutes(rg *gin.RouterGroup, tx *gorm.DB)
}

var modules []IRouterModule

func RegisterModule(module IRouterModule) {
	modules = append(modules, module)
}

func SetModuleRoutes(rg *gin.RouterGroup, tx *gorm.DB) {
	for _, module := range modules {
		module.RegisterRoutes(rg, tx)
	}
}
