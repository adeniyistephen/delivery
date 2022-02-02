DROP TABLE IF EXISTS `delivery_option`;

CREATE TABLE `delivery_option` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `delivery_option` */

insert  into `delivery_option`(`id`,`name`) values 
(4,'Dropship'),
(3,'Parcel');

/* region table */

DROP TABLE IF EXISTS `region`;

CREATE TABLE `region` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `region` */

insert  into `region`(`id`,`name`) values 
(3,'Luzon'),
(4,'Vis/Min');

/*Table structure for table `sysparam` */

DROP TABLE IF EXISTS `sysparam`;

CREATE TABLE `sysparam` (
  `key` varchar(255) NOT NULL,
  `value` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `sysparam` */

insert  into `sysparam`(`key`,`value`) values 
('HANDLER_DROPSHIP_LUZON','Max88official@gmail.com'),
('HANDLER_DROPSHIP_VISMIN','Max88cebustaff@gmail.com'),
('HANDLER_PACKAGE_LUZON','marcoavilam88@gmail.com'),
('HANDLER_PACKAGE_VISMIN','beautybeyondhealth@gmail.com'),
('HANDLER_PARCEL_LUZON','marcoavilam88@gmail.com'),
('HANDLER_PARCEL_VISMIN','beautybeyondhealth@gmail.com'),
('PACKAGE_PURCHASE_DEDUCTION','10750'),
('PACKAGE_PURCHASE_LIMIT','10750'),
('PRODUCT_EXCLUSIONS_DROPSHIP','Max-Cee Blister'),
('PRODUCT_EXCLUSIONS_PACKAGE','Max-Cee'),
('PRODUCT_EXCLUSIONS_PARCEL','Max-Cee'),
('WITHDRAWAL_FEE','50');

/*Table structure for table `user_type` */

DROP TABLE IF EXISTS `user_type`;

CREATE TABLE `user_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=latin1;

/*Data for the table `user_type` */

insert  into `user_type`(`id`,`name`) values 
(15,'Admin'),
(16,'Seller'),
(17,'Dropshipper');

/*Table structure for table `bank_type` */

DROP TABLE IF EXISTS `bank_type`;

CREATE TABLE `bank_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `date_added` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `date_modified` timestamp NULL DEFAULT NULL,
  `voided` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=latin1;

/*Data for the table `bank_type` */

insert  into `bank_type`(`id`,`name`,`date_added`,`date_modified`,`voided`) values 
(1,'BDO','2020-05-31 12:58:33',NULL,NULL),
(2,'UNION BANK','2020-05-31 12:58:38',NULL,NULL),
(3,'METROBANK','2020-05-31 12:58:44',NULL,NULL),
(4,'Asia United Bank (AUB)','2020-08-11 05:07:37',NULL,NULL),
(5,'Bank of Commerce','2020-08-11 05:07:37',NULL,NULL),
(6,'BPI','2020-08-11 05:07:37',NULL,NULL),
(7,'ChinaBank','2020-08-11 05:07:37',NULL,NULL),
(8,'Development Bank of Phil (DBP)','2020-08-11 05:07:37',NULL,NULL),
(9,'Eastwest Bank','2020-08-11 05:07:37',NULL,NULL),
(10,'Land Bank of the Phil','2020-08-11 05:07:37',NULL,NULL),
(11,'Maybank','2020-08-11 05:07:37',NULL,NULL),
(12,'PBCom','2020-08-11 05:07:37',NULL,NULL),
(13,'RCBC','2020-08-11 05:07:37',NULL,NULL),
(14,'Robinsons Bank','2020-08-11 05:07:37',NULL,NULL),
(15,'Security Bank','2020-08-11 05:07:37',NULL,NULL),
(16,'UCPB','2020-08-11 05:07:37',NULL,NULL),
(17,'GCASH','2020-08-12 18:25:44',NULL,NULL);


CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `mobile_number` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `user_type_id` int NOT NULL,
  `created_by` int DEFAULT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` int DEFAULT NULL,
  `is_active` tinyint DEFAULT '1',
  `bank_type_id` int NOT NULL,
  `bank_no` varchar(255) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `m88_account` varchar(255) DEFAULT NULL,
  `region_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `user_type` (`user_type_id`),
  KEY `user_ibfk_2` (`created_by`),
  KEY `bank_type_id` (`bank_type_id`),
  KEY `region_id` (`region_id`),
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_ibfk_4` FOREIGN KEY (`user_type_id`) REFERENCES `user_type` (`id`),
  CONSTRAINT `user_ibfk_5` FOREIGN KEY (`bank_type_id`) REFERENCES `bank_type` (`id`),
  CONSTRAINT `user_ibfk_6` FOREIGN KEY (`region_id`) REFERENCES `region` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1416 DEFAULT CHARSET=utf8mb3;

/*Data for the table `user` */

insert  into `user`(`id`,`firstname`,`lastname`,`email`,`mobile_number`,`password`,`user_type_id`,`created_by`,`created_date`,`last_updated`,`updated_by`,`is_active`,`bank_type_id`,`bank_no`,`address`,`birthday`,`gender`,`m88_account`,`region_id`) values 
(26,'Demby','Demby','dmcchy@gmail.com','09176799255','$2a$10$QLqxZEw9FBFeLni2zaxRWeXKkN4Fo2YNvLYY3OZrroc7CkzaZGDJK',15,NULL,'2020-06-06 14:14:55','2020-09-02 12:27:01',NULL,1,1,'091232','Palao','2020-07-17','F','2321',NULL),
(37,'Al John','Timogan','aljohntimogan@gmail.com','09176336836','$2a$10$VGenW2O2IB.LRDWvD2TDoeL8LhOLg8iz40W3GuedKPsOZtE884nLq',15,NULL,'2020-07-20 11:08:35','2020-07-25 18:37:27',NULL,1,1,'003128012043','0007 Villa Donna Subd., Purok 5, Upper Hinaplanon, Iligan City','1993-06-02','M','alphaajt',NULL),
(38,'William','Chavez','wilsh_lc@yahoo.com','09173770565','$2a$10$7Dp40h.5.wLFZjrL0qikpuzMEOYcis0OfJdm7AsnF5GwN9bjWbNqG',15,NULL,'2020-07-20 18:18:29','2020-07-23 20:22:00',NULL,1,1,'100660060168','23d padgett place','1980-08-26','M','My Account',NULL),
(40,'dgdsagadsgsa','gasdgdsaga','gasgsadgsad@gmail.com','9176336836','$2a$10$rV0UkIF3MeaZYBUUB2ry1OYpx2gpINYz4iRZKrsuDN6NDJ6RnWSNm',16,NULL,'2020-07-22 17:51:39','2020-08-06 15:12:40',NULL,0,1,'003128012043','0007 Villa Donna Subd., Purok 5, Upper Hinaplanon, Iligan City','2020-07-14','F','alphaajt',NULL),
(147,'M88','Luzon','Max88official@gmail.com','09351415087','$2a$10$oVoRXVxNGPvcgw56tQ1/au3IVJVlLNuJedCpXry96xQIsIMXOv6Fu',17,NULL,'2020-08-09 11:44:26',NULL,NULL,1,1,'005548024688','Mandaluyong City','1999-01-01','M','n/a',NULL),
(132,'Droppy','Luzon','marcoavilam88@gmail.com','09199844180','$2a$10$PuBCd7WdoGGiShcVwKAbder9hNy2a9gWCHKApsEUSmIgedINuLc5K',17,NULL,'2020-08-06 16:13:46','2020-08-09 11:34:20',NULL,1,1,'1094-5245-6563','581 Lagman 2, Coloong 1, Valenzuela City','1992-10-29','M','MarcoAvila24',NULL),
(43,'Droppy','VisMin','beautybeyondhealth@gmail.com','9176336836','$2a$10$mUATLSI9Dch29.5Pk/d8MeZWsrJB/udA2dc2WwoFGMW19IYC7gPqG',17,NULL,'2020-08-04 10:30:29','2020-08-09 11:09:15',NULL,1,1,'9359458565','Purok 5, Upper Hinaplanon','1993-06-02','M','n/a',NULL),
(41,'My','Seller','my_seller@gg.com','12321','$2a$10$24WGeC9w/S3MEbUMD2OOPeN5EmeNn8.AiuyTC7So9QWtsZ1rFxFOe',16,NULL,'2020-07-25 19:00:45','2020-08-06 09:28:26',NULL,0,1,'123','123','2020-07-29','F','123123',NULL);

DROP TABLE IF EXISTS `user_total`;

CREATE TABLE `user_total` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` decimal(65,2) NOT NULL,
  `coin_amount` decimal(65,2) NOT NULL,
  `created_by` int DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  `last_updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_2` (`user_id`),
  KEY `user_id` (`user_id`),
  KEY `created_by` (`created_by`),
  KEY `updated_by` (`updated_by`),
  -- CONSTRAINT `user_total_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_2` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_4` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1856 DEFAULT CHARSET=latin1;

/*Data for the table `user_total` */

insert  into `user_total`(`id`,`user_id`,`amount`,`coin_amount`,`created_by`,`updated_by`,`last_updated`) values 
(735,26,0.00,0.00,38,NULL,NULL),
(736,37,20474558.00,783902.00,38,NULL,NULL),
(737,38,-5066412.00,14986187.00,38,NULL,NULL),
(738,40,0.00,0.00,38,NULL,NULL),
(739,147,0.00,0.00,38,NULL,NULL),
(740,132,0.00,0.00,38,NULL,NULL),
(741,43,0.00,0.00,38,NULL,NULL),
(742,41,0.00,0.00,38,NULL,NULL);