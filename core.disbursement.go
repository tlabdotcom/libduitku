package libduitku

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	pathDisbursementInquirySandbox  = "webapi/api/disbursement/inquirysandbox"
	pathDisbursementInquiryProd     = "webapi/api/disbursement/inquiry"
	pathDisbursementTransferSandbox = "webapi/api/disbursement/transfersandbox"
	pathDisbursementTransferProd    = "webapi/api/disbursement/transfer"
	pathInquiryStatus               = "webapi/api/disbursement/inquirystatus"
	pathInquiryCheckBalance         = "webapi/api/disbursement/checkbalance"
	pathInquiryListBank             = "webapi/api/disbursement/listBank"
)

// InquiryListBank is used to inquery list bank
func (core *CoreDuitku) InquiryListBank() (res DisbursementResponse, err error) {
	var req DisbursementInquiryRequest
	req.UserID = core.Client.DisbursementUserID
	req.Email = core.Client.DisbursementEmail
	req.Timestamp = core.makeTimestamp()
	sign := fmt.Sprintf("%s%v%s",
		req.Email,
		req.Timestamp,
		core.Client.DisbursementKey,
	)
	req.Signature = core.hashToSHA256(sign)
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, pathInquiryListBank, headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// InquiryCheckBalance is used to ...
func (core *CoreDuitku) InquiryCheckBalance() (res DisbursementResponse, err error) {
	var req DisbursementInquiryRequest
	req.UserID = core.Client.DisbursementUserID
	req.Email = core.Client.DisbursementEmail
	req.Timestamp = core.makeTimestamp()
	sign := fmt.Sprintf("%s%v%s",
		req.Email,
		req.Timestamp,
		core.Client.DisbursementKey,
	)
	req.Signature = core.hashToSHA256(sign)
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, pathInquiryCheckBalance, headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// InquiryStatus is used to chech status disbursement
func (core *CoreDuitku) InquiryStatus(disburseID string) (res DisbursementResponse, err error) {
	var req DisbursementInquiryRequest
	req.Timestamp = core.makeTimestamp()
	req.UserID = core.Client.DisbursementUserID
	req.Email = core.Client.DisbursementEmail
	req.DisburseID = disburseID
	sign := fmt.Sprintf("%s%v%s%s",
		req.Email,
		req.Timestamp,
		req.DisburseID,
		core.Client.DisbursementKey,
	)
	req.Signature = core.hashToSHA256(sign)
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, pathInquiryStatus, headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// DisbursementTransfer is used to ...
func (core *CoreDuitku) DisbursementTransfer(req DisbursementTransferRequest) (res DisbursementResponse, err error) {
	req.Timestamp = core.makeTimestamp()
	req.UserID = core.Client.DisbursementUserID
	req.Email = core.Client.DisbursementEmail
	sign := fmt.Sprintf("%s%v%s%s%s%s%v%s%v%s",
		req.Email,
		req.Timestamp,
		req.BankCode,
		req.BankAccount,
		req.AccountName,
		req.CustRefNumber,
		req.AmountTransfer,
		req.Purpose,
		req.DisburseID,
		core.Client.DisbursementKey)
	req.Signature = core.hashToSHA256(sign)
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, core.isSanbox(pathDisbursementTransferProd, pathDisbursementTransferSandbox), headers, buf, &res)
	if err != nil {
		return
	}
	return
}

// DisbursementInquiry is used to ...
func (core *CoreDuitku) DisbursementInquiry(req DisbursementInquiryRequest) (res DisbursementResponse, err error) {
	req.UserID = core.Client.DisbursementUserID
	req.Email = core.Client.DisbursementEmail
	req.Timestamp = core.makeTimestamp()
	sign := fmt.Sprintf("%s%v%s%s%v%s%s", req.Email, req.Timestamp, req.BankCode, req.BankAccount, req.AmountTransfer, req.Purpose, core.Client.DisbursementKey)
	req.Signature = core.hashToSHA256(sign)

	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": fmt.Sprintf("%x", buf.Len()),
	}
	err = core.Call(fasthttp.MethodPost, core.isSanbox(pathDisbursementInquiryProd, pathDisbursementInquirySandbox), headers, buf, &res)
	if err != nil {
		return
	}
	return
}
