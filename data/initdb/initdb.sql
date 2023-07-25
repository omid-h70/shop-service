SET time_zone="+3:30";

DROP DATABASE IF EXISTS shop_service_db;
CREATE DATABASE shop_service_db;
USE shop_service_db;

DROP TABLE IF EXISTS `vendors`;
CREATE TABLE `vendors`(
`vendor_id` int(5) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`phone_number` varchar(32) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`vendor_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

INSERT INTO `vendors` VALUES
(1001, 'Test Shop 1',  '+989123993699', 1),
(1002, 'Test Shop 2',  '+989033934262', 1),
(1003, 'Test Shop 3',  '+989123993699', 1),
(1004, 'Test Shop 4',  '+989033934262', 1),
(1005, 'Test Shop 5',  '+989123993699', 1);

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`(
`order_id` int(5) NOT NULL AUTO_INCREMENT,
`vendor_id` int(5) NOT NULL,
`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
#`delivery_time` TIMESTAMP NOT NULL DEFAULT (NOW()+INTERVAL 1 HOUR),
`delivery_time` TIMESTAMP NOT NULL DEFAULT (NOW()+INTERVAL 1 MINUTE ),
`order_status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`order_id`),
KEY `vendors_fk` (`vendor_id`),
CONSTRAINT `vendors_fk` FOREIGN KEY (`vendor_id`) REFERENCES `vendors` (`vendor_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 2007 DEFAULT CHARSET=latin1;

CREATE INDEX `idx_orders_created_at` ON `orders`(`created_at`);
CREATE INDEX `idx_orders_delivery_time` ON `orders`(`delivery_time`);

INSERT INTO `orders` (`order_id`, `vendor_id`) VALUES
(2001, 1001),
(2002, 1002),
(2003, 1003),
(2004, 1004),
(2005, 1005);


DROP TABLE IF EXISTS `trips`;
CREATE TABLE `trips`(
`trip_id` int(5) NOT NULL AUTO_INCREMENT,
`order_id` int(5) NOT NULL,
`trip_status` ENUM('AT_VENDOR', 'ASSIGNED', 'PICKED', 'DELIVERED') DEFAULT 'AT_VENDOR',
PRIMARY KEY (`trip_id`),
KEY `trips_order_id_fk` (`order_id`),
CONSTRAINT `trips_order_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 3002 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `agents`;
CREATE TABLE `agents`(
`agent_id` int(5) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
PRIMARY KEY (`agent_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 4002 DEFAULT CHARSET=latin1;

INSERT INTO `agents` (`agent_id`, `name`) VALUES
(4001, 'admin 1'),
(4002, 'admin 2'),
(4003, 'admin 3');

DROP TABLE IF EXISTS `delay_reports`;
CREATE TABLE `delay_reports`(
`delay_report_id` int(5) NOT NULL AUTO_INCREMENT,
`order_id` int(5) NOT NULL,
`vendor_id` int(5) NOT NULL,
`agent_id` int(5) DEFAULT NULL,
`report_count` int(5) NOT NULL DEFAULT 1,
`delay_report_status` ENUM('OPEN', 'CLOSED') DEFAULT 'OPEN',
`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`delay_report_id`),
KEY `vendor_id_fk` (`vendor_id`),
CONSTRAINT `vendor_id_fk` FOREIGN KEY (`vendor_id`) REFERENCES `vendors` (`vendor_id`),
KEY `order_id_fk` (`order_id`),
CONSTRAINT `order_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`),
KEY `agent_id_fk` (`agent_id`),
CONSTRAINT `agent_id_fk` FOREIGN KEY (`agent_id`) REFERENCES `agents` (`agent_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 5001 DEFAULT CHARSET=latin1;

CREATE INDEX `idx_delay_reports_created_at` ON `delay_reports`(`created_at`);
CREATE INDEX `idx_delay_reports_updated_at` ON `delay_reports`(`updated_at`);