package domain_transaction_test

import (
	"errors"
	"os"
	domain_transaction "ppob/transaction/domain"
	transactiontMocks "ppob/transaction/domain/mocks"
	service_transaction "ppob/transaction/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

var (
	transactionService domain_transaction.Service
	transactionDomain  domain_transaction.Transaction
	detailDomain       domain_transaction.Detail_Transaction
	paymentDomain      domain_transaction.Payment
	callbackDomain     domain_transaction.Callback_Invoice
	transactionRepo    transactiontMocks.Repository
)

func TestMain(m *testing.M) {
	transactionService = service_transaction.NewTransactionService(&transactionRepo)
	transactionDomain = domain_transaction.Transaction{
		Transaction_Code:    "transaction-0b15cadb-dd2b-4cb2-8b29-47e334f4640e",
		ID_Customer:         "+6895631948686",
		Phone:               "+62895631948686",
		Amount:              15000,
		Category_Slug:       "pulsa",
		Detail_Product_Slug: "paket-xl-20rb",
		Payment_Id:          "62d5550b7c7cc1f5a367c5ce",
		Status:              "PENDING",
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
	}
	callbackDomain = domain_transaction.Callback_Invoice{
		ID:             "62bb210af01c3236811fc564",
		PaymentMethod:  "BANK",
		Status:         "PAID",
		PaidAmount:     15000,
		PaymentChannel: "BCA",
		PaidAt:         time.Time{},
	}
	detailDomain = domain_transaction.Detail_Transaction{
		Product_Detail_code: "paket-xl-20rb",
		Transaction_Code:    "transaction-0b15cadb-dd2b-4cb2-8b29-47e334f4640e",
		ID_Customer:         "12345",
		Customer_Name:       "Aziz",
		Price:               15000,
		Fee:                 2000,
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
	}
	paymentDomain = domain_transaction.Payment{
		Payment_Id: "62bb210af01c3236811fc564",
		Method:     "BANK",
		Channel:    "BCA",
		Paid_at:    time.Time{},
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	os.Exit(m.Run())
}

func TestGetTransactionsByPhone(t *testing.T) {
	t.Run("success get transaction by phone", func(t *testing.T) {
		transactionRepo.On("GetTransactionByPhone", mock.Anything).Return([]domain_transaction.Transaction{transactionDomain})
		sliceTransaction := transactionService.GetTransactionsByPhone(transactionDomain.Phone)
		assert.Equal(t, []domain_transaction.Transaction{transactionDomain}, sliceTransaction)
	})
}

func TestGetTransactionAll(t *testing.T) {
	t.Run("success get all transaction", func(t *testing.T) {
		transactionRepo.On("GetTransactions").Return([]domain_transaction.Transaction{transactionDomain})
		sliceTransaction := transactionService.GetTransactionAll()
		assert.Equal(t, []domain_transaction.Transaction{transactionDomain}, sliceTransaction)
	})
}

func TestAddDetailTransaction(t *testing.T) {
	t.Run("success add detail transaction", func(t *testing.T) {
		transactionRepo.On("StoreDetailTransaction", mock.Anything).Return(nil).Once()
		_, err := transactionService.AddDetailTransaction("paket-xl-20rb", detailDomain)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed add detail transaction", func(t *testing.T) {
		transactionRepo.On("StoreDetailTransaction", mock.Anything).Return(errors.New("failed store detail")).Once()
		_, err := transactionService.AddDetailTransaction("paket-xl-20rb", detailDomain)

		assert.Error(t, err)
		assert.Equal(t, err, err)
	})
}

func TestGetDetailTransaction(t *testing.T) {
	t.Run("success get detail transaction", func(t *testing.T) {
		transactionRepo.On("GetDetailTransaction", mock.Anything).Return(detailDomain, nil).Once()
		res, err := transactionService.GetDetailTransaction(transactionDomain.Transaction_Code)
		assert.NoError(t, err)
		assert.Equal(t, detailDomain, res)
	})
	t.Run("failed get detail transaction", func(t *testing.T) {
		transactionRepo.On("GetDetailTransaction", mock.Anything).Return(domain_transaction.Detail_Transaction{}, errors.New("error")).Once()
		res, err := transactionService.GetDetailTransaction(transactionDomain.Transaction_Code)
		assert.Error(t, err)
		assert.Equal(t, domain_transaction.Detail_Transaction{}, res)
	})
}

func TestAddTransaction(t *testing.T) {
	t.Run("success, add transaction", func(t *testing.T) {
		transactionRepo.On("StoreTransaction", mock.Anything).Return(nil).Once()
		item := xendit.InvoiceItem{
			Name:     "pulsa xl 20rb",
			Category: "Pulsa",
		}
		err := transactionService.AddTransaction(&xendit.Invoice{Customer: xendit.InvoiceCustomer{MobileNumber: transactionDomain.Phone}, Amount: 400, Items: []xendit.InvoiceItem{item}, ID: "1231231412311", Status: "PENDING"}, detailDomain)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed, add transaction", func(t *testing.T) {
		transactionRepo.On("StoreTransaction", mock.Anything).Return(errors.New("error add transaction")).Once()
		item := xendit.InvoiceItem{
			Name:     "pulsa xl 20rb",
			Category: "Pulsa",
		}
		err := transactionService.AddTransaction(&xendit.Invoice{Customer: xendit.InvoiceCustomer{MobileNumber: transactionDomain.Phone}, Amount: 400, Items: []xendit.InvoiceItem{item}, ID: "1231231412311", Status: "PENDING"}, detailDomain)

		assert.Error(t, err)
		assert.Equal(t, err, err)
	})
}

func TestEditTransaction(t *testing.T) {
	t.Run("success edit transaction", func(t *testing.T) {
		transactionRepo.On("GetTransactionByPaymentId", mock.Anything).Return(transactionDomain, nil).Once()
		transactionRepo.On("UpdateTransaction", mock.Anything).Return(nil).Once()
		transactionRepo.On("StorePayment", mock.Anything).Return(nil).Once()

		err := transactionService.EditTransaction(callbackDomain)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})

	t.Run("failed edit transaction", func(t *testing.T) {
		transactionRepo.On("GetTransactionByPaymentId", mock.Anything).Return(transactionDomain, errors.New("transaction not found")).Once()
		transactionRepo.On("UpdateTransaction", mock.Anything).Return(errors.New("failed update")).Once()
		transactionRepo.On("StorePayment", mock.Anything).Return(errors.New("failed add payment")).Once()

		err := transactionService.EditTransaction(domain_transaction.Callback_Invoice{})

		assert.Error(t, err)
		assert.Equal(t, err, err)
	})
}

func TestGetPayment(t *testing.T) {
	t.Run("success get payments", func(t *testing.T) {
		transactionRepo.On("GetPayment", mock.Anything).Return(paymentDomain).Once()
		res := transactionService.GetPayment("abcd-12345")

		assert.Equal(t, paymentDomain, res)
	})
	t.Run("failed get payments", func(t *testing.T) {
		transactionRepo.On("GetPayment", mock.Anything).Return(domain_transaction.Payment{}).Once()
		res := transactionService.GetPayment("")

		assert.Equal(t, domain_transaction.Payment{}, res)
	})
}

func TestGetFavoritesByPhone(t *testing.T) {
	t.Run("success get favorite by phone", func(t *testing.T) {
		transactionRepo.On("GetFavorite", mock.Anything, mock.Anything).Return([]domain_transaction.Transaction{transactionDomain}).Once()
		transactionRepo.On("GetTransactionByPaymentId", mock.Anything).Return(transactionDomain, nil)
		data := transactionService.GetFavoritesByPhone("pulsa", "62895631948686")
		assert.Equal(t, transactionDomain, data)
	})
	t.Run("success get favorite 2 by phone", func(t *testing.T) {
		transactionRepo.On("GetFavorite", mock.Anything, mock.Anything).Return([]domain_transaction.Transaction{transactionDomain, transactionDomain}).Once()
		transactionRepo.On("GetTransactionByPaymentId", mock.Anything).Return(transactionDomain, nil)
		transactionRepo.On("Count", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain.Payment_Id, 2)
		data := transactionService.GetFavoritesByPhone("pulsa", "62895631948686")
		assert.Equal(t, transactionDomain, data)
	})
}

func TestGetTransactionByPaymentId(t *testing.T) {
	t.Run("err get transaction with xendit", func(t *testing.T) {
		transactionRepo.On("GetTransactionByPaymentId", mock.Anything).Return(domain_transaction.Transaction{}, errors.New("bad request")).Once()
		data, err := transactionService.GetTransactionByPaymentId(transactionDomain.Payment_Id)
		assert.Error(t, err)
		assert.Equal(t, domain_transaction.Transaction{}, data)
	})
}

func TestCount(t *testing.T) {
	t.Run("count transaction", func(t *testing.T) {
		transactionRepo.On("Counts").Return(1).Once()
		data := transactionService.CountTransaction()
		assert.Equal(t, 1, data)
	})
}

// $ go test ./transaction/domain/abstraction_test.go -coverpkg=./transaction/service/...
// ok      command-line-arguments  0.389s  coverage: 93.1% of statements in ./transaction/service/...
