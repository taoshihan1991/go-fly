DROP TABLE IF EXISTS `user`|
CREATE TABLE `user` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `name` varchar(50) NOT NULL DEFAULT '',
 `password` varchar(50) NOT NULL DEFAULT '',
 `nickname` varchar(50) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `avator` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
TRUNCATE TABLE `user`|
INSERT INTO `user` (`id`, `name`, `password`, `nickname`, `created_at`, `updated_at`, `deleted_at`, `avator`) VALUES
(1, 'kefu2', '202cb962ac59075b964b07152d234b70', '小白菜', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg'),
(2, 'kefu3', '202cb962ac59075b964b07152d234b70', '中白菜', '2020-07-02 14:36:46', '2020-07-05 08:46:57', NULL, '/static/images/11.jpg')|

DROP TABLE IF EXISTS `visitor`|
CREATE TABLE `visitor` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `name` varchar(50) NOT NULL DEFAULT '',
 `avator` varchar(500) NOT NULL DEFAULT '',
 `source_ip` varchar(50) NOT NULL DEFAULT '',
 `to_id` varchar(50) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `visitor_id` varchar(100) NOT NULL DEFAULT '',
 `status` tinyint(4) NOT NULL DEFAULT '0',
 `refer` varchar(500) NOT NULL DEFAULT '',
 `city` varchar(100) NOT NULL DEFAULT '',
 `client_ip` varchar(100) NOT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `visitor_id` (`visitor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|

DROP TABLE IF EXISTS `message`|
CREATE TABLE `message` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `kefu_id` varchar(100) NOT NULL DEFAULT '',
 `visitor_id` varchar(100) NOT NULL DEFAULT '',
 `content` varchar(2048) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `mes_type` enum('kefu','visitor') NOT NULL DEFAULT 'visitor',
 `status` enum('read','unread') NOT NULL DEFAULT 'unread',
 PRIMARY KEY (`id`),
 KEY `kefu_id` (`kefu_id`),
 KEY `visitor_id` (`visitor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|

DROP TABLE IF EXISTS `user_role`|
CREATE TABLE `user_role` (
 `id` int(11) NOT  NULL AUTO_INCREMENT,
 `user_id` int(11) NOT NULL DEFAULT '0',
 `role_id` int(11) NOT NULL DEFAULT '0',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8|
INSERT INTO `user_role` (`id`, `user_id`, `role_id`) VALUES
(1, 1, 1),
(2, 2, 2)|

DROP TABLE IF EXISTS `role`|
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `method` varchar(100) NOT NULL DEFAULT '',
  `path` varchar(2048) NOT NULL DEFAULT '',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
INSERT INTO `role` (`id`, `name`, `method`, `path`) VALUES
(1, '普通客服', 'GET', '/kefuinfo,/kefulist,/roles'),
(2, '管理员', '*', '*')|

DROP TABLE IF EXISTS `welcome`|
CREATE TABLE `welcome` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` varchar(100) NOT NULL DEFAULT '',
 `content` varchar(500) NOT NULL DEFAULT '',
 `is_default` tinyint(3) unsigned NOT NULL DEFAULT '0',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`) VALUES
(1, 'kefu2', '本客服代码开源,欢迎star,开源地址:https://github.com/taoshihan1991/go-fly', 1, '2020-08-24 02:57:49')|

DROP TABLE IF EXISTS `ipblack`|
CREATE TABLE `ipblack` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `ip` varchar(100) NOT NULL DEFAULT '',
 `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `kefu_id` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|

DROP TABLE IF EXISTS `config`|
CREATE TABLE `config` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `conf_name` varchar(255) NOT NULL DEFAULT '',
 `conf_key` varchar(255) NOT NULL DEFAULT '',
 `conf_value` varchar(255) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `conf_key` (`conf_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '是否开启Server酱微信提醒', 'NoticeServerJiang', 'false');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'Server酱API', 'ServerJiangAPI', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '微信小程序Token', 'WeixinToken', '');

