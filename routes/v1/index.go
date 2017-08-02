package v1

import "github.com/gin-gonic/gin"

//InitRoutes init routes
func InitRoutes(g *gin.RouterGroup) {
	// g.Use(appid.AppIDMiddleWare())
	SetPerformanceRoutes(g)
	// g.Use(token.TokenAuthMiddleWare()) //secure the API From this line to bottom with JSON Auth
	// g.Use(appid.ValidateAppIDMiddleWare())
	// SetTaskRoutes(g)
	// SetUserRoutes(g)
}
