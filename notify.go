package wxpay

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type NotifyRequest struct {
	PublicResponse

	ResultCode string `xml:"result_code"`
	OutTradeNo string `xml:"out_trade_no"`
	TotalFee string `xml:"total_fee"`//int 单位：分
	FeeType string `xml:"fee_type"`
	Attach string `xml:"attach"`
}

type NotifyResponse struct {
	XMLName xml.Name `xml:"xml"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
}

func AccessNotify(req http.Request) (notifyReq NotifyRequest,err error) {
	body,err:=ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(body,&notifyReq)
	return
}