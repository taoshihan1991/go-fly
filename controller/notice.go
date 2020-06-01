package controller
import(
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{}
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
		err = c.WriteMessage(mt, []byte("{\"title\":\"测试\"}"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
