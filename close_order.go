package wxpay

import (
	"encoding/xml"
	"fmt"
)

type CloseOrderRequest struct {
	QueryOrderRequest
}

type CloseOrderResponse struct {
	PublicResponse
}

func (req CloseOrderRequest)CloseOrder(signKey string)(response CloseOrderResponse,err error){
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
	resp,err:=HttpPost(wxpayCloseOrderUrl,byteReq)
	if err != nil {
		return
	}

	err  = xml.Unmarshal(resp,&response)

	return
}