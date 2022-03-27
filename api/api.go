package api

import (
	"github.com/harisaginting/ginting/pkg/wire"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/harisaginting/ginting/pkg/utils/helper"

	// middleware
	kc "github.com/harisaginting/ginting/pkg/keycloak_client"		
)

func RestV1(r *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	apiUser := wire.ApiUser(db)

	cfgkc 	 := helper.ForceInt(helper.MustGetEnv("KEUCLOAK"))
	keycloak := kc.Start(cfgkc)

	// group rest
	rest := r.Group("rest")
	{
		// group v1
		v1 := rest.Group("v1")
		{
			// user
			apiUserGroup := v1.Group("user")
			{
				apiUserGroup.GET("/", keycloak.Validate([]string{"customerservice:dashboard:email-logs:read"}), apiUser.List)
			}
		}
	}

	return
}