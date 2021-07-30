package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Config"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//PlaceOrder Places order for a product if it exists, and updates product quantity
func (cs *ControllerStruct) PlaceOrder(c *gin.Context) {
	var order Models.Order
	if err := c.BindJSON(&order); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	productId := order.ProductID
	customerId := order.CustomerID
	orderQuantity := order.Quantity
	order.Status = "order placed"
	lastOrderTime, ok := Config.CustomerOrders[customerId]
	if ok && time.Now().Sub(lastOrderTime).Milliseconds() < 300000 {
		c.JSON(http.StatusTooEarly, "Customer is in cool-down period")
		return
	}
	product, exists, err := Repo.GetProductByID(productId, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNoContent, "Product does not exist")
		return
	}
	if product.Quantity - orderQuantity < 0 {
		order.Status = "failed"
	}
	if order.ID == "" {
		order.ID = "ORD" + Config.GetOrderCounter()
	}
	err = Repo.PlaceOrder(&order, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	Config.CustomerOrders[customerId] = order.CreatedAt
	if order.Status == "failed" {
		c.JSON(http.StatusOK, order)
		return
	}
	if err = Repo.UpdateProductByID(&product, cs.DB, product.Price, product.Quantity-orderQuantity); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, order)
}

func (cs *ControllerStruct) GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, exists, err := Repo.GetOrderByID(id, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusOK, "Order does not exist")
		return
	}
	c.JSON(http.StatusOK, product)
}

//GetAllOrders get all orders by calling another function for a DB call
func (cs *ControllerStruct) GetAllOrders(c *gin.Context) {
	orders, err := Repo.GetAllOrders(cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders" : orders})
}

func (cs *ControllerStruct) GetCustomerOrders(c *gin.Context) {
	id := c.Params.ByName("id")
	exists, err := Repo.CheckCustomerExists(id, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var orders []Models.Order
	orders, exists, err = Repo.GetCustomerOrders(id, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, "No orders for this customer")
		return
	}
	c.JSON(http.StatusOK, orders)
}
