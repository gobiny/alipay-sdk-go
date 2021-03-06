package request

const AlipayFundAuthOrderAppFreezeMethod = "alipay.fund.auth.order.app.freeze"

type AlipayFundAuthOrderAppFreeze struct {
	//AlipaySdk         string `json:"alipay_sdk"`
	OutOrderNo        string `json:"out_order_no"`        // 商户授权资金订单号 ,不能包含除中文、英文、数字以外的字符，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo      string `json:"out_request_no"`      // 商户本次资金操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OrderTitle        string `json:"order_title"`         // 业务订单的简单描述，如商品名称等长度不超过100个字母或50个汉字
	Amount            string `json:"amount"`              // 需要冻结的金额，单位为：元（人民币），精确到小数点后两位取值范围：[0.01,100000000.00]
	ProductCode       string `json:"product_code"`        // 销售产品码，新接入线上预授权的业务，本字段取值固定为PRE_AUTH_ONLINE 。
	PayeeUserId       string `json:"payee_user_id"`       // 收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId      string `json:"payee_logon_id"`      // 收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayTimeout        string `json:"pay_timeout"`         // 该笔订单允许的最晚付款时间，逾期将关闭该笔订单取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m如果为空，默认15m
	ExtraParam        string `json:"extra_param"`         // 业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.授权业务对应的类目，key为category，value由支付宝分配，比如充电桩业务传 "CHARGE_PILE_CAR"； 2. 外部商户的门店编号，key为outStoreCode，可选； 3. 外部商户的门店简称，key为outStoreAlias，可选。
	EnablePayChannels string `json:"enable_pay_channels"` //  商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。
}
