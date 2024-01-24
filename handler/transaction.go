package handler

import (
	"challenge-goapi/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	var newTransaction entity.Laundry

	// date_in, err := time.Parse(time.DateOnly, newTransaction.DateIn)
	// date_out, err := time.Parse(time.DateOnly, newTransaction.DateIn)

	err := c.ShouldBind(&newTransaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO trn_laundry (unit, amount, date_in, date_out, id_customer, id_service) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	var laundryId int
	err = db.QueryRow(query, newTransaction.Unit, newTransaction.Amount, newTransaction.DateIn, newTransaction.DateOut, newTransaction.IdCustomer, newTransaction.IdService).Scan(&laundryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service", "resp": err.Error()})
		return
	}

	newTransaction.Id = laundryId
	c.JSON(http.StatusCreated, gin.H{"data": newTransaction, "status": "insert success"})
}
