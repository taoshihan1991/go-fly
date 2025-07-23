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
(1, 'agent', 'b33aed8f3134996703dc39f9a7c95783', 'Open Source LiveChat Support', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg');

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
 `user_id` varchar(500) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 UNIQUE KEY `conf_key` (`conf_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Announcement', 'AllNotice', 'Open source customer support system at your service','agent');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Offline Message', 'OfflineMessage', 'I am currently offline and will reply to you later!','agent');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Welcome Message', 'WelcomeMessage', 'How may I help you?','agent');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Email Address (SMTP)', 'NoticeEmailSmtp', '','agent');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Email Account', 'NoticeEmailAddress', '','agent');
INSERT INTO `config` (`id`, `conf_name`, `conf_key`, `conf_value`, `user_id`) VALUES
(NULL, 'Email Password (SMTP)', 'NoticeEmailPassword', '','agent');


DROP TABLE IF EXISTS `reply_group`;
CREATE TABLE `reply_group` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `group_name` varchar(50) NOT NULL DEFAULT '',
 `user_id` varchar(50) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`),
 KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `reply_group` (`id`, `group_name`, `user_id`) VALUES
(NULL, 'Frequently Asked Questions', 'agent');

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