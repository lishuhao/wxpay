package wxpay

const (
	xmlContentType = "application/xml; charset=utf-8"

	wxpayUnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	wxpayOrderQueryUrl   = "https://api.mch.weixin.qq.com/pay/orderquery"
	wxpayCloseOrderUrl   = "https://api.mch.weixin.qq.com/pay/closeorder"
)

const (
	signType = "MD5"
	tradeType = "APP"
)