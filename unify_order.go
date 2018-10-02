package wxpay

import (
	"encoding/xml"
	"errors"
	"strconv"
	"time"
)

type UnifiedOrderRequest struct {
	XMLName xml.Name `xml:"xml"`
	Appid string `xml:"appid"`
	MchId string `xml:"mch_id"`
	Body string `xml:"body"`
	OutTradeNo string `xml:"out_trade_no"`
	TotalFee string `xml:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	NotifyUrl string `xml:"notify_url"`
	Sign string `xml:"sign"`
	TimeExpire string `xml:"time_expire"`

	TradeType string `xml:"trade_type"`
	NonceStr string `xml:"nonce_str"`
	SignType string `xml:"sign_type,omitempty"`
}

type UnifiedOrderResponse struct {
	PublicResponse
	ResultCode string `xml:"result_code"`
	PrepayId string `xml:"prepay_id"`
}

type PublicResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	ErrCode string `xml:"err_code"`
	ErrCodeDesc string `xml:"err_code_desc"`
}

//统一下单
func (req UnifiedOrderRequest)Request(signKey string) (response UnifiedOrderResponse,err error) {
	if req.Appid == ""{
		err = errors.New("appid 不能为空")
		return
	}
	_,err=strconv.ParseInt(req.TotalFee,10,64)
	if err != nil {
		err = errors.New("TotalFee 格式不正确")
		return
	}
	req.NonceStr = nonceStr()
	req.TradeType = tradeType
	req.SignType = signType
	req.TimeExpire = time.Now().Add(time.Minute * 20).Format("20060102150405")//30分钟后订单过期

	sign,err:=signReq(req,signKey)
	if err != nil {
		return
	}
	req.Sign = sign
	byteReq,err:=xml.Marshal(req)
	if err != nil {
		return
	}
	resp,err:=HttpPost(wxpayUnifiedOrderUrl,byteReq)
	if err != nil {
		return
	}

	err  = xml.Unmarshal(resp,&response)

	return
}

func (resp UnifiedOrderResponse)Success() bool {
	return resp.ReturnCode == "SUCCESS" && resp.ResultCode == "SUCCESS"
}