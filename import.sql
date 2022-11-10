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
INSERT INTO `user` (`id`, `name`, `password`, `nickname`, `created_at`, `updated_at`, `deleted_at`, `avator`) VALUE
(1, 'kefu2', '202cb962ac59075b964b07152d234b70', '智能客服系统', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg');

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
INSERT INTO `user_role` (`id`, `user_id`, `role_id`) VALUE
(1, 1, 1);

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `method` varchar(100) NOT NULL DEFAULT '',
  `path` varchar(2048) NOT NULL DEFAULT '',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `role` (`id`, `name`, `method`, `path`) VALUES
(1, '普通客服', '*', '*');

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
(NULL, 'kefu2', '请问有什么可以帮您？', 0, '2020-08-24 02:57:49','welcome');

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
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '公告信息', 'AllNotice', '开源智能客服系统为您服务');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '离线消息', 'OfflineMessage', '我现在离线，稍后回复您！');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '欢迎消息', 'WelcomeMessage', '有什么可以帮您？');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '邮箱地址(SMTP地址)', 'NoticeEmailSmtp', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '邮箱账户', 'NoticeEmailAddress', '');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`) VALUES (NULL, '邮箱密码(SMTP密码)', 'NoticeEmailPassword', '');

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