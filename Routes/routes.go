package Routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kabos0809/redis_jwt_gin/Models"
	"github.com/kabos0809/redis_jwt_gin/Controller"
	"github.com/kabos0809/redis_jwt_gin/MiddleWare"
)

func SetRoutes(ctrl Models.Model) *gin.Engine {
	e := gin.Default()

	v1 := e.Group("/v1")
	{
		v1.POST("/signup", Controller.Controller{Model: ctrl}.SignUp)
		v1.POST("/login", Controller.Controller{Model: ctrl}.Login)
		v1.GET("/logout", MiddleWare.CheckJWT(), Controller.Controller{Model: ctrl}.Logout)
		v1.GET("/refresh", MiddleWare.CheckRefresh(), Controller.Controller{Model: ctrl}.GetReToken)
		v1.GET("/home", MiddleWare.CheckJWT(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hellooooooooooo!!!!",
			})
		})
	}

	return e
}