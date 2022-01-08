package main

import (
	"fmt"

	"github.com/piLinux/postmark/controller"
	"github.com/piLinux/postmark/migration"

	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/lib/middleware"

	"github.com/gin-gonic/gin"
)

var configure = config.Config()

func main() {
	if err := database.InitDB().Error; err != nil {
		fmt.Println(err)
	}

	// Migrate database
	// DBMigrate(true): drop previous tables
	// DBMigrate(false): create tables with missing columns and missing indexes
	migration.DBMigrate(false)

	router := SetupRouter()
	router.Run(":" + configure.Server.ServerPort)
}

// SetupRouter ...
func SetupRouter() *gin.Engine {
	// debug or release mode
	if configure.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin.Default() = gin.New() + gin.Logger() + gin.Recovery()
	router := gin.Default()

	// Which proxy to trust
	if configure.Security.TrustedIP == "nil" {
		router.SetTrustedProxies(nil)
	} else {
		if configure.Security.TrustedIP != "" {
			router.SetTrustedProxies([]string{configure.Security.TrustedIP})
		}
	}

	router.Use(middleware.CORS())
	router.Use(middleware.SentryCapture(configure.Logger.SentryDsn))
	router.Use(middleware.Firewall(
		configure.Security.Firewall.ListType,
		configure.Security.Firewall.IP,
	))

	// API:v1
	v1 := router.Group("webhooks/v1")
	{
		// basic auth
		user := configure.Security.BasicAuth.Username
		pass := configure.Security.BasicAuth.Password
		v1.Use(gin.BasicAuth(gin.Accounts{user: pass}))

		// postmark webhooks for outbound
		v1.POST("outbound-events", controller.Outbound)
	}

	return router
}
