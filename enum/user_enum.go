package enum

type ResponseType int

const (
	OperateOK   ResponseType = 200
	OperateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOK:
		return "ok"
	case OperateFail:
		return "fail"
	default:
		return "UNKNOWN"
	}
}
