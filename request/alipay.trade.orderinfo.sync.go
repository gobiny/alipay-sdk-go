package request

const AlipayTradeOrderInfoSyncMethod = "alipay.trade.orderinfo.sync"

type AlipayTradeOrderInfoSyncRequest struct {
	TradeNo       string `json:"trade_no"`
	OrigRequestNo string `json:"orig_request_no"`
	OutRequestNo  string `json:"out_request_no"`
	BizType       string `json:"biz_type"`
	OrderBizInfo  string `json:"order_biz_info"`
}
