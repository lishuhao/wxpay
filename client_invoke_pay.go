package wxpay

import (
	"strconv"
	"time"
)

type ClientInvokePay struct {
	Appid string `xml:"appid" json:"appid"`
	Partnerid string `xml:"partnerid" json:"partnerid"`
	Prepayid string `xml:"prepayid" json:"prepayid"`
	Package string `xml:"package" json:"package"`
	Noncestr string `xml:"noncestr" json:"noncestr"`
	Timestamp string `xml:"timestamp" json:"timestamp"`
	Sign string `xml:"sign" json:"sign"`
}
//APP端调起支付的参数列表 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12
func NewClientInvokePay(appId,partnerId,prepayId,signKey string)(ClientInvokePay,error)  {
	req:= ClientInvokePay{
		Appid:appId,
		Partnerid:partnerId,
		Prepayid:prepayId,
		Package:"Sign=WXPay",
		Noncestr:nonceStr(),
		Timestamp:strconv.FormatInt(time.Now().Unix(),10),
	}
	sign,err := signReq(req,signKey)
	if err != nil {
		return req,err
	}
	req.Sign = sign
	return req,nil
}
