package tracer

import (
	"log"
)

var tsrv = otel.Tracer("gin-server")

func Tracer() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("==tracer==")
		context.Set("key-1", "data-1")
		context.Next()
	}
}
