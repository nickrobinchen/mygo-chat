package proto

type SuccessReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RedisMsg struct {
	Op           int               `json:"op"`
	ServerId     string            `json:"serverId,omitempty"`
	RoomId       int               `json:"roomId,omitempty"`
	UserId       int               `json:"userId,omitempty"`
	Msg          []byte            `json:"msg"`
	Count        int               `json:"count"`
	RoomUserInfo map[string]string `json:"roomUserInfo"`
}
