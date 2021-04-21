package test

import (
	"testing"

	"github.com/starudream/alipay-sdk-go/request"
)

func TestAlipayFundAuthOrderUnfreeze(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOrderUnfreezeRequest{
		AuthNo:       "20180901000000AFAOU00001",
		OutRequestNo: "20180901000000AFAOU00001",
		Amount:       "1.11",
		Remark:       "解冻 1.11 元",
	}
	client.SendRequest(request.AlipayFundAuthOrderUnfreezeMethod, data)
}
