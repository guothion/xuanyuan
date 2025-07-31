package api

import (
	"github.com/gin-gonic/gin"
	"github.com/guothion/xuanyuan/pkg/api/controller"
	"github.com/guothion/xuanyuan/pkg/api/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func Server() {
	profile := strings.ToLower(os.Getenv("SERVICE_ACTIVE_PROFILE"))
	// 设置模式
	if profile != "production" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	go func() {
		engine := gin.New()
		engine.Use(gin.Recovery())
		setupAdminAPIs(engine)

		addr := "localhost:8082"
		logrus.Info(profile)
		if profile != "production" {
			addr = "localhost:14082"
		}
		logrus.Printf("Admin server is listening on %s", addr)
		if err := engine.Run(addr); err != nil {
			logrus.Fatalf("launch admin server failed: %v", err)
		}
	}()

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(corsMiddleware())
	//engine.GET("/swagger/*any",)
	setupAPIs(engine)

	addr := "localhost:8083"
	if profile != "production" {
		addr = "localhost:14083"
	}
	logrus.Printf("Admin server is listening on %s", addr)
	if err := engine.Run(addr); err != nil {
		logrus.Fatalf("launch admin server failed: %v", err)
	}
}

func setupAPIs(engine *gin.Engine) {
	middleware.Init()

	engine.POST("/login", middleware.LoginHandler)
	controller.Init(engine)
}

func setupAdminAPIs(engine *gin.Engine) {
	engine.GET("/envVars", controller.ShowEnvVars)
	engine.GET("/configs", controller.ShowConfig)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Fl-Service-Profile,X-Gl-Account")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
