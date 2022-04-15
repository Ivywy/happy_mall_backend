package enum

type PayType int

const (
	Bank   PayType = 0
	WeChat PayType = 1
	AllPay PayType = 2
)

func (p PayType) String() string {
	switch p {
	case Bank:
		return "银行卡"
	case WeChat:
		return "微信"
	case AllPay:
		return "支付宝"
	default:
		return "UNKNOWN"
	}
}
