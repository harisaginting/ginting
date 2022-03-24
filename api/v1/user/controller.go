package user

import "github.com/gin-gonic/gin"
import "github.com/harisaginting/ginting/pkg/http/response"
import "github.com/harisaginting/ginting/pkg/tracer"

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) List(c *gin.Context) {
	span := tracer.Span("ListUserController")
	defer span.End()

	var resData ResponseList
	ctrl.service.List(&resData)
	
	// return
	response.Json(c,resData)
	return
}