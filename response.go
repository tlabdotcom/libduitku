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
		Amount               int    `json:"amount"`
		QRString             string `json:"qrString"`
	}
)
