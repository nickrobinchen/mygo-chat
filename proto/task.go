package proto

type SuccessReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
