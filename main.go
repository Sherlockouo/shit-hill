package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFS("/public", http.Dir("./public"))

	r = RegistryRoute(r)
	// log.Default().Printf("env port:%v", os.Getenv("PORT"))
	r.Run(":8088")
}
