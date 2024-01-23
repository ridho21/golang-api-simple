package handler

import (
	"challenge-goapi/entity"
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
	c.JSON(http.StatusCreated, newService)
}
