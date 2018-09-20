package wxpay

import (
	"fmt"
	"testing"
)

func TestQueryOrderRequest_Query(t *testing.T) {
	req:=QueryOrderRequest{
		Appid:Appid,
		MchId:MechId,
		OutTradeNo:"1537354134",
	}

	resp,err:=req.Query(SignKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v",resp)
}
