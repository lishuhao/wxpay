package wxpay

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestUnifyOrderRequest(t *testing.T)  {
	orderNo:=strconv.FormatInt(time.Now().Unix(),10)
	fmt.Println(orderNo)
	req:=UnifiedOrderRequest{
		Appid:Appid,
		MchId:MechId,
		TotalFee:"2324",
		Body:"Body",
		OutTradeNo:orderNo,
		SpbillCreateIp:"127.0.0.1",
		NotifyUrl:"http://baidu.com",
	}
	res,err := req.Request(SignKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v",res)
}

func TestUnifyOrderResponse(t *testing.T)  {
	resp:=`<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[APP]]></trade_type>
</xml>`
	res:=UnifiedOrderResponse{}
	err := xml.Unmarshal([]byte(resp),&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v",res)
}
