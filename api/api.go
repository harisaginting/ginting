package api

import (
	"github.com/harisaginting/ginting/pkg/wire"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestV1(r *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	apiUser := wire.ApiUser(db)


	// group rest
	rest := r.Group("rest")
	{
		// group v1
		v1 := rest.Group("v1")
		{
			// user
			apiUserGroup := v1.Group("user")
			{
				apiUserGroup.GET("/", apiUser.List)
			}
		}
	}

	return
}