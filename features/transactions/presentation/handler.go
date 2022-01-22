package presentation

import (
	"ezpay/features/middlewares"
	"ezpay/features/transactions"
	"ezpay/features/transactions/presentation/request"
	"ezpay/features/transactions/presentation/response"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	TransactionBusiness transactions.Business
}

func NewTransactionHandler(transactionBusiness transactions.Business) *TransactionHandler {
	return &TransactionHandler{TransactionBusiness: transactionBusiness}
}

func (ph *TransactionHandler) CreateTransactionHandler(e echo.Context) error {
	userId, _, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create transaction",
			"err":     "token is invalid",
		})
	}

	transactionData := request.TransactionRequest{}

	e.Bind(&transactionData)

	err = ph.TransactionBusiness.CreateTransaction(transactionData.ToTransactionCore(userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create transaction",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "new transaction is created",
	})
}

func (ph *TransactionHandler) GetAllTransactionsHandler(e echo.Context) error {
	data, err := ph.TransactionBusiness.GetAllTransactions()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all transactions",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToTransactionResponseList(data),
	})

}

func (ph *TransactionHandler) GetTransactionByIdHandler(e echo.Context) error {
	transactionId, err := strconv.Atoi(e.Param("transactionId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get transaction",
			"err":     err.Error(),
		})
	}

	data, err := ph.TransactionBusiness.GetTransactionById(transactionId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get transaction",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToTransactionResponse(data),
	})

}

func (ah *TransactionHandler) UpdateTransactionByIdHandler(e echo.Context) error {
	userId, _, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create transaction",
			"err":     "token is invalid",
		})
	}

	transactionId, err := strconv.Atoi(e.Param("transactionId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update transaction",
			"err":     err.Error(),
		})
	}

	transactionData := request.TransactionRequest{}
	e.Bind(&transactionData)

	err = ah.TransactionBusiness.UpdateTransactionById(transactionId, transactionData.ToTransactionCore(userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update transaction",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update transaction",
	})

}

func (uh *TransactionHandler) DeleteTransactionByIdHandler(e echo.Context) error {
	transactionId, err := strconv.Atoi(e.Param("transactionId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete transaction",
			"err":     err.Error(),
		})
	}

	err = uh.TransactionBusiness.DeleteTransactionById(transactionId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete transaction",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete transaction",
	})

}
