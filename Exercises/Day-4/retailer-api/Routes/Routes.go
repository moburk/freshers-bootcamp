package Routes

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Config"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter configuring routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	controllerStruct := &Controllers.ControllerStruct{
		DB: Config.GetDB(),
	}
	grp1 := r.Group("/retail-api/")
	{
		grp1.POST("product", controllerStruct.AddProduct)
		grp1.PATCH("product/:id", controllerStruct.UpdateProductByID)
		grp1.GET("product/:id", controllerStruct.GetProductByID)
		grp1.GET("products", controllerStruct.GetProducts)
		grp1.POST("order", controllerStruct.PlaceOrder)
		grp1.GET("order/:id", controllerStruct.GetOrderByID)
		grp1.GET("customer/:id", controllerStruct.GetCustomerOrders)
	}

	authorized := r.Group("/retail-api/owner", gin.BasicAuth(gin.Accounts{
		"admin": "1234",
	}))

	authorized.GET("", controllerStruct.GetAllOrders)
	return r
}
