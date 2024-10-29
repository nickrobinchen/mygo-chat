package proto

type LoginRequest struct {
	Name     string
	Password string
}

type LoginResponse struct {
	Code      int
	AuthToken string
}

type RegisterRequest struct {
	Name     string
	Password string
}

type RegisterReply struct {
	Code      int
	AuthToken string
}

type GetUserInfoRequest struct {
	UserId int
}

type CheckAuthRequest struct {
	AuthToken string
}

type LogoutRequest struct {
	AuthToken string
}

type ConnectRequest struct {
	AuthToken string `json:"authToken"`
	RoomId    int    `json:"roomId"`
	ServerId  string `json:"serverId"`
}

type ConnectReply struct {
	UserId int
}
type DisConnectRequest struct {
	RoomId int
	UserId int
}

type DisConnectReply struct {
	Has bool
}
type CheckAuthResponse struct {
	Code     int
	UserId   int
	UserName string
}

type GetUserInfoResponse struct {
	Code     int
	UserId   int
	UserName string
}

type LogoutResponse struct {
	Code int
}

type Send struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	FromUserId   int    `json:"fromUserId"`
	FromUserName string `json:"fromUserName"`
	ToUserId     int    `json:"toUserId"`
	ToUserName   string `json:"toUserName"`
	RoomId       int    `json:"roomId"`
	Op           int    `json:"op"`
	CreateTime   string `json:"createTime"`
}
