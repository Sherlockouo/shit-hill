package main

import (
	"gin-toy/data"
	"gin-toy/handler"
	"gin-toy/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegistryRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	r.GET("/search", handleSearch)
	r.GET("/ping", handlePing)

	return r
}

func handlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func handleSearch(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	p := ctx.Query("page")
	page, _ := strconv.ParseInt(p, 10, 64)
	// 解析参数
	search := &data.Search{
		Q:          keywords,
		Categories: 111,
		Purities:   110,
		Sorting:    data.Hot,
		Order:      data.Desc,
		Page:       int(page),
	}
	res, err := handler.SearchWallHavenPapers(search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "请求错误")
	}
	ctx.JSON(http.StatusOK, res)
}
