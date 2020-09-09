DROP TABLE IF EXISTS `user`;
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
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `visitor_id` varchar(100) NOT NULL DEFAULT '',
 `status` tinyint(4) NOT NULL DEFAULT '0',
 `refer` varchar(500) NOT NULL DEFAULT '',
 `city` varchar(100) NOT NULL DEFAULT '',
 `client_ip` varchar(100) NOT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `visitor_id` (`visitor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `message`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
 `id` int(11) NOT  NULL AUTO_INCREMENT,
 `user_id` int(11) NOT NULL DEFAULT '0',
 `role_id` int(11) NOT NULL DEFAULT '0',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
INSERT INTO `user_role` (`id`, `user_id`, `role_id`) VALUES
(1, 1, 1),
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
(1, '普通客服', 'GET', '/kefuinfo,/kefulist,/roles'),
(2, '管理员', '*', '*');

DROP TABLE IF EXISTS `welcome`;
CREATE TABLE `welcome` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` varchar(100) NOT NULL DEFAULT '',
 `content` varchar(500) NOT NULL DEFAULT '',
 `is_default` tinyint(3) unsigned NOT NULL DEFAULT '0',
 `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `welcome` (`id`, `user_id`, `content`, `is_default`, `ctime`) VALUES
(1, 'kefu2', '本客服代码开源,欢迎star,开源地址:https://github.com/taoshihan1991/go-fly', 1, '2020-08-24 02:57:49');

DROP TABLE IF EXISTS `ipblack`;
CREATE TABLE `ipblack` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `ip` varchar(100) NOT NULL DEFAULT '',
 `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `kefu_id` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;