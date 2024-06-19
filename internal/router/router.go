package router

import (
	"fmt"
	"gin-web/internal/handler"
	"gin-web/web"
	"net/http"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(logger.SetLogger(logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
		return l.Output(gin.DefaultWriter).With().Logger()
	})))

	// pprof.Register(r)
	// default is "debug/pprof"
	pprof.Register(r, "dev/pprof")

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
