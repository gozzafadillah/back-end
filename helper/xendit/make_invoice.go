package helper_xendit

import (
	"errors"
	"fmt"

	domain_products "ppob/products/domain"
	domain_transaction "ppob/transaction/domain"
	domain_users "ppob/users/domain"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func Xendit_Invoice(DetailTransaction domain_transaction.Detail_Transaction, product domain_products.Detail_Product, user domain_users.Users, cat string) (*xendit.Invoice, error) {
	xendit.Opt.SecretKey = "xnd_development_I0guK5bOcQB3AVQ8pYUXMtXltsVvfqsyPU2dz1RJvTDNVrsLVxqC8K0KJc3YhlZE"

	customer := xendit.Customer{
		GivenNames:   user.Name,
		Email:        user.Email,
		MobileNumber: user.Phone,
	}

	invoiceCustomer := xendit.InvoiceCustomer{
		GivenNames:   customer.GivenNames,
		Email:        customer.Email,
		MobileNumber: customer.MobileNumber,
	}
	item := xendit.InvoiceItem{
		Name:     product.Name,
		Price:    float64(product.Price),
		Quantity: 1,
		Category: cat,
	}
	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: float64(DetailTransaction.Fee),
	}

	NotificationType := []string{"whatsapp", "email", "sms"}
	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID:                     DetailTransaction.Transaction_Code,
		Amount:                         item.Price + fee.Value,
		PayerEmail:                     "admin@bayeue.com",
		Description:                    item.Name,
		Items:                          []xendit.InvoiceItem{item},
		Customer:                       invoiceCustomer,
		Fees:                           []xendit.InvoiceFee{fee},
		CustomerNotificationPreference: customerNotificationPreference,
		// success redirect di mobile
		// SuccessRedirectURL: ,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return resp, errors.New("internal server error, xendit")
	}
	fmt.Println("response ", resp)
	return resp, nil

}
