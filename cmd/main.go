package main

import (
	orderapi "Api/api/order_api"
	productapi "Api/api/product_api"
	userapi "Api/api/user_api"
	"Api/clients/orderclient"
	"Api/clients/productclient"
	"Api/clients/userclient"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	userclient := userclient.DialGrpcUser()
	userhandler := &userapi.UserHandler{ClientUser: userclient}

	productclient := productclient.DialGrpcProduct()
	producthandler := &productapi.ProductHandler{ClientProduct: productclient}

	orderclient := orderclient.DialGrpcOrder()
	orderhandler := &orderapi.OrderHandler{ClientOrder: orderclient}
	
	router := gin.Default()

	// user
	router.POST("/create", userhandler.CreateUser)
	router.POST("/login", userhandler.Login)
	router.GET("/user/:id", userhandler.GetByIDUser)
	router.GET("/alluser", userhandler.GetAllUser)
	// product
	router.POST("/product/create", producthandler.CreateProduct)
	router.GET("/product/:id", producthandler.GetByIdProduct)
	router.GET("/allproduct", producthandler.GetAllProducts)
	router.PUT("/update", producthandler.UpdateProduct)
	router.DELETE("delete/:id", producthandler.DeleteProduct)
	// order
	router.POST("/order/create", orderhandler.CreateOrder)
	router.POST("/createorders", orderhandler.CreateOrders)
	router.GET("/allorder", orderhandler.GetAllOrders)


	router.Run(":9888")

	log.Println("Client is Listening on port", ":9888")
	if err := http.ListenAndServe("localhost:9888", router); err != nil {
		log.Fatal("Unable to listen and serve :", err)
	}
}
