package Config

import (
	"strconv"
	"time"
)

var orderCounter int = 1233
var productCounter int = 12345
var CustomerOrders = make(map[string]time.Time)

func GetOrderCounter() string {
	orderCounter++
	return strconv.Itoa(orderCounter)
}

func GetProductCounter() string {
	productCounter++
	return strconv.Itoa(productCounter)
}
