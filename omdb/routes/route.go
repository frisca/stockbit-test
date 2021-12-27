package routes

import (
	"stockbit-test/omdb/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/omdb")
	{
		v1.GET("/movie", controllers.List)
		v1.GET("/movie/detail/:imdbID", controllers.Get)
	}

	return r
}
