CREATE TABLE `user` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `name` varchar(50) NOT NULL DEFAULT '',
 `password` varchar(50) NOT NULL DEFAULT '',
 `nickname` varchar(50) NOT NULL DEFAULT '',
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL DEFAULT NULL,
 `avator` varchar(100) NOT NULL DEFAULT '',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

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
 PRIMARY KEY (`id`),
 UNIQUE KEY `visitor_id` (`visitor_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;