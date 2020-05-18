package libduitku

import "github.com/shopspring/decimal"

type (
	// ErrorResponse is struct response error of services
	ErrorResponse struct {
		Error   error  `json:"error"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	// DisbursementResponse struct model
	DisbursementResponse struct {
		UserID           int             `json:"userId,omitempty"`
		Balance          decimal.Decimal `json:"balance,omitempty"`
		EffectiveBalance decimal.Decimal `json:"effectiveBalance,omitempty"`
		Email            string          `json:"email,omitempty"`
		Reference        string          `json:"reference,omitempty"`
		BankCode         string          `json:"bankCode,omitempty"`
		BankAccount      string          `json:"bankAccount,omitempty"`
		AmountTransfer   decimal.Decimal `json:"amountTransfer,omitempty"`
		CustRefNumber    string          `json:"custRefNumber,omitempty"`
		AccountName      string          `json:"accountName,omitempty"`
		DisburseID       int             `json:"disburseId,omitempty"`
		ResponseCode     string          `json:"responseCode,omitempty"`
		ResponseDesc     string          `json:"responseDesc,omitempty"`
		Banks            []Banks         `json:"banks,omitempty"`
	}

	// Banks struct model
	Banks struct {
		BankCode string `json:"bankCode"`
		BankName string `json:"bankName"`
	}

	// CheckTransactionResponse struct model
	CheckTransactionResponse struct {
		MerchantCode  string `json:"merchantCode"`
		Amount        string `json:"amount"`
		Reference     string `json:"reference"`
		Fee           string `json:"fee"`
		StatusCode    string `json:"statusCode"`
		StatusMessage string `json:"statusMessage"`
	}

	// TransactionResponse struct model
	TransactionResponse struct {
		MerchantCode         string `json:"merchantCode"`
		Reference            string `json:"reference"`
		PaymentURL           string `json:"paymentUrl"`
		VirtualAccountNumber string `json:"vaNumber"`
		Amount               string `json:"amount"`
		QRString             string `json:"qrString"`
	}

	// CallbackResponse struct model
	CallbackResponse struct {
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

	// RedirectResponse struct model
	RedirectResponse struct {
		MerchantOrderID string `json:"merchantOrderId"`
		Reference       string `json:"reference"`
		ResultCode      string `json:"resultCode"`
	}
)
