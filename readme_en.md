# <b>GOFLY</b> [V1KF] GOFLY LIVE CHAT FOR CUSTOMER SUPPORT SERVICE
<a href="readme.md">中文</a> |
<a href="readme_en.md">English</a> |
<a href="https://gofly.v1kf.com">The official website</a>

### Please note that this project is for personal learning and testing only, and is prohibited for all online commercial use and illegal use!

### It appears that you are providing information about purchasing a paid version of the software and receiving an installation package and authorization. 

## The main technology stack
gin + jwt-go + websocket + go.uuid + gorm + cobra + VueJS + ElementUI + MySQL

### Preview

![Image text](https://img2022.cnblogs.com/blog/726254/202211/726254-20221108002459990-32759129.png)

![Image text](https://img2022.cnblogs.com/blog/726254/202211/726254-20221108002516168-1465488645.png)

![Image text](https://img2022.cnblogs.com/blog/726254/202211/726254-20221108002619691-1817390882.png)



### To install and use:


#### 1. Install and run MySQL >=5.5, and create the gofly database.
 
    create database gofly charset utf8;
   
   edit config/mysql.json
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"gofly",
	"Username":"go-fly",
	"Password":"go-fly"
}
```
        
#### 2. Run the source code:

1. Go module:

   go env -w GO111MODULE=on
   
   go env -w GOPROXY=https://goproxy.cn,direct
   
   git clone https://github.com/taoshihan1991/go-fly.git
   
   go run go-fly.go install
   
   go run go-fly.go server

3. Source code packaging: go build go-fly.go, which will generate the go-fly executable file.

4. Import the database (will delete the table and clear the data): ./go-fly install

5. Binary file execution:
 
   linux:   ./go-fly server [optional  -p 8082 -d]
   
   windows: go-fly.exe server [optional  -p 8082 -d]

6. Close the program:
   ./go-fly stop  

    For Linux, use the ps and kill commands to kill the process:
    
    ps -ef|grep go-fly
    
    kill process parent process id; kill process child process id
   
    or  killall go-fly

#### Usage
The server installation is complete and the service is running, and the client can be accessed through the browser.

The default port is 8081. If you use the -p parameter to specify the port, you can access it through the browser http://127.0.0.1:port.

The default user name and password are kefu2 and 123
   


   
### Nginx

```php
server {
       listen 443 ssl http2;
        ssl on;
        ssl_certificate   conf.d/cert/4263285_gofly.sopans.com.pem;
        ssl_certificate_key  conf.d/cert/4263285_gofly.sopans.com.key;
        ssl_session_timeout 5m;
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
        ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
        ssl_prefer_server_ciphers on;
        #listen          80; 
        server_name  gofly.sopans.com;
        access_log  /var/log/nginx/gofly.sopans.com.access.log  main;
        location /static {
                root /var/www/html/go-fly;//自己的部署路径,静态文件直接nginx响应
        }
        location / {
                proxy_pass http://127.0.0.1:8081;
                    proxy_http_version 1.1;
                    proxy_set_header X-Real-IP $remote_addr;
                    proxy_set_header Upgrade $http_upgrade;
                    proxy_set_header Connection "upgrade";
                    proxy_set_header Origin "";
        }
}
server{
       listen 80;
        server_name  gofly.sopans.com;
        access_log  /var/log/nginx/gofly.sopans.com.access.log  main;
        location /static {
                root /var/www/html/go-fly;//自己的部署路径,静态文件直接nginx响应
        }        
        location / {
                proxy_pass http://127.0.0.1:8081;
                    proxy_http_version 1.1;
                    proxy_set_header X-Real-IP $remote_addr;
                    proxy_set_header Upgrade $http_upgrade;
                    proxy_set_header Connection "upgrade";
                    proxy_set_header Origin "";
        }
}
```

#### If you encounter a pop-up window requiring authentication after logging into the backend, please go to the official website to register an account and bind your phone. Entering your bound phone number + your own password will allow you to pass the authentication. If you do not have a Chinese phone number, please contact me on the official website to obtain a test phone number and password.

### Copyright 

This project is a complete code with full functionality, but it is still only for personal demonstration and testing and does not include online use. 

All commercial activities are prohibited. When using this software, please comply with local laws and regulations. Any illegal use is at your own risk.