package wxpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewClientInvokePay(t *testing.T) {
	payInfo,err:=NewClientInvokePay(Appid,MechId,"3452",SignKey)
	js,err:=json.Marshal(payInfo)
	fmt.Println(string(js))
	fmt.Println(err)
}