DROP TABLE IF EXISTS `about`|
CREATE TABLE `about` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title_cn` varchar(255) NOT NULL DEFAULT '',
  `title_en` varchar(255) NOT NULL DEFAULT '',
  `keywords_cn` varchar(255) NOT NULL DEFAULT '',
  `keywords_en` varchar(255) NOT NULL DEFAULT '',
  `desc_cn` varchar(1024) NOT NULL DEFAULT '',
  `desc_en` varchar(1024) NOT NULL DEFAULT '',
  `css_js` text NOT NULL,
  `html_cn` text NOT NULL,
  `html_en` text NOT NULL,
  `page` varchar(50) NOT NULL DEFAULT '',
PRIMARY KEY (`id`),
UNIQUE KEY `page` (`page`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
INSERT INTO `about` (`id`, `title_cn`, `title_en`, `keywords_cn`, `keywords_en`, `desc_cn`, `desc_en`, `css_js`, `html_cn`, `html_en`, `page`) VALUES
(1, 'GO-FLY - 极简强大的Go语言在线客服系统', 'GO-FLY - Simple and Powerful Go language online customer chat system', 'GO-FLY，GO-FLY', 'GO-FLY，GO-FLY', 'GO-FLY，GO-FLY，一套为PHP工程师、Golang工程师准备的基于 Vue 2.0的在线客服即时通讯系统', 'GO-FLY，GO-FLY, a Vue 2.0-based online customer service instant messaging system for PHP engineers and Golang engineers', '    <style>\r\n        *{\r\n            margin: 0;padding: 0;\r\n        }\r\n        .header{\r\n            height: 80px;\r\n            background-color: #fff;\r\n            color: #fff;\r\n            top: 0;\r\n            left: 0;\r\n            width: 100%;\r\n            line-height: 80px;\r\n            z-index: 100;\r\n            position: relative;\r\n        }\r\n        .container{\r\n            width: 1140px;\r\n            padding: 0;\r\n            margin: 0 auto;\r\n        }\r\n        .header .container{\r\n            height: 100%;\r\n            box-sizing: border-box;\r\n            border-bottom: 1px solid #dcdfe6;\r\n        }\r\n        .header h1{\r\n            margin: 0;\r\n            float: left;\r\n            font-size: 32px;\r\n            font-weight: 400;\r\n        }\r\n        .header a{\r\n            color: #519eff;\r\n            font-family: \"Microsoft JhengHei\";\r\n            text-decoration: none;\r\n        }\r\n        .header h1 a{\r\n            font-size: 30px;\r\n            font-weight: bold;\r\n        }\r\n        .header .navBtn{\r\n            float: right;\r\n            margin-left: 20px;\r\n        }\r\n        .banner{\r\n            padding-top: 20px;\r\n            text-align: center;\r\n        }\r\n        .banner h1{\r\n            font-size: 34px;\r\n            margin: 0;\r\n            line-height: 48px;\r\n            color: #555;\r\n            font-weight: 500;\r\n            font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,SimSun,sans-serif;\r\n        }\r\n        .banner p{\r\n            font-size: 18px;\r\n            line-height: 28px;\r\n            color: #888;\r\n            margin: 10px 0 5px;\r\n        }\r\n        .jumbotron{\r\n            width: 587px;\r\n            height: 560px;\r\n            margin: 30px auto;\r\n        }\r\n        .footer {\r\n            clear: both;\r\n            background-color: #f7fbfd;\r\n            width: 100%;\r\n            padding: 40px 150px;\r\n            box-sizing: border-box;\r\n        }\r\n        .copyright{\r\n            color: #6c757d;\r\n            text-align: center;\r\n            margin: 60px 0;\r\n        }\r\n        .mainTechLeft{\r\n            width: 300px;\r\n            float: left;\r\n        }\r\n        .mainTechLeft h1{\r\n            font-size: 34px;\r\n            margin: 0;\r\n            line-height: 48px;\r\n            color: #555;\r\n            font-weight: 500;\r\n            font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,SimSun,sans-serif;\r\n        }\r\n        .mainTechLeft p{\r\n            font-size: 18px;\r\n            line-height: 28px;\r\n            color: #888;\r\n            margin: 10px 0 5px;\r\n        }\r\n        .floatRight{\r\n            width: 700px;\r\n            border: 1px solid #e1e1e1;\r\n            padding: 4px;\r\n            margin-top: 35px;\r\n            display: block;\r\n            float: right;\r\n        }\r\n    </style>', '<div class=\"banner\">\r\n    <h1>极简强大的Go语言在线客服系统</h1>\r\n    <p>GO-FLY，一套为PHP工程师、Golang工程师准备的基于 Vue 2.0的在线客服即时通讯系统</p>\r\n</div>\r\n<div class=\"jumbotron\">\r\n    \r\n    <img src=\"/static/images/intro1.jpg\"/>\r\n    \r\n</div>\r\n<div class=\"container\">\r\n    <div class=\"mainTechLeft\">\r\n        <h1>主要技术架构</h1>\r\n        <p>github.com/dgrijalva/jwt-go</p>\r\n        <p>github.com/gin-gonic/gin</p>\r\n        <p> github.com/go-sql-driver/mysql</p>\r\n        <p>  github.com/gobuffalo/packr/v2</p>\r\n        <p>  github.com/gorilla/websocket</p>\r\n        <p>   github.com/ipipdotnet/ipdb-go</p>\r\n        <p>    github.com/jinzhu/gorm</p>\r\n        <p>    github.com/satori/go.uuid</p>\r\n        <p>   github.com/spf13/cobra</p>\r\n        <p>   github.com/swaggo/gin-swagger</p>\r\n        <p>  github.com/swaggo/swag\r\n        </p>\r\n    </div>\r\n    <img src=\"/static/images/admin.png\" class=\"floatRight\"/>\r\n</div>', '<div class=\"banner\">     <h1>Simple and Powerful Go language online customer chat system</h1>     <p>GO-FLY, a Vue 2.0-based online customer service instant messaging system for PHP engineers and Golang engineers</p> </div> <div class=\"jumbotron\">          <img src=\"/static/images/intro3.png\"/>      </div> <div class=\"container\">     <div class=\"mainTechLeft\">         <h1>Main technical architecture</h1>         <p>github.com/dgrijalva/jwt-go</p>         <p>github.com/gin-gonic/gin</p>         <p> github.com/go-sql-driver/mysql</p>         <p>  github.com/gobuffalo/packr/v2</p>         <p>  github.com/gorilla/websocket</p>         <p>   github.com/ipipdotnet/ipdb-go</p>         <p>    github.com/jinzhu/gorm</p>         <p>    github.com/satori/go.uuid</p>         <p>   github.com/spf13/cobra</p>         <p>   github.com/swaggo/gin-swagger</p>         <p>  github.com/swaggo/swag         </p>     </div>     <img src=\"/static/images/admin.png\" class=\"floatRight\"/> </div>','index')|
