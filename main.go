/**
* @Author: Gandalfwang
* @Date: 2022/10/13
* @Mail: 516575345@qq.com
 */
package main

import (
	"time"

	"TCPMySQLServer/util"
	router "TCPMySQLServer/zrouter"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

// 创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	time0 := time.Now().Format("2006-01-02 15:04:05")
	zlog.Debug("DoConnecionBegin is Called ... ")

	// 设置链接属性，在连接创建之后
	zlog.Debug("Set conn Name, Home done!")
	conn.SetProperty("CreateTime", time0)

	// err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	// if err != nil {
	// 	zlog.Error(err)
	// }
}

// 连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	//在连接销毁之前，查询conn的属性
	if time0, err := conn.GetProperty("CreateTime"); err == nil {
		zlog.Error("Conn Property CreateTime = ", time0)
	}

	zlog.Debug("DoConneciotnLost is Called ... ")
}

func main() {
	// init util
	util.Init()

	//创建一个server句柄
	s := znet.NewServer()

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//配置路由
	s.AddRouter(0, &router.Ping{})
	s.AddRouter(1, &router.Business{})

	//开启服务
	s.Serve()
}
