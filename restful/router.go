package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/ontio/mdserver/middleware/cors"
	"github.com/ontio/mdserver/restful/api"
)

func NewRouter() *gin.Engine {
	gin.DisableConsoleColor()
	root := gin.Default()
	root.Use(cors.Cors())
	api.RoutesApi(root)
	return root
}
