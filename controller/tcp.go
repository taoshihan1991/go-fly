package controller

import (
	"log"
	"net"
)
var clientTcpList = make([]net.Conn,0)
func NewTcpServer(tcpBaseServer string){
	listener, err := net.Listen("tcp", tcpBaseServer)
	if err != nil {
		log.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting", err.Error())
			return // 终止程序
		}
		clientTcpList=append(clientTcpList,conn)
	}
}
func PushServerTcp(str []byte){
	for index,conn:=range clientTcpList{
		_,err:=conn.Write(str)
		log.Println(index,err)
		if err!=nil{
			clientTcpList=append(clientTcpList[:index],clientTcpList[index+1:]...)
		}
	}
}