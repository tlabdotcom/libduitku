package libduitku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func initClient() CoreDuitku {
	c := NewClient()
	c.Host = "https://sandbox.duitku.com"
	c.MerchantCode = "D6080"
	c.APIKey = "79915d65a1e918c054b2724bdfe5d574"
	c.ReturnURL = "https://api.nedinasek.com/stag/e2pay/from/duitku/return-url"
	c.CallbackURL = "https://api.nedinasek.com/stag/e2pay/from/duitku/callback"
	c.DisbursementUserID = 3551
	c.DisbursementEmail = "test@chakratechnology.com"
	c.DisbursementKey = "de56f832487bc1ce1de5ff2cfacf8d9486c61da69df6fd61d5537b6b7d6d354d"
	return CoreDuitku{
		Client: c,
	}
}

func TestInquiryStatus(t *testing.T) {
	c := initClient()
	resp, err := c.InquiryStatus("6261")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestDisbursementTransfer(t *testing.T) {
	c := initClient()
	req := DisbursementTransferRequest{
		DisburseID:     6261,
		BankAccount:    "8760673566",
		AmountTransfer: 10000,
		BankCode:       "014",
		Purpose:        "Pencairan saldo ",
	}
	resp, err := c.DisbursementTransfer(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp) // {0 0 0 test@chakratechnology.com  014 8760673566 10000 00006261 Test Account 0 00 Success []}
}

func TestDisbursementInquiry(t *testing.T) {
	c := initClient()
	req := DisbursementInquiryRequest{
		BankAccount:    "8760673566",
		AmountTransfer: 10000,
		BankCode:       "014",
		SenderID:       time.Now().Minute() + rand.Int(),
		SenderName:     "Jihar",
		Purpose:        "Pencairan saldo kesan",
	}
	resp, err := c.DisbursementInquiry(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp) // {0 0 0 test@chakratechnology.com  014 8760673566 10000 00006261 Test Account 6261 00 Success []}
}

func TestInquiryCheckBalance(t *testing.T) {
	c := initClient()
	resp, err := c.InquiryCheckBalance()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestInquiryListBank(t *testing.T) {
	c := initClient()
	resp, err := c.InquiryListBank()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.Banks)
}

func TestRequestTransaction(t *testing.T) {
	c := initClient()
	resData := TransactionRequest{
		PaymentMethod:   CreditCard,
		PaymentAmount:   500000,
		MerchantOrderID: fmt.Sprintf("KESAN%v", time.Now().Unix()),
		ProductDetails:  "Transaksi pembayaran kesan",
		Email:           "jihar@tlab.co.id",
		CustomerVAName:  "Jihar Al Gifari",
		PhoneNumber:     "08992939929",
		ItemDetails: []ItemDetail{
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
