package v1

import (
	book "github.com/Tez-Private/jetbook/api/book/v1"
	rent "github.com/Tez-Private/jetbook/api/rent/v1"
	user "github.com/Tez-Private/jetbook/api/user/v1"

	"github.com/gin-gonic/gin"
)

// RouteUser is routing to user
func RouteUser(router *gin.RouterGroup) {
	userRoute := router.Group("/users")
	userRoute.POST("", user.Create)      // POST api/v1/users
	userRoute.GET("", user.GetAll)       // GET api/v1/users
	userRoute.GET(":id", user.Get)       // GET api/v1/users/[id]
	userRoute.PATCH(":id", user.Update)  // PATCH api/v1/users/[id]
	userRoute.DELETE(":id", user.Delete) // DELETE api/v1/users/[id]
}

// RouteBook is routing to book
func RouteBook(router *gin.RouterGroup) {
	bookRoute := router.Group("/books")
	bookRoute.POST("", book.Create)        // POST api/v1/books
	bookRoute.GET("", book.GetAll)         // GET api/v1/books
	bookRoute.GET(":isbn", book.Get)       // GET api/v1/books/[isbn]
	bookRoute.DELETE(":isbn", book.Delete) // DELETE api/v1/books/[isbn]
}

// RouteRent is routing to rent
func RouteRent(router *gin.RouterGroup) {
	rentRoute := router.Group("/rents")
	rentRoute.POST("", rent.Create)            // POST api/v1/rents
	rentRoute.GET("", rent.GetAll)             // GET api/v1/rents
	rentRoute.GET(":id", rent.Get)             // GET api/v1/rents/[id]
	rentRoute.PATCH(":id/return", rent.Return) // PATCH api/v1/rents/[id]/return
	//rentRoute.PATCH(":id/changereturndate", rent.ChangeReturnDate) // PATCH api/v1/[id]/changereturndate
	//rentRoute.DELETE(":id", rent.Delete)       // DELETE api/v1/rents/[id]
}
