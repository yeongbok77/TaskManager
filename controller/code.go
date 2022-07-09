package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:    "操作成功",
	CodeServerBusy: "系统繁忙，请稍后再试",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
