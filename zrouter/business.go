package zrouter

import (
	"encoding/json"

	"TCPMySQLServer/dao"
	"TCPMySQLServer/vo"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// ping test 自定义路由
type Business struct {
	znet.BaseRouter
}

// Business Handle
func (this *Business) Handle(request ziface.IRequest) {
	zlog.Debug("Call Business Handle")
	msgId0 := request.GetMsgID()
	msgData := request.GetData()
	zlog.Debug("recv from client : msgId=", msgId0, ", data=", string(msgData))

	var book vo.Book
	err := json.Unmarshal(msgData, &book)
	if err != nil {
		zlog.Errorf("%s|%v", msgId0, err)
	}

	result0 := dao.Query(book.Id)
	replyMsg0, err := json.Marshal(result0)
	if err != nil {
		zlog.Errorf("Marshal|%+v|%v", result0, err)
	}

	err = request.GetConnection().SendBuffMsg(0, replyMsg0)
	if err != nil {
		zlog.Errorf("%s|%v", msgId0, err)
	}
}
