package helper_xendit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Data struct {
	ID                        string    `json:"id"`
	ExternalID                string    `json:"external_id"`
	UserID                    string    `json:"user_id"`
	Status                    string    `json:"status"`
	MerchantName              string    `json:"merchant_name"`
	MerchantProfilePictureURL string    `json:"merchant_profile_picture_url"`
	Amount                    int       `json:"amount"`
	PayerEmail                string    `json:"payer_email"`
	Description               string    `json:"description"`
	ExpiryDate                time.Time `json:"expiry_date"`
	InvoiceURL                string    `json:"invoice_url"`
	AvailableBanks            []struct {
		BankCode          string `json:"bank_code"`
		CollectionType    string `json:"collection_type"`
		TransferAmount    int    `json:"transfer_amount"`
		BankBranch        string `json:"bank_branch"`
		AccountHolderName string `json:"account_holder_name"`
		IdentityAmount    int    `json:"identity_amount"`
	} `json:"available_banks"`
	AvailableRetailOutlets []struct {
		RetailOutletName string `json:"retail_outlet_name"`
	} `json:"available_retail_outlets"`
	AvailableEwallets []struct {
		EwalletType string `json:"ewallet_type"`
	} `json:"available_ewallets"`
	AvailableDirectDebits []struct {
		DirectDebitType string `json:"direct_debit_type"`
	} `json:"available_direct_debits"`
	AvailablePaylaters      []interface{} `json:"available_paylaters"`
	ShouldExcludeCreditCard bool          `json:"should_exclude_credit_card"`
	ShouldSendEmail         bool          `json:"should_send_email"`
	Created                 time.Time     `json:"created"`
	Updated                 time.Time     `json:"updated"`
	Currency                string        `json:"currency"`
	Items                   []struct {
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Category string `json:"category"`
	} `json:"items"`
	Fees []struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"fees"`
	Customer struct {
		GivenNames   string `json:"given_names"`
		Email        string `json:"email"`
		MobileNumber string `json:"mobile_number"`
	} `json:"customer"`
	CustomerNotificationPreference struct {
		InvoiceCreated  []string `json:"invoice_created"`
		InvoiceReminder []string `json:"invoice_reminder"`
		InvoicePaid     []string `json:"invoice_paid"`
		InvoiceExpired  []string `json:"invoice_expired"`
	} `json:"customer_notification_preference"`
}

func DetailTransactionXendit(paymentID string) interface{} {

	url := "https://api.xendit.co/v2/invoices/" + paymentID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic eG5kX2RldmVsb3BtZW50X0kwZ3VLNWJPY1FCM0FWUThwWVVYTXRYbHRzVnZmcXN5UFUyZHoxUkp2VEROVnJzTFZ4cUM4SzBLSmMzWWhsWkU6")
	req.Header.Add("Cookie", "incap_ses_7264_2182539=TDK6VWyh7gJagAQpnenOZJlH1mIAAAAACRNwEzYd9ElMYWlIWUUQ7Q==; nlbi_2182539=66GMWKBkMiAwMfOrjjCKbQAAAACDB8IL0atKp47HSWHkW8Ka")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := Data{}
	json.Unmarshal(body, &result)
	return result
}
