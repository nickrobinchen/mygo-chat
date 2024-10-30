package connect

import (
	"mygochat/proto"
	"net"

	"github.com/gorilla/websocket"
)

// in fact, Channel it's a user Connect session
type Channel struct {
	Room      *Room
	Next      *Channel
	Prev      *Channel
	broadcast chan *proto.Msg
	userId    int
	conn      *websocket.Conn
	connTcp   *net.TCPConn
}

func NewChannel(size int) (c *Channel) {
	c = new(Channel)
	c.broadcast = make(chan *proto.Msg, size)
	c.Next = nil
	c.Prev = nil
	return
}

// 函数分析
// select 语句:
// select 语句用于处理多个通道操作。在这个例子中，它用于尝试将消息发送到 ch.broadcast 通道。
// case ch.broadcast <- msg: 这是一个非阻塞的发送操作。如果 ch.broadcast 通道有空间接收新消息，消息将被发送。
// default:: 如果 ch.broadcast 通道没有空间（即通道已满），则不执行任何操作，控制流将跳过此 case。
// 返回值:
// return: 函数在尝试发送消息后返回，返回的错误（err）始终为 nil，因为此函数没有显式处理错误或发送失败的情况。
// 可能的作用
// 非阻塞推送: 该函数的设计使得推送消息的过程是非阻塞的。如果通道满了，消息不会被发送，也不会阻塞调用者。这适合于需要快速处理消息的场景，例如广播通知。
// 适合高并发场景: 由于采用非阻塞的方式，该函数可以在高并发的环境中高效运行，避免因等待通道可用而导致的延迟。
// 广播消息: 函数的命名和 broadcast 通道的使用暗示它可能用于将消息广播给多个接收者，适合用于聊天室、消息推送等应用场景。
// 总结
// 这个 Push 函数实现了一种高效的、非阻塞的消息推送机制，适用于需要快速广播消息的应用场景。
func (ch *Channel) Push(msg *proto.Msg) (err error) {
	// select {
	// case <- chan1:
	// 	// 如果 chan1 成功读到数据，则进行该 case 处理语句
	// case chan2 <- 1:
	// 	// 如果成功向 chan2 写入数据，则进行该 case 处理语句
	// default:
	// 	// 如果上面都没有成功，则进入default处理流程
	//}
	select {
	case ch.broadcast <- msg:
	default:
	}
	return
}
