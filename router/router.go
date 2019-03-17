package router

import (
	"github.com/gin-gonic/gin"

	// "github.com/lotteryjs/ten-minutes-api/api"
	"github.com/lotteryjs/ten-minutes-api/database"
	"github.com/lotteryjs/ten-minutes-api/error"
	"github.com/lotteryjs/ten-minutes-api/model"
)

// Create creates the gin engine with all routes.
func Create(db *database.TenDatabase, vInfo *model.VersionInfo) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery(), error.Handler())
	g.NoRoute(error.NotFound())

	// userHandler := api.UserAPI{DB: db, PasswordStrength: conf.PassStrength}

	// authAdmin := g.Group("/user")
	// {
	// 	authAdmin.GET("", userHandler.GetUsers)
	// 	authAdmin.POST("", userHandler.CreateUser)
	// }

	// swagger:operation GET /version version getVersion
	//
	// Get version information.
	//
	// ---
	// produces: [application/json]
	// responses:
	//   200:
	//     description: Ok
	//     schema:
	//         $ref: "#/definitions/VersionInfo"
	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	return g
}
