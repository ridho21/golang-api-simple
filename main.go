package main

import (
	"challenge-goapi/config"
	"challenge-goapi/entity"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = config.ConnectDB()

func main() {
	router := gin.Default()

	groupRouter := router.Group("/api")
	{
		custRoute := groupRouter.Group("/customer")
		{
			custRoute.GET("/", getAllCustomer)
			custRoute.POST("/", addCustomer)
		}
	}

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}

func addCustomer(c *gin.Context) {
	var newCustomer entity.Customer

	err := c.ShouldBind(&newCustomer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO mst_customer (name, no_hp, address) VALUES ($1, $2, $3) RETURNING id"

	var custId int

	err = db.QueryRow(query, newCustomer.Name, newCustomer.Phone, newCustomer.Address).Scan(&custId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	newCustomer.Id = custId
	c.JSON(http.StatusCreated, newCustomer)
}

func getAllCustomer(c *gin.Context) {
	params := c.Query("name")

	query := "SELECT id, name, no_hp, address FROM mst_customer"

	var rows *sql.Rows
	var err error

	if params != "" {
		query += " WHERE name ILIKE '%' || $1 || '%'"
		rows, err = db.Query(query, params)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	defer rows.Close()

	var matchedCust []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "res": err.Error()})
			return
		}
		matchedCust = append(matchedCust, customer)
	}

	if len(matchedCust) > 0 {
		c.JSON(http.StatusOK, matchedCust)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Custommer not found"})
	}
}

func updateCustomer(c *gin.Context) {

}
