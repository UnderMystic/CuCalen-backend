package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	ar "cucalen/essential/autorouter"
	cf "cucalen/essential/config"

	api "cucalen/api/api"
	auth "cucalen/api/auth"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	cfg, _ := cf.Config()

	if cfg.Env.Typ == "Pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/api/*path", ar.AutoRoute(&api.GET{}))
	r.POST("/api/*path", ar.AutoRoute(&api.POST{}))
	r.GET("/auth/*path", ar.AutoRoute(&auth.GET{}))
	r.POST("/auth/*path", ar.AutoRoute(&auth.POST{}))
	host := cfg.Conf.Host + ":" + cfg.Conf.Port
	fmt.Println(cfg.Conf.Host)
	fmt.Println(cfg.Conf.Port)
	r.Run(host)
}
