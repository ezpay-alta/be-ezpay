package presentation

import (
	"ezpay/features/middlewares"
	"ezpay/features/products"
	"ezpay/features/transactions"
	"ezpay/features/transactions/presentation/request"
	"ezpay/features/transactions/presentation/response"
	"net/http"
	"os"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type TransactionHandler struct {
	TransactionBusiness transactions.Business
	ProductBusiness     products.Business
}

func NewTransactionHandler(transactionBusiness transactions.Business, productBusiness products.Business) *TransactionHandler {
	return &TransactionHandler{TransactionBusiness: transactionBusiness, ProductBusiness: productBusiness}
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

	productId, _ := ph.ProductBusiness.GetProductByName(transactionData.Product)

	transactionId, err := ph.TransactionBusiness.CreateTransaction(transactionData.ToTransactionCore(userId, productId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create transaction",
			"err":     err.Error(),
		})
	}

	transaction, _ := ph.TransactionBusiness.GetTransactionById(transactionId)

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET_KEY")

	data := invoice.CreateParams{
		ExternalID:      strconv.Itoa(transactionId),
		Amount:          float64(transaction.Total),
		PayerEmail:      transaction.User.Email,
		Description:     transaction.Product.Name,
		ShouldSendEmail: &[]bool{true}[0],
		Customer: xendit.InvoiceCustomer{
			GivenNames:   transaction.User.Fullname,
			Email:        transaction.User.Email,
			MobileNumber: transaction.User.Phone,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"whatsapp", "sms", "email"},
			InvoiceReminder: []string{"whatsapp", "sms", "email"},
			InvoicePaid:     []string{"whatsapp", "sms", "email"},
			InvoiceExpired:  []string{"whatsapp", "sms", "email"},
		},
	}

	invoice.Create(&data)

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

	productId, _ := ah.ProductBusiness.GetProductByName(transactionData.Product)

	err = ah.TransactionBusiness.UpdateTransactionById(transactionId, transactionData.ToTransactionCore(userId, productId))
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

func (ah *TransactionHandler) UpdateTransactionByXenditHandler(e echo.Context) error {

	if e.Request().Header.Get("x-callback-token") != os.Getenv("X_CALLBACK_TOKEN") {
		return e.NoContent(http.StatusForbidden)
	}

	transactionData := request.XenditRequest{}
	e.Bind(&transactionData)

	err := ah.TransactionBusiness.UpdateTransactionById(transactionData.ToTransactionCore().ID, transactionData.ToTransactionCore())
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
