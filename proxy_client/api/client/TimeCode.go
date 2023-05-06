package client

type TimeOutCode int

const (
	TimeOutServerInitStart TimeOutCode = 2000 + iota
	TimeOutServerInitSuccess
	TimeOutServerInitError
)

var codeMsgMap = map[TimeOutCode]string{
	TimeOutServerInitStart:   "创建开始",
	TimeOutServerInitSuccess: "创建成功",
	TimeOutServerInitError:   "创建失败",
}

func (r TimeOutCode) GetCode() string {
	v, ok := codeMsgMap[r]
	if !ok {
		return v
	}
	return v
}
