package main

import (
	"challenge-goapi/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	groupRouter := router.Group("/api")
	{
		custRoute := groupRouter.Group("/customer")
		{
			custRoute.GET("/", handler.GetAllCustomer)
			custRoute.POST("/", handler.AddCustomer)
			custRoute.PUT("/:id", handler.UpdateCustomer)
			custRoute.DELETE("/:id", handler.DeleteCustomer)
		}

		serviceRoute := groupRouter.Group("/service")
		{
			serviceRoute.GET("/", handler.GetAllServices)
			serviceRoute.POST("/", handler.AddService)
		}

		trnRoute := groupRouter.Group("/transaction")
		{
			// trnRoute.GET("/", handler.GetTransaction)
			trnRoute.POST("/", handler.AddTransaction)
		}
	}

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
