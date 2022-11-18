package swagger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func SwaggerDocRoute(route *gin.Engine) {
	//swagger integration
	route.GET("/swagger.json", gin.WrapH(http.FileServer(http.Dir("./swagger/"))))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	route.GET("/docs", gin.WrapH(sh))
}
