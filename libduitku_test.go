package libduitku

import "testing"

func initClient() CoreDuitku {
	c := NewClient()
	c.Host = "https://sandbox.duitku.com",
	c.MerchantCode = ""
	c.APIKey = ""
	c.ReturnURL = ""
	c.CallbackURL = ""
	return CoreDuitku{
		Client: c,
	}
}

func TestRequestTransaction(t *testing.T) {
	c := initClient()
}
