package controller

import (
	"log"
	"net"
)
var clientTcpList = make(map[string]net.Conn)
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
		var remoteIpAddress = conn.RemoteAddr()
		clientTcpList[remoteIpAddress.String()]=conn
		log.Println(remoteIpAddress,clientTcpList)
		//clientTcpList=append(clientTcpList,conn)
	}
}
func PushServerTcp(str []byte){
	for ip,conn:=range clientTcpList{
		_,err:=conn.Write(str)
		log.Println(ip,err)
		if err!=nil{
			delete(clientTcpList,ip)
			//clientTcpList=append(clientTcpList[:index],clientTcpList[index+1:]...)
		}
	}
}