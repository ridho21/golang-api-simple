package handler

import (
	"challenge-goapi/entity"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var db = config.ConnectDB()

func AddService(c *gin.Context) {
	var newService entity.Services

	err := c.ShouldBind(&newService)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO mst_services (service, price) VALUES ($1, $2) RETURNING id"

	var servId int

	err = db.QueryRow(query, newService.Service, newService.Price).Scan(&servId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service", "resp": err.Error()})
		return
	}

	newService.Id = servId
	c.JSON(http.StatusCreated, gin.H{"data": newService, "status": "insert success"})
}

func GetAllServices(c *gin.Context) {
	params := c.Query("service")

	query := "SELECT id, service, price FROM mst_services"

	var rows *sql.Rows
	var err error

	if params != "" {
		query += " WHERE name ILIKE '%' || $1 || '%'"
		rows, err = db.Query(query, params)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "res": err.Error()})
		return
	}

	defer rows.Close()

	var matchedService []entity.Services
	for rows.Next() {
		var service entity.Services
		err := rows.Scan(&service.Id, &service.Service, &service.Price)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "res": err.Error()})
			return
		}
		matchedService = append(matchedService, service)
	}

	if len(matchedService) > 0 {
		c.JSON(http.StatusOK, matchedService)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Services not found"})
	}
}
