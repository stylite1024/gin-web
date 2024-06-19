package router

import (
	"fmt"
	"gin-web/internal/handler"
	"gin-web/web"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 静态文件路由
	staticFileRouter(r)
	// api路由
	g := r.Group("")
	apiRouter(g)
	return r
}

func staticFileRouter(r *gin.Engine) {
	r.Use(static.Serve("/", static.EmbedFolder(web.StaticFS, "static")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exists, redirect on /\n", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
}

func apiRouter(g *gin.RouterGroup) {
	v1 := g.Group("/api/v1")
	v1.GET("/test", handler.TestApi)
}
