DROP TABLE IF EXISTS `user`|
CREATE TABLE `user` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `name` varchar(50) NOT NULL DEFAULT '',
 `password` varchar(50) NOT NULL DEFAULT '',
 `nickname` varchar(50) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `avator` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_name` (`name`)
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
 `updated_at` timestamp NULL DEFAULT NULL,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `visitor_id` varchar(100) NOT NULL DEFAULT '',
 `status` tinyint(4) NOT NULL DEFAULT '0',
 `refer` varchar(500) NOT NULL DEFAULT '',
 `city` varchar(100) NOT NULL DEFAULT '',
 `client_ip` varchar(100) NOT NULL DEFAULT '',
 `extra` varchar(2048) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `visitor_id` (`visitor_id`),
 KEY `to_id` (`to_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|

DROP TABLE IF EXISTS `message`|
CREATE TABLE `message` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `kefu_id` varchar(100) NOT NULL DEFAULT '',
 `visitor_id` varchar(100) NOT NULL DEFAULT '',
 `content` varchar(2048) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NULL DEFAULT NULL,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `mes_type` enum('kefu','visitor') NOT NULL DEFAULT 'visitor',
 `status` enum('read','unread') NOT NULL DEFAULT 'unread',
 PRIMARY KEY (`id`),
 KEY `kefu_id` (`kefu_id`),
 KEY `visitor_id` (`visitor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4|

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
(1, '普通客服', 'GET', 'GET:/kefuinfo,GET:/kefulist,GET:/roles,POST:/notice_save,POST:/notice'),
(2, '管理员', '*', '*')|

DROP TABLE IF EXISTS `welcome`|
CREATE TABLE `welcome` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` varchar(100) NOT NULL DEFAULT '',
 `keyword` varchar(100) NOT NULL DEFAULT '',
 `content` varchar(500) NOT NULL DEFAULT '',
 `is_default` tinyint(3) unsigned NOT NULL DEFAULT '0',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`),
 KEY `keyword` (`keyword`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`, `keyword`) VALUES
