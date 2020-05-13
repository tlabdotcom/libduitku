package libduitku

import (
	"fmt"
	"testing"
	"time"
)

func initClient() CoreDuitku {
	c := NewClient()
	c.Host = "https://sandbox.duitku.com"
	c.MerchantCode = "code"
	c.APIKey = "api-key"
	c.ReturnURL = "https://yourweb.com/return-url"
	c.CallbackURL = "https://yourweb.com/callback"
	return CoreDuitku{
		Client: c,
	}
}

func TestRequestTransaction(t *testing.T) {
	c := initClient()
	resData := DuitkuTransactionRequest{
		PaymentMethod:   CreditCard,
		PaymentAmount:   500000,
		MerchantOrderID: fmt.Sprintf("KESAN%v", time.Now().Unix()),
		ProductDetails:  "Transaksi pembayaran kesan",
		Email:           "jihar@tlab.co.id",
		CustomerVAName:  "Jihar Al Gifari",
		PhoneNumber:     "08992939929",
		ItemDetails: []DuitkuItemDetail{
			{
				Name:     "Item 1",
				Price:    500000,
				Quantity: 1,
			},
		},
	}
	resp, err := c.RequestTransaction(resData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestCheckTransaction(t *testing.T) {
	c := initClient()
	resp, err := c.CheckTransaction("KESAN1589352383")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
