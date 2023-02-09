package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApi(router *gin.RouterGroup, server *Server) {

	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })

}
