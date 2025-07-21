## GOFLY LIVE CHAT
An open-source customer service system developed in Golang

### Project Overview  
- Backend: `gin`, `jwt-go`, `websocket`, `go.uuid`, `gorm`, `cobra`  
- Frontend: `VueJS`, `ElementUI`  
- Database: `MySQL`  

---

### Installation & Usage  

#### 1. Set Up MySQL Database  
- Install and run MySQL (version ≥ 5.5).  
- Create a database:  
  ```sql
  CREATE DATABASE gofly CHARSET utf8mb4;
   
*  Configure Database Connection
   Edit mysql.json in the config directory:
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"goflychat",
	"Username":"goflychat",
	"Password":"goflychat"
}
```
* Install and Configure Golang
  Run the following commands:
```php
wget https://studygolang.com/dl/golang/go1.20.2.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.20.2.linux-amd64.tar.gz
mv go1.20.2.linux-amd64.tar.gz /tmp
echo "PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
source /etc/profile
go version
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
* Download the Source Code

  Clone the repository in any directory:
```php
git clone https://github.com/taoshihan1991/go-fly.git 
cd go-fly  
 ```  
* Initialize the Database
 ```php
 go run go-fly.go install
 ```  
* Run the Application
```php
 go run go-fly.go server
 ```
* ​​Build executable
```php
 go build -o kefu
```
* ​​Run binary​​:
```php
  Linux: ./kefu server (optional flags: -p 8082 -d)
  
  Windows: kefu.exe server (optional flags: -p 8082 -d)
```  
* Terminate the Process
```php
   killall kefu
``` 

Once running, the service listens on port 8081. Access via http://[your-ip]:8081.

For domain access, configure a reverse proxy to port 8081 to hide the port number.
### Customer Service Integration
Chat Link

http://127.0.0.1:8081/chatIndex?kefu_id=kefu2

Popup Integration

```
    (function(a, b, c, d) {
        let h = b.getElementsByTagName('head')[0];let s = b.createElement('script');
        s.type = 'text/javascript';s.src = c+"/static/js/kefu-front.js";s.onload = s.onreadystatechange = function () {
            if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete") d(c);
        };h.appendChild(s);
    })(window, document,"http://127.0.0.1:8081",function(u){
        KEFU.init({
            KEFU_URL:u,
            KEFU_KEFU_ID: "kefu2",
        })
    });

```
### Important Notice  
The use of this project for illegal or non-compliant purposes, including but not limited to viruses, trojans, pornography, gambling, fraud, prohibited items, counterfeit products, false information, cryptocurrencies, and financial violations, is strictly prohibited.  

This project is intended solely for personal learning and testing purposes. Any commercial use or illegal activities are explicitly forbidden!!!  



### Copyright Notice
This project provides full-featured code but is intended ​​only for personal demonstration and testing​​. Commercial use is strictly prohibited.

By using this software, you agree to comply with all applicable local laws and regulations. ​​You are solely responsible for any legal consequences arising from misuse.​