package controller
import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/taoshihan1991/imaptool/tools"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{}
var oldFolders map[string]int
//推送新邮件到达
func PushMailServer(w http.ResponseWriter, r *http.Request){
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
			err = c.WriteMessage(mt,msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}else{
			folders:=tools.GetMailNum(mailServer.Server, mailServer.Email, mailServer.Password)
			log.Println(folders)
			log.Println(oldFolders)
			for name,num:=range folders{
				if oldFolders[name]!=num{
					msg, _ = json.Marshal(tools.JsonResult{Code: 200, Msg: fmt.Sprintf("%s:%d",name,num)})
					c.WriteMessage(mt,msg)
				}
			}
			oldFolders=folders
		}
	}
}