(NULL, 'kefu2', '我暂时离线，留言已转发到我的邮箱，稍后回复~', 1, '2020-08-24 02:57:49','offline')|
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`, `keyword`) VALUES
(NULL, 'kefu2', '本客服代码开源,欢迎star,开源地址:https://github.com/taoshihan1991/go-fly', 0, '2020-08-24 02:57:49','welcome')|

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
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '是否开启Server酱微信提醒', 'NoticeServerJiang', 'false')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'Server酱API', 'ServerJiangAPI', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '微信小程序Token', 'WeixinToken', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '当前小程序审核状态', 'MiniAppAudit', 'yes')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '是否允许上传附件', 'SendAttachment', 'true')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(SMTP地址)', 'NoticeEmailSmtp', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(邮箱)', 'NoticeEmailAddress', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(密码)', 'NoticeEmailPassword', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(Token)', 'GetuiToken', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppID)', 'GetuiAppID', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppKey)', 'GetuiAppKey', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppSecret)', 'GetuiAppSecret', '')|
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppMasterSecret)', 'GetuiMasterSecret', '')|
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
(NULL, '免费开源客服系统GOFLY-演示页', 'Simple and Powerful Golang customer live chat-GOFLY-demo', 'GO-FLY，GO-FLY', 'GO-FLY，GO-FLY', 'GO-FLY，GO-FLY，一套为PHP工程师、Golang工程师准备的基于 Vue 2.0的在线客服即时通讯系统', 'GO-FLY，GO-FLY, a Vue 2.0-based online customer service instant messaging system for PHP engineers and Golang engineers', '    <style>\n        *{\n            margin: 0;padding: 0;\n        }\n        .header{\n            height: 80px;\n            background-color: #fff;\n            color: #fff;\n            top: 0;\n            left: 0;\n            width: 100%;\n            line-height: 80px;\n            z-index: 100;\n            position: relative;\n        }\n        .container{\n            width: 1140px;\n            padding: 0;\n            margin: 0 auto;\n        }\n        .header .container{\n            height: 100%;\n            box-sizing: border-box;\n            border-bottom: 1px solid #dcdfe6;\n        }\n        .header h1{\n            margin: 0;\n            float: left;\n            font-size: 32px;\n            font-weight: 400;\n        }\n        .header a{\n            color: #519eff;\n            font-family: \"Microsoft JhengHei\";\n            text-decoration: none;\n        }\n        .header h1 a{\n            font-size: 30px;\n            font-weight: bold;\n        }\n        .header .navBtn{\n            float: right;\n            margin-left: 20px;\n        }\n        .banner{\n            padding-top: 20px;\n            text-align: center;\n        }\n        .banner h1{\n            font-size: 34px;\n            margin: 0;\n            line-height: 48px;\n            color: #555;\n            font-weight: 500;\n            font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,SimSun,sans-serif;\n        }\n        .banner p{\n            font-size: 18px;\n            line-height: 28px;\n            color: #888;\n            margin: 10px 0 5px;\n        }\n        .jumbotron{\n            width: 587px;\n            height: 560px;\n            margin: 30px auto;\n        }\n        .footer {\n            clear: both;\n            background-color: #f7fbfd;\n            width: 100%;\n            padding: 40px 150px;\n            box-sizing: border-box;\n        }\n        .copyright{\n            color: #6c757d;\n            text-align: center;\n            margin: 60px 0;\n        }\n        .mainTechLeft{\n            width: 300px;\n            float: left;\n        }\n        .mainTechLeft h1{\n            font-size: 34px;\n            margin: 0;\n            line-height: 48px;\n            color: #555;\n            font-weight: 500;\n            font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,SimSun,sans-serif;\n        }\n        .mainTechLeft p{\n            font-size: 18px;\n            line-height: 28px;\n            color: #888;\n            margin: 10px 0 5px;\n        }\n        .floatRight{\n            width: 700px;\n            border: 1px solid #e1e1e1;\n            padding: 4px;\n            margin-top: 35px;\n            display: block;\n            float: right;\n        }\n    </style>', '<header class=\"header\">\n    <div class=\"container\">\n        <h1><a href=\"/\">GO-FLY</a></h1>\n        <a class=\"navBtn\" href=\"/index_en\">English (United States)</a>\n        <a class=\"navBtn\" href=\"/index_cn\">中文版 (简体)</a>\n        <a class=\"navBtn\" href=\"https://github.com/taoshihan1991/go-fly\" target=\"_blank\">Github</a>\n        <a class=\"navBtn\" href=\"/login\">客服入口</a>\n        <a class=\"navBtn\" href=\"/docs/index.html\" target=\"_blank\">接口文档</a>\n    </div>\n</header>\n<div class=\"banner\">\n    <h1>极简强大的Go语言在线客服系统</h1>\n    <p>GO-FLY，一套为PHP工程师、Golang工程师准备的基于 Vue 2.0的在线客服即时通讯系统</p>\n</div>\n<div class=\"jumbotron\">\n    \n    <img src=\"/static/images/intro1.jpg\"/>\n    \n</div>\n<div class=\"container\">\n    <div class=\"mainTechLeft\">\n        <h1>主要技术架构</h1>\n        <p>github.com/dgrijalva/jwt-go</p>\n        <p>github.com/gin-gonic/gin</p>\n        <p> github.com/go-sql-driver/mysql</p>\n        <p>  github.com/gobuffalo/packr/v2</p>\n        <p>  github.com/gorilla/websocket</p>\n        <p>   github.com/ipipdotnet/ipdb-go</p>\n        <p>    github.com/jinzhu/gorm</p>\n        <p>    github.com/satori/go.uuid</p>\n        <p>   github.com/spf13/cobra</p>\n        <p>   github.com/swaggo/gin-swagger</p>\n        <p>  github.com/swaggo/swag\n        </p>\n    </div>\n    <img src=\"/static/images/admin.png\" class=\"floatRight\"/>\n</div>\n\n\n<footer class=\"footer\">\n    <div class=\"container\">\n\n    </div>\n    <div class=\"copyright\">\n        陶士涵的菜地版权所有&copy; 2020\n    </div>\n</footer>', '<header class=\"header\">\n    <div class=\"container\">\n        <h1><a href=\"/\">GO-FLY</a></h1>\n        <a class=\"navBtn\" href=\"/index_en\">English (United States)</a>\n        <a class=\"navBtn\" href=\"/index_cn\">中文版 (简体)</a>\n        <a class=\"navBtn\" href=\"https://github.com/taoshihan1991/go-fly\" target=\"_blank\">Github</a>\n        <a class=\"navBtn\" href=\"/login\">Agents Here</a>\n        <a class=\"navBtn\" href=\"/docs/index.html\" target=\"_blank\">API Documents</a>\n    </div>\n</header>\n<div class=\"banner\">     <h1>Simple and Powerful Go language online customer chat system</h1>     <p>GO-FLY, a Vue 2.0-based online customer service instant messaging system for PHP engineers and Golang engineers</p> </div> <div class=\"jumbotron\">          <img src=\"/static/images/intro3.png\"/>      </div> <div class=\"container\">     <div class=\"mainTechLeft\">         <h1>Main technical architecture</h1>         <p>github.com/dgrijalva/jwt-go</p>         <p>github.com/gin-gonic/gin</p>         <p> github.com/go-sql-driver/mysql</p>         <p>  github.com/gobuffalo/packr/v2</p>         <p>  github.com/gorilla/websocket</p>         <p>   github.com/ipipdotnet/ipdb-go</p>         <p>    github.com/jinzhu/gorm</p>         <p>    github.com/satori/go.uuid</p>         <p>   github.com/spf13/cobra</p>         <p>   github.com/swaggo/gin-swagger</p>         <p>  github.com/swaggo/swag         </p>     </div>     <img src=\"/static/images/admin.png\" class=\"floatRight\"/> </div>\n\n\n<footer class=\"footer\">\n    <div class=\"container\">\n\n    </div>\n    <div class=\"copyright\">\n        TaoShihan&copy; 2020\n    </div>\n</footer>', 'index')|
DROP TABLE IF EXISTS `reply_group`|
CREATE TABLE `reply_group` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `group_name` varchar(50) NOT NULL DEFAULT '',
 `user_id` varchar(50) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
DROP TABLE IF EXISTS `reply_item`|
CREATE TABLE `reply_item` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `content` varchar(1024) NOT NULL DEFAULT '',
 `group_id` int(11) NOT NULL DEFAULT '0',
 `user_id` varchar(50) NOT NULL DEFAULT '',
 `item_name` varchar(50) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`),
 KEY `group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|
DROP TABLE IF EXISTS `land_page`|
CREATE TABLE `land_page` (
  `id` int(11) NOT NULL,
  `title` varchar(125) NOT NULL DEFAULT '',
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `language` varchar(50) NOT NULL DEFAULT '',
  `page_id` varchar(50) NOT NULL DEFAULT '',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci|
DROP TABLE IF EXISTS `language`|
CREATE TABLE `language` (
  `id` int(11) NOT NULL,
  `country` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `short_key` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci|
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (1, '中文简体', 'zh-cn')|
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (2, '正體中文', 'zh-tw')|
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (3, 'English', 'en_us')|
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (4, '日本語', 'ja_jp')|
DROP TABLE IF EXISTS `user_client`|
CREATE TABLE `user_client` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `kefu` varchar(100) NOT NULL DEFAULT '',
 `client_id` varchar(100) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 UNIQUE KEY `idx_user` (`kefu`,`client_id`),
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8|