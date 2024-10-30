package handler

import (
	"challenge-goapi/entity"
	"database/sql"
	"net/http"
	"strconv"

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

func GetTransactions(c *gin.Context) {
	params := c.Query("id")

	query := "SELECT id, unit, amount, date_in, date_out, id_customer, id_service FROM trn_laundry"

	var rows *sql.Rows
	var err error

	if params != "" {
		query += " WHERE id=$1"
		rows, err = db.Query(query, params)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "res": err.Error()})
		return
	}

	defer rows.Close()

	var matchedTransaction []entity.Laundry
	for rows.Next() {
		var transaction entity.Laundry
		err := rows.Scan(&transaction.Id, &transaction.Unit, &transaction.Amount, &transaction.DateIn, &transaction.DateOut, &transaction.IdCustomer, &transaction.IdService)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "res": err.Error()})
			return
		}
		matchedTransaction = append(matchedTransaction, transaction)
	}

	if len(matchedTransaction) > 0 {
		c.JSON(http.StatusOK, matchedTransaction)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
	}
}

func DeleteTransactions(c *gin.Context) {
	var transaction entity.Laundry

	queryParams := c.Param("id")
	param, err := strconv.Atoi(queryParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := "DELETE FROM TRN_LAUNDRY WHERE id = $1 RETURNING id"

	err = db.QueryRow(query, param).Scan(&transaction.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "delete success", "id": transaction.Id})
}
