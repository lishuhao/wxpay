微信支付sdk
测试文件中的Appid,MechId,SignKey 属于敏感数据，若要执行测试文件，请新建demo_data.go 文件，并定义以下常量

//敏感数据，做测试用
const(
	Appid = "你的App id"
	MechId = "你的 商户号"
	SignKey = "你的密钥"
)
