package zrouter

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// ping test 自定义路由
type Ping struct {
	znet.BaseRouter
}

// Ping Handle
func (this *Ping) Handle(request ziface.IRequest) {
	zlog.Debug("Call PingRouter Handle")
	zlog.Debug("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("Hello! This is Ping From Server!!!"))
	if err != nil {
		zlog.Error(err)
	}
}
