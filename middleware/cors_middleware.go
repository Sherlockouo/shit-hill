package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Allow-Credentials", "true")
		}()

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}

	}
}
