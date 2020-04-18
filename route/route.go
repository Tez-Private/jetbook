package route

import (
	v1route "github.com/Tez-Private/jetbook/route/v1"

	"github.com/gin-gonic/gin"
)

// V1 is application v1 route.
func V1(api *gin.Engine) {
	router := api.Group("api/v1")
	v1route.RouteUser(router) // api/v1/users
	v1route.RouteBook(router) // api/v1/books
}
