package controller

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var oldFolders map[string]int

//推送新邮件到达
func PushMailServer(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		mailServer := tools.GetMailServerFromCookie(r)
		var msg []byte
		if mailServer == nil {
			msg, _ = json.Marshal(tools.JsonResult{Code: 400, Msg: "验证失败"})
			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		} else {
			folders := tools.GetMailNum(mailServer.Server, mailServer.Email, mailServer.Password)
			for name, num := range folders {
				if oldFolders[name] != num {
					result := make(map[string]interface{})
					result["folder_name"] = name
					result["new_num"] = num - oldFolders[name]
					msg, _ := json.Marshal(tools.JsonListResult{
						JsonResult: tools.JsonResult{Code: 200, Msg: "获取成功"},
						Result:     result,
					})
					c.WriteMessage(mt, msg)
				}
			}
			oldFolders = folders
		}
	}
}
