package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Config"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AddProduct Passes arguments to AddProduct to make a database call
func (cs *ControllerStruct) AddProduct(c *gin.Context){
	var product Models.Product
	err := c.BindJSON(&product)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if product.ID == "" {
		product.ID = "PROD" + Config.GetProductCounter()
	}
	err = Repo.AddProduct(&product, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"id":product.ID,
		"product_name":product.ProductName,
		"price":product.Price,
		"quantity":product.Quantity,
		"message":"product successfully added",
	})
}

//UpdateProductByID Check if product exists. if yes, update
func (cs *ControllerStruct) UpdateProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, exists, err := Repo.GetProductByID(id, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, "Product does not exist")
		return
	}
	requestJson := struct {
		Price int `json:"price"`
		Quantity int `json:"quantity"`
	}{}
	err = c.BindJSON(&requestJson) //not working
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err = Repo.UpdateProductByID(&product, cs.DB, requestJson.Price, requestJson.Quantity); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

func (cs *ControllerStruct) GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, exists, err := Repo.GetProductByID(id, cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, "Product does not exist")
		return
	}
	c.JSON(http.StatusOK, product)
}


func (cs *ControllerStruct) GetProducts(c *gin.Context) {
	products, err := Repo.GetProducts(cs.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"products" : products})
}