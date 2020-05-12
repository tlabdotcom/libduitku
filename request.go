package libduitku

type (
	// DuitkuTransactionRequest struct model
	DuitkuTransactionRequest struct {
		MerchantCode     string             `json:"merchantCode,required"`
		PaymentAmount    int                `json:"paymentAmount,required"`
		MerchantOrderID  string             `json:"merchantOrderId,required"`
		ProductDetails   string             `json:"productDetails,required"`
		Email            string             `json:"email,required"`
		AdditionalParam  string             `json:"additionalParam,omitempty"`
		PaymentMethod    string             `json:"paymentMethod,required"`
		MerchantUserInfo string             `json:"merchantUserInfo,omitompty"`
		CustomerVAName   string             `json:"customerVaName,required"`
		PhoneNumber      string             `json:"phoneNumber,required"`
		ItemDetails      []DuitkuItemDetail `json:"itemDetails,required"`
		ReturnURL        string             `json:"returnUrl,required"`
		CallbackURL      string             `json:"callbackUrl,required"`
		Signature        string             `json:"signature,required"`
		ExpiryPeriod     int                `json:"expiryPeriod,omitempty"` // 5, 10 or 60 (in minutes)
	}

	// DuitkuItemDetail struct models
	DuitkuItemDetail struct {
		Name     string `json:"name,required"`
		Quantity int    `json:"quantity,required"`
		Price    int    `json:"price,required"`
	}

	// DuitkuCheckTransactionRequest struct model
	DuitkuCheckTransactionRequest struct {
		MerchantCode    string `json:"merchantCode"`
		MerchantOrderID string `json:"merchantOrderId"`
		Signature       string `json:"signature"`
	}
)
