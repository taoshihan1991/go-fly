package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
)

var clientTcpList = make(map[string]net.Conn)

func NewTcpServer(tcpBaseServer string) {
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
		clientTcpList[remoteIpAddress.String()] = conn
		log.Println(remoteIpAddress, clientTcpList)
		//clientTcpList=append(clientTcpList,conn)
	}
}
func PushServerTcp(str []byte) {
	for ip, conn := range clientTcpList {
		line := append(str, []byte("\r\n")...)
		_, err := conn.Write(line)
		log.Println(ip, err)
		if err != nil {
			conn.Close()
			delete(clientTcpList, ip)
			//clientTcpList=append(clientTcpList[:index],clientTcpList[index+1:]...)
		}
	}
}
func DeleteOnlineTcp(c *gin.Context) {
	ip := c.Query("ip")
	for ipkey, conn := range clientTcpList {
		if ip == ipkey {
			conn.Close()
			delete(clientTcpList, ip)
		}
		if ip == "all" {
			conn.Close()
			delete(clientTcpList, ipkey)
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
