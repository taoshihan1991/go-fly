DROP TABLE IF EXISTS `user`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
TRUNCATE TABLE `user`;
INSERT INTO `user` (`id`, `name`, `password`, `nickname`, `created_at`, `updated_at`, `deleted_at`, `avator`) VALUES
(1, 'kefu2', '202cb962ac59075b964b07152d234b70', '小白菜', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg'),
(2, 'kefu3', '202cb962ac59075b964b07152d234b70', '中白菜', '2020-07-02 14:36:46', '2020-07-05 08:46:57', NULL, '/static/images/11.jpg');

DROP TABLE IF EXISTS `visitor`;
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
 KEY `to_id` (`to_id`),
 KEY `idx_update` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `message`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
 `id` int(11) NOT  NULL AUTO_INCREMENT,
 `user_id` int(11) NOT NULL DEFAULT '0',
 `role_id` int(11) NOT NULL DEFAULT '0',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
INSERT INTO `user_role` (`id`, `user_id`, `role_id`) VALUES
(1, 1, 2),
(2, 2, 2);

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `method` varchar(100) NOT NULL DEFAULT '',
  `path` varchar(2048) NOT NULL DEFAULT '',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `role` (`id`, `name`, `method`, `path`) VALUES
(1, '普通客服', 'GET', 'GET:/kefuinfo,GET:/kefulist,GET:/roles,POST:/notice_save,POST:/notice'),
(2, '管理员', '*', '*');

DROP TABLE IF EXISTS `welcome`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`, `keyword`) VALUES
(NULL, 'kefu2', '我暂时离线，留言已转发到我的邮箱，稍后回复~', 1, '2020-08-24 02:57:49','offline');
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`, `keyword`) VALUES
(NULL, 'kefu2', '本客服代码开源,欢迎star,开源地址:https://github.com/taoshihan1991/go-fly', 0, '2020-08-24 02:57:49','welcome');

DROP TABLE IF EXISTS `ipblack`;
CREATE TABLE `ipblack` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `ip` varchar(100) NOT NULL DEFAULT '',
 `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `kefu_id` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `conf_name` varchar(255) NOT NULL DEFAULT '',
 `conf_key` varchar(255) NOT NULL DEFAULT '',
 `conf_value` varchar(255) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `conf_key` (`conf_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '是否开启Server酱微信提醒', 'NoticeServerJiang', 'false');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'Server酱API', 'ServerJiangAPI', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '微信小程序Token', 'WeixinToken', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '当前小程序审核状态', 'MiniAppAudit', 'yes');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '是否允许上传附件', 'SendAttachment', 'true');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(SMTP地址)', 'NoticeEmailSmtp', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(邮箱)', 'NoticeEmailAddress', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '发送通知邮件(密码)', 'NoticeEmailPassword', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(Token)', 'GetuiToken', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppID)', 'GetuiAppID', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppKey)', 'GetuiAppKey', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppSecret)', 'GetuiAppSecret', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, 'App个推(AppMasterSecret)', 'GetuiMasterSecret', '');
DROP TABLE IF EXISTS `about`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `about` (`id`, `title_cn`, `title_en`, `keywords_cn`, `keywords_en`, `desc_cn`, `desc_en`, `css_js`, `html_cn`, `html_en`, `page`) VALUES
(NULL, '免费开源客服系统GOFLY0.4.1-演示页',
'Free Customer Live Chat GOFLY0.4.1-demo',
'GOFLY，GO-FLY',
'GOFLY，GO-FLY',
'一款开箱即用的在线客服系统',
'a free customer live chat',
'<style>body{color: #333;padding-left: 40px;}h1{font-size: 6em;}h2{font-size: 3em;font-weight: normal;}a{color: #333;}</style>',
'<script src="/assets/js/gofly-front.js?v=1"></script><script>
    GOFLY.init({
        GOFLY_URL:"",
        GOFLY_KEFU_ID: "kefu2",
        GOFLY_BTN_TEXT: "GOFLY 在线客服!",
        GOFLY_LANG:"cn",
    })
</script>
 <h1>:)</h1><h2>你好 <a href="https://gofly.sopans.com">GOFLY0.4.1</a> 在线客服系统 !</h2><h3><a href="/login">Administrator</a>&nbsp;<a href="/index_en">English</a>&nbsp;<a href="/index_cn">中文</a></h3>',
'<script src="/assets/js/gofly-front.js?v=1"></script><script>
    GOFLY.init({
        GOFLY_URL:"",
        GOFLY_KEFU_ID: "kefu2",
        GOFLY_BTN_TEXT: "GOFLY LIVE CHAT!",
        GOFLY_LANG:"en",
    })
</script>
 <h1>:)</h1><h2>HELLO <a href="https://gofly.sopans.com">GOFLY0.4.1</a> LIVE CHAT !</h2><h3><a href="/login">Administrator</a>&nbsp;<a href="/index_en">English</a>&nbsp;<a href="/index_cn">中文</a></h3>',
 'index');
DROP TABLE IF EXISTS `reply_group`;
CREATE TABLE `reply_group` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `group_name` varchar(50) NOT NULL DEFAULT '',
 `user_id` varchar(50) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `reply_group` (`id`, `group_name`, `user_id`) VALUES (NULL, '常见问题', 'kefu2');
DROP TABLE IF EXISTS `reply_item`;
CREATE TABLE `reply_item` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `content` varchar(1024) NOT NULL DEFAULT '',
 `group_id` int(11) NOT NULL DEFAULT '0',
 `user_id` varchar(50) NOT NULL DEFAULT '',
 `item_name` varchar(50) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`),
 KEY `group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `reply_item` (`id`, `content`, `group_id`, `user_id`, `item_name`) VALUES (NULL, '在这里[官网]link[https://gofly.sopans.com]!','1','kefu2', '官方地址在哪?');
DROP TABLE IF EXISTS `land_page`;
CREATE TABLE `land_page` (
  `id` int(11) NOT NULL,
  `title` varchar(125) NOT NULL DEFAULT '',
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `language` varchar(50) NOT NULL DEFAULT '',
  `page_id` varchar(50) NOT NULL DEFAULT '',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
DROP TABLE IF EXISTS `language`;
CREATE TABLE `language` (
  `id` int(11) NOT NULL,
  `country` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `short_key` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (1, '中文简体', 'zh-cn');
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (2, '正體中文', 'zh-tw');
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (3, 'English', 'en_us');
INSERT INTO `language` (`id`, `country`, `short_key`) VALUES (4, '日本語', 'ja_jp');
DROP TABLE IF EXISTS `user_client`;
CREATE TABLE `user_client` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `kefu` varchar(100) NOT NULL DEFAULT '',
 `client_id` varchar(100) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 UNIQUE KEY `idx_user` (`kefu`,`client_id`),
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
