package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusBadGateway, fmt.Errorf("err msg:%v", err))
			}
		}()
		ctx.Next()
	}
}
