package handler

import (
	"challenge-goapi/config"
	"challenge-goapi/entity"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = config.ConnectDB()

func AddCustomer(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add customer"})
		return
	}

	newCustomer.Id = custId
	c.JSON(http.StatusCreated, newCustomer)
}

func GetAllCustomer(c *gin.Context) {
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

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	custId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer id"})
		return
	}

	var cust entity.Customer

	if err := c.ShouldBind(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE MST_CUSTOMER SET name = $2, no_hp = $3, address = $4 WHERE id = $1"

	_, err = db.Query(query, custId, cust.Name, cust.Phone, cust.Address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	cust.Id = custId
	c.JSON(http.StatusOK, gin.H{"status": "update success", "data": cust})
}

func DeleteCustomer(c *gin.Context) {
	pathParams := c.Param("id")

	param, err := strconv.Atoi(pathParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := "DELETE FROM MST_CUSTOMER WHERE id = $1"

	rows, err := db.Query(query, param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var matchedCustomer []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(param)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		matchedCustomer = append(matchedCustomer, customer)
	}

	if len(matchedCustomer) > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "delete success", "id": param})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
	}
}
