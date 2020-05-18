package libduitku

type (
	// DisbursementTransferRequest struct model
	DisbursementTransferRequest struct {
		DisburseID     int    `json:"disburseId"`
		UserID         int    `json:"userId"`
		Email          string `json:"email"`
		BankCode       string `json:"bankCode"`
		BankAccount    string `json:"bankAccount"`
		AmountTransfer int    `json:"amountTransfer"`
		AccountName    string `json:"accountName"`
		CustRefNumber  string `json:"custRefNumber"`
		Purpose        string `json:"purpose"`
		Timestamp      int64  `json:"timestamp"`
		Signature      string `json:"signature"`
	}

	// DisbursementInquiryRequest struct model
	DisbursementInquiryRequest struct {
		UserID         int    `json:"userId,omitempty"`
		DisburseID     string `json:"disburseId,omitempty"`
		AmountTransfer int    `json:"amountTransfer,omitempty"`
		BankAccount    string `json:"bankAccount,omitempty"`
		BankCode       string `json:"bankCode,omitempty"`
		Email          string `json:"email,omitempty"`
		Purpose        string `json:"purpose,omitempty"`
		Timestamp      int64  `json:"timestamp,omitempty"`
		SenderID       int    `json:"senderId,omitempty"`
		SenderName     string `json:"senderName,omitempty"`
		Signature      string `json:"signature,required"`
	}

	// TransactionRequest struct model
	TransactionRequest struct {
		MerchantCode     string       `json:"merchantCode,required"`
		PaymentAmount    int          `json:"paymentAmount,required"`
		MerchantOrderID  string       `json:"merchantOrderId,required"`
		ProductDetails   string       `json:"productDetails,required"`
		Email            string       `json:"email,required"`
		AdditionalParam  string       `json:"additionalParam,omitempty"`
		PaymentMethod    string       `json:"paymentMethod,required"`
		MerchantUserInfo string       `json:"merchantUserInfo,omitompty"`
		CustomerVAName   string       `json:"customerVaName,required"`
		PhoneNumber      string       `json:"phoneNumber,required"`
		ItemDetails      []ItemDetail `json:"itemDetails,required"`
		ReturnURL        string       `json:"returnUrl,required"`
		CallbackURL      string       `json:"callbackUrl,required"`
		Signature        string       `json:"signature,required"`
		ExpiryPeriod     int          `json:"expiryPeriod,omitempty"` // 5, 10 or 60 (in minutes)
	}

	// ItemDetail struct models
	ItemDetail struct {
		Name     string `json:"name,required"`
		Quantity int    `json:"quantity,required"`
		Price    int    `json:"price,required"`
	}

	// CheckTransactionRequest struct model
	CheckTransactionRequest struct {
		MerchantCode    string `json:"merchantCode"`
		MerchantOrderID string `json:"merchantOrderId"`
		Signature       string `json:"signature"`
	}
)
