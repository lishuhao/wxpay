package wxpay

import (
	"encoding/xml"
	"fmt"
)
//查询订单
type QueryOrderRequest struct {
	XMLName xml.Name `xml:"xml"`
	Appid string `xml:"appid"`
	MchId string `xml:"mch_id"`
	OutTradeNo string `xml:"out_trade_no"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
}

type QueryOrderResponse struct {
	PublicResponse

	ResultCode string `xml:"result_code"`
	TradeState string `xml:"trade_state"`
	TotalFee string `xml:"total_fee"` //int
	FeeType string `xml:"fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo string `xml:"out_trade_no"`
	TimeEnd string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}

//统一下单
func (req QueryOrderRequest)Query(signKey string) (response QueryOrderResponse,err error) {
	req.NonceStr = nonceStr()
	sign,err:=signReq(req,signKey)
	if err != nil {
		return
	}
	req.Sign = sign
	byteReq,err:=xml.MarshalIndent(req,"  ","  ")
	if err != nil {
		return
	}
	fmt.Println(string(byteReq))
	resp,err:=HttpPost(wxpayOrderQueryUrl,byteReq)
	if err != nil {
		return
	}

	err  = xml.Unmarshal(resp,&response)

	return
}
