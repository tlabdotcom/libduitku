package libduitku

type (
	// ErrorResponse is struct response error of services
	ErrorResponse struct {
		Error   error  `json:"error"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	// DuitkuCheckTransactionResponse struct model
	DuitkuCheckTransactionResponse struct {
		MerchantCode  string `json:"merchantCode"`
		Amount        string `json:"amount"`
		Reference     string `json:"reference"`
		Fee           string `json:"fee"`
		StatusCode    string `json:"statusCode"`
		StatusMessage string `json:"statusMessage"`
	}

	// DuitkuTransactionResponse struct model
	DuitkuTransactionResponse struct {
		MerchantCode         string `json:"merchantCode"`
		Reference            string `json:"reference"`
		PaymentURL           string `json:"paymentUrl"`
		VirtualAccountNumber string `json:"vaNumber"`
		Amount               string `json:"amount"`
		QRString             string `json:"qrString"`
	}

	// DuitkuCallbackResponse struct model
	DuitkuCallbackResponse struct {
		MerchantCode    string `json:"merchantCode"`
		Amount          string `json:"amount"`
		MerchantOrderID string `json:"merchantOrderId"`
		ProductDetail   string `json:"productDetail"`
		AdditionalParam string `json:"additionalParam"`
		PaymentCode     string `json:"paymentCode"`
		ResultCode      string `json:"resultCode"`
		MerchantUserID  string `json:"merchantUserId"`
		Reference       string `json:"reference"`
		Signature       string `json:"signature"`
	}

	// DuitkuRedirectResponse struct model
	DuitkuRedirectResponse struct {
		MerchantOrderID string `json:"merchantOrderId"`
		Reference       string `json:"reference"`
		ResultCode      string `json:"resultCode"`
	}
)
