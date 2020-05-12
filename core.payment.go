package libduitku

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	pathRequestTransaction = "webapi/api/merchant/v2/inquiry"
)

// RequestTransaction is used to request payment to Duitku
func (core *CoreDuitku) RequestTransaction(req DuitkuTransactionRequest) (res DuitkuTransactionResponse, err error) {
	req.MerchantCode = core.Client.MerchantCode
	req.Signature = HashToMD5(fmt.Sprintf("%s%s%v%s",
		core.Client.MerchantCode,
		req.MerchantOrderID,
		req.PaymentAmount,
		core.Client.APIKey,
	))
	req.CallbackURL = core.Client.CallbackURL
	req.ReturnURL = core.Client.ReturnURL
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, pathRequestTransaction, headers, buf, &res)
	if err != nil {
		return
	}
	return
}
