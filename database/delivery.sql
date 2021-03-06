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

/*Table structure for table `delivery_status` */

DROP TABLE IF EXISTS `delivery_status`;

CREATE TABLE `delivery_status` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `delivery_status` */

insert  into `delivery_status`(`id`,`name`) values 
(7,'Accepted'),
(9,'Delivered'),
(8,'Fulfilled'),
(6,'Proposed'),
(10,'Rejected'),
(11,'Returned'),
(12,'Voided');

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
(101,'Jan Dane Mar','Tura','jtura115@gmail.com','09664153128','$2a$10$OeobOMP/LCPDeiErrCg2Ju92eg07dwCVyZaCYlaeL.N4e6ytd5QiW',16,NULL,'2020-08-06 10:11:56','2020-09-17 03:34:04',NULL,1,1,'304010036333','Sitio Pundok Pulangbato st. Pit os Cebu Cuty','1999-01-01','M','816603609',NULL),
(41,'My','Seller','my_seller@gg.com','12321','$2a$10$24WGeC9w/S3MEbUMD2OOPeN5EmeNn8.AiuyTC7So9QWtsZ1rFxFOe',16,NULL,'2020-07-25 19:00:45','2020-08-06 09:28:26',NULL,0,1,'123','123','2020-07-29','F','123123',NULL);

DROP TABLE IF EXISTS `user_total`;

CREATE TABLE `user_total` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userid` int NOT NULL,
  `amount` decimal(65,2) NOT NULL,
  `coinamount` decimal(65,2) NOT NULL,
  `createdby` int DEFAULT NULL,
  `updatedby` int DEFAULT NULL,
  `lastupdated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_2` (`userid`),
  KEY `userid` (`userid`),
  KEY `createdby` (`createdby`),
  KEY `updatedby` (`updatedby`),
  -- CONSTRAINT `user_total_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_2` FOREIGN KEY (`createdby`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_3` FOREIGN KEY (`userid`) REFERENCES `user` (`id`),
  CONSTRAINT `user_total_ibfk_4` FOREIGN KEY (`updatedby`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1856 DEFAULT CHARSET=latin1;

/*Data for the table `user_total` */

insert  into `user_total`(`id`,`userid`,`amount`,`coinamount`,`createdby`,`updatedby`,`lastupdated`) values 
(735,26,0.00,0.00,38,NULL,NULL),
(736,37,20474558.00,783902.00,38,NULL,NULL),
(737,38,-5066412.00,14986187.00,38,NULL,NULL),
(738,40,0.00,0.00,38,NULL,NULL),
(739,147,0.00,0.00,38,NULL,NULL),
-- (740,132,0.00,0.00,38,NULL,NULL),
(741,43,0.00,0.00,38,NULL,NULL),
(742,41,0.00,0.00,38,NULL,NULL),
(743,101,0.00,250.00,38,38,'2020-09-02 12:27:01');

/*Table structure for table `product_type` */

DROP TABLE IF EXISTS `product_type`;

CREATE TABLE `product_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `product_type` */

insert  into `product_type`(`id`,`name`) values 
(1,'Consumables'),
(2,'Cosmetics'),
(3,'Others');

/*Table structure for table `product` */

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `producttypeid` int NOT NULL,
  `createdby` int DEFAULT NULL,
  `createddate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `lastupdated` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedby` varchar(255) DEFAULT NULL,
  `isactive` tinyint DEFAULT '1',
  `url` text,
  `priceperitem` decimal(65,2) NOT NULL DEFAULT '500.00',
  `priceperitemdropshipper` decimal(65,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `producttype` (`producttypeid`),
  KEY `user_ibfk_3` (`createdby`),
  CONSTRAINT `product_ibfk_5` FOREIGN KEY (`producttypeid`) REFERENCES `product_type` (`id`),
  CONSTRAINT `user_ibfk_3` FOREIGN KEY (`createdby`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=latin1;

/*Data for the table `product` */

insert  into `product`(`id`,`name`,`producttypeid`,`createdby`,`createddate`,`lastupdated`,`updatedby`,`isactive`,`url`,`priceperitem`,`priceperitemdropshipper`) values 
(15,'Max-Cee',1,26,'2020-06-08 14:48:30','2020-08-10 09:12:35',NULL,1,'/images/Max-Cee.JPG',1560.00,1092.00),
(16,'PPAR',1,26,'2020-06-08 14:49:48','2020-08-13 06:25:47','',1,'/images/PPAR.jpeg',1600.00,1120.00),
(17,'Maxijuice',1,26,'2020-06-08 15:19:45','2020-08-09 13:12:43','',1,'/images/Maxijuice.jpeg',560.00,392.00),
(24,'Tamaraw +',1,26,'2020-07-18 23:39:02','2020-08-09 13:12:55',NULL,1,'/images/Tamaraw +.jpeg',1950.00,1365.00),
(31,'Vert',2,26,'2020-07-25 18:05:19','2020-08-09 13:13:05',NULL,1,'/images/Vert.jpeg',250.00,175.00);


/*Table structure for table `inventory` */

DROP TABLE IF EXISTS `inventory`;

CREATE TABLE `inventory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `productid` int NOT NULL,
  `quantity` int NOT NULL,
  `createdby` int NOT NULL,
  `updatedby` int DEFAULT NULL,
  `createddate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `isactive` tinyint NOT NULL,
  `regionid` int NOT NULL,
  `sellerid` int NOT NULL,
  `dropshipperid` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `productid` (`productid`,`regionid`,`sellerid`,`dropshipperid`),
  KEY `createdby` (`createdby`),
  KEY `updatedby` (`updatedby`),
  KEY `regionid` (`regionid`),
  KEY `sellerid` (`sellerid`),
  KEY `dropshipperid` (`dropshipperid`),
  CONSTRAINT `inventory_ibfk_1` FOREIGN KEY (`productid`) REFERENCES `product` (`id`),
  CONSTRAINT `inventory_ibfk_2` FOREIGN KEY (`createdby`) REFERENCES `user` (`id`),
  CONSTRAINT `inventory_ibfk_3` FOREIGN KEY (`updatedby`) REFERENCES `user` (`id`),
  CONSTRAINT `inventory_ibfk_5` FOREIGN KEY (`regionid`) REFERENCES `region` (`id`),
  CONSTRAINT `inventory_ibfk_6` FOREIGN KEY (`sellerid`) REFERENCES `user` (`id`),
  CONSTRAINT `inventory_ibfk_7` FOREIGN KEY (`dropshipperid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4641 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `inventory` */

insert  into `inventory`(`id`,`productid`,`quantity`,`createdby`,`updatedby`,`createddate`,`isactive`,`regionid`,`sellerid`,`dropshipperid`) values 
(550,16,10,101,101,'2021-07-15 10:30:12',1,3,101,132),
(551,17,8,101,101,'2021-01-13 06:15:34',1,3,101,147),
(552,31,0,101,NULL,'2021-12-01 11:21:28',1,4,101,43);

/*Table structure for table `delivery` */

DROP TABLE IF EXISTS `delivery`;

CREATE TABLE `delivery` (
  `id` int NOT NULL AUTO_INCREMENT,
  `createdby` int NOT NULL,
  `updatedby` int DEFAULT NULL,
  `lastupdated` varchar(255) NOT NULL DEFAULT '',
  `createddate` varchar(255) NOT NULL DEFAULT '',
  `isactive` tinyint NOT NULL,
  `name` text NOT NULL,
  `address` text NOT NULL,
  `regionid` int NOT NULL,
  `servicefee` decimal(65,2) NOT NULL,
  `declaredamount` decimal(65,2) NOT NULL,
  `deliveryoptionid` int DEFAULT NULL,
  `deliverystatusid` int NOT NULL,
  `sellerid` int NOT NULL,
  `dropshipperid` int NOT NULL,
  `riderid` int DEFAULT NULL,
  `trackingnumber` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `contactnumber` varchar(255) NOT NULL,
  `note` text NOT NULL,
  `baseprice` decimal(65,2) NOT NULL,
  `amountdistributor` decimal(65,2) DEFAULT NULL,
  `voidorrejectreason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `trackingnumber` (`trackingnumber`),
  KEY `createdby` (`createdby`),
  KEY `updatedby` (`updatedby`),
  KEY `regionid` (`regionid`),
  KEY `deliveryoptionid` (`deliveryoptionid`),
  KEY `deliverystatusid` (`deliverystatusid`),
  KEY `sellerid` (`sellerid`),
  KEY `dropshipperid` (`dropshipperid`),
  KEY `riderid` (`riderid`),
  CONSTRAINT `delivery_ibfk_1` FOREIGN KEY (`createdby`) REFERENCES `user` (`id`),
  CONSTRAINT `delivery_ibfk_2` FOREIGN KEY (`updatedby`) REFERENCES `user` (`id`),
  CONSTRAINT `delivery_ibfk_3` FOREIGN KEY (`regionid`) REFERENCES `region` (`id`),
  CONSTRAINT `delivery_ibfk_4` FOREIGN KEY (`deliveryoptionid`) REFERENCES `delivery_option` (`id`),
  CONSTRAINT `delivery_ibfk_5` FOREIGN KEY (`deliverystatusid`) REFERENCES `delivery_status` (`id`),
  CONSTRAINT `delivery_ibfk_6` FOREIGN KEY (`sellerid`) REFERENCES `user` (`id`),
  CONSTRAINT `delivery_ibfk_7` FOREIGN KEY (`dropshipperid`) REFERENCES `user` (`id`),
  CONSTRAINT `delivery_ibfk_8` FOREIGN KEY (`riderid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10604 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `delivery` */

insert  into `delivery`(`id`,`createdby`,`updatedby`,`lastupdated`,`createddate`,`isactive`,`name`,`address`,`regionid`,`servicefee`,`declaredamount`,`deliveryoptionid`,`deliverystatusid`,`sellerid`,`dropshipperid`,`riderid`,`trackingnumber`,`contactnumber`,`note`,`baseprice`,`amountdistributor`,`voidorrejectreason`) values 
(337,101,101,"",'2020-08-12 18:36:07',1,'Jethro C. Sanchez','Caminade Compund, Hi-way 77, Talamban, Cebu City',4,195.00,800.00,4,9,101,147,NULL,'4735-4293-MFNQ','09177024565','Landmark: infront of F&M townhomes, green gate, Talamban, Cebu City',130.00,350.00,NULL),
(338,101,101,"",'2020-08-13 04:03:08',1,'Maria Lourder Jugueta','30 macapuno st. brgy ugong, valle verde 1, pasig city 1604',3,195.00,550.00,4,9,101,147,NULL,'4735-5408-SBZC','09275610095','(Live test)',130.00,350.00,NULL);


/*Table structure for table `delivery_detail` */

DROP TABLE IF EXISTS `delivery_detail`;

CREATE TABLE `delivery_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `deliveryid` int NOT NULL,
  `productid` int NOT NULL,
  `quantity` int NOT NULL,
  `priceperitemdistributor` decimal(65,2) DEFAULT NULL,
  `totalpricedistributor` decimal(65,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deliveryid` (`deliveryid`),
  KEY `productid` (`productid`),
  CONSTRAINT `delivery_detail_ibfk_1` FOREIGN KEY (`deliveryid`) REFERENCES `delivery` (`id`),
  CONSTRAINT `delivery_detail_ibfk_2` FOREIGN KEY (`productid`) REFERENCES `product` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35926 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `delivery_detail` */

insert  into `delivery_detail`(`id`,`deliveryid`,`productid`,`quantity`,`priceperitemdistributor`,`totalpricedistributor`) values 
(996,337,31,2,175.00,350.00),
(997,338,31,2,175.00,350.00);

/*Table structure for table `delivery_tracking` */

DROP TABLE IF EXISTS `delivery_tracking`;

CREATE TABLE `delivery_tracking` (
  `id` int NOT NULL AUTO_INCREMENT,
  `deliveryid` int NOT NULL,
  `deliverystatusid` int NOT NULL,
  `lastupdated` varchar(255) DEFAULT '',
  `updatedby` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `deliveryid` (`deliveryid`),
  KEY `deliverystatusid` (`deliverystatusid`),
  KEY `updatedby` (`updatedby`),
  CONSTRAINT `delivery_tracking_ibfk_1` FOREIGN KEY (`deliveryid`) REFERENCES `delivery` (`id`),
  CONSTRAINT `delivery_tracking_ibfk_2` FOREIGN KEY (`deliverystatusid`) REFERENCES `delivery_status` (`id`),
  CONSTRAINT `delivery_tracking_ibfk_3` FOREIGN KEY (`updatedby`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37014 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `delivery_tracking` */

insert  into `delivery_tracking`(`id`,`deliveryid`,`deliverystatusid`,`lastupdated`,`updatedby`) values 
(447,337,6,'2020-08-12 18:36:07',101),
(448,337,7,'2020-08-12 18:36:46',101);

/*Table structure for table `transaction` */

DROP TABLE IF EXISTS `transaction`;

CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_by` int DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  `created_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `is_active` tinyint DEFAULT NULL,
  `amount` decimal(15,2) DEFAULT NULL,
  `coin_amount` decimal(15,2) DEFAULT NULL,
  `admin_allotted_id` int DEFAULT NULL,
  `user_allotted_id` int DEFAULT NULL,
  `money_in` tinyint(1) DEFAULT NULL,
  `bank_type_id` int NOT NULL,
  `reference_number` varchar(255) NOT NULL,
  `description` text,
  `withdrawal_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `reference_number` (`reference_number`,`admin_allotted_id`,`user_allotted_id`,`money_in`),
  KEY `created_by` (`created_by`),
  KEY `updated_by` (`updated_by`),
  KEY `admin_alotted_id` (`admin_allotted_id`),
  KEY `user_alloted_id` (`user_allotted_id`),
  KEY `bank_type_id` (`bank_type_id`),
  CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT `transaction_ibfk_3` FOREIGN KEY (`admin_allotted_id`) REFERENCES `user` (`id`),
  CONSTRAINT `transaction_ibfk_4` FOREIGN KEY (`user_allotted_id`) REFERENCES `user` (`id`),
  CONSTRAINT `transaction_ibfk_5` FOREIGN KEY (`bank_type_id`) REFERENCES `bank_type` (`id`),
  CONSTRAINT `transaction_ibfk_6` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5692 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `transaction` */

insert  into `transaction`(`id`,`created_by`,`updated_by`,`created_date`,`is_active`,`amount`,`coin_amount`,`admin_allotted_id`,`user_allotted_id`,`money_in`,`bank_type_id`,`reference_number`,`description`,`withdrawal_id`) values 
(107,37,NULL,'2020-08-12 18:18:46',1,3000.00,3000.00,37,43,1,6,'1597220287871','for dropship',NULL),
(108,37,NULL,'2020-08-12 18:21:00',1,400.00,400.00,37,40,1,2,'155709','for dropship',NULL);

/*Table structure for table `withdrawal_status` */

DROP TABLE IF EXISTS `withdrawal_status`;

CREATE TABLE `withdrawal_status` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `withdrawal_status` */

insert  into `withdrawal_status`(`id`,`name`) values 
(1,'Completed'),
(2,'Rejected'),
(3,'Voided'),
(4,'Pending');

/*Table structure for table `withdrawal` */

DROP TABLE IF EXISTS `withdrawal`;

CREATE TABLE `withdrawal` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_by` int NOT NULL,
  `updated_by` int DEFAULT NULL,
  `created_date` timestamp NOT NULL,
  `is_active` tinyint NOT NULL DEFAULT '1',
  `void_or_reject_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_updated` timestamp NULL DEFAULT NULL,
  `amount` decimal(65,2) NOT NULL,
  `fee` decimal(65,2) NOT NULL,
  `total_amount` decimal(65,2) NOT NULL,
  `withdrawal_status_id` int NOT NULL,
  `user_id` int NOT NULL,
  `bank_no` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `bank_type_id` int NOT NULL,
  `bank_account_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reference_number` text,
  `contact_no` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `created_by` (`created_by`),
  KEY `withdrawal_status_id` (`withdrawal_status_id`),
  KEY `user_id` (`user_id`),
  KEY `bank_type_id` (`bank_type_id`),
  KEY `updated_by` (`updated_by`),
  CONSTRAINT `withdrawal_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT `withdrawal_ibfk_3` FOREIGN KEY (`withdrawal_status_id`) REFERENCES `withdrawal_status` (`id`),
  CONSTRAINT `withdrawal_ibfk_4` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `withdrawal_ibfk_5` FOREIGN KEY (`bank_type_id`) REFERENCES `bank_type` (`id`),
  CONSTRAINT `withdrawal_ibfk_6` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1660 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `withdrawal` */

insert  into `withdrawal`(`id`,`created_by`,`updated_by`,`created_date`,`is_active`,`void_or_reject_reason`,`last_updated`,`amount`,`fee`,`total_amount`,`withdrawal_status_id`,`user_id`,`bank_no`,`bank_type_id`,`bank_account_name`,`reference_number`,`contact_no`) values 
(1,38,38,'2020-09-06 01:45:54',1,NULL,'2020-09-07 07:42:42',500.00,20.00,500.00,1,40,'003128012043',1,'dgdsagadsgsa','185773-858751',NULL),
(2,38,38,'2020-09-06 05:11:46',1,'Test withdrawal','2020-09-06 06:10:25',610.00,20.00,610.00,3,43,'9359458565',1,'Droppy',NULL,NULL);

/*Table structure for table `order` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_by` int NOT NULL,
  `created_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_updated` timestamp NULL DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  `is_active` tinyint NOT NULL,
  `order_status_id` int NOT NULL,
  `amount` decimal(15,2) NOT NULL,
  `tracking_id` varchar(255) DEFAULT NULL,
  `region_id` int NOT NULL,
  `dropshipper_id` int NOT NULL,
  `seller_id` int NOT NULL,
  `admin_id` int NOT NULL,
  `void_or_reject_reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`id`),
  KEY `created_by` (`created_by`),
  KEY `updated_by` (`updated_by`),
  KEY `region_id` (`region_id`),
  KEY `dropshipper_id` (`dropshipper_id`),
  KEY `seller_id` (`seller_id`),
  KEY `admin_id` (`admin_id`),
  CONSTRAINT `order_ibfk_1` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT `order_ibfk_2` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`),
  CONSTRAINT `order_ibfk_4` FOREIGN KEY (`region_id`) REFERENCES `region` (`id`),
  CONSTRAINT `order_ibfk_5` FOREIGN KEY (`dropshipper_id`) REFERENCES `user` (`id`),
  CONSTRAINT `order_ibfk_6` FOREIGN KEY (`seller_id`) REFERENCES `user` (`id`),
  CONSTRAINT `order_ibfk_7` FOREIGN KEY (`admin_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2618 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `order` */

insert  into `order`(`id`,`created_by`,`created_date`,`last_updated`,`updated_by`,`is_active`,`order_status_id`,`amount`,`tracking_id`,`region_id`,`dropshipper_id`,`seller_id`,`admin_id`,`void_or_reject_reason`) values 
(371,101,'2020-08-13 17:01:57',NULL,NULL,1,5,10750.00,NULL,3,132,101,38,NULL),
(372,101,'2020-08-13 17:03:03',NULL,NULL,1,5,10750.00,NULL,4,43,101,38,NULL);

/*Table structure for table `coin_transaction` */

DROP TABLE IF EXISTS `coin_transaction`;

CREATE TABLE `coin_transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `createdby` int DEFAULT NULL,
  `updatedby` int DEFAULT NULL,
  `createddate` varchar(255) DEFAULT '',
  `isactive` tinyint DEFAULT NULL,
  `transactionid` int DEFAULT NULL,
  `userid` int DEFAULT NULL,
  `type` enum('D','C') DEFAULT NULL,
  `amount` decimal(65,2) DEFAULT NULL,
  `cointransactionid` int DEFAULT NULL,
  `orderid` int DEFAULT NULL,
  `deliveryid` int DEFAULT NULL,
  `withdrawalid` int DEFAULT NULL,
  `lastupdated` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `createdby` (`createdby`),
  KEY `updatedby` (`updatedby`),
  KEY `transactionid` (`transactionid`),
  KEY `userid` (`userid`),
  KEY `cointransactionid` (`cointransactionid`),
  KEY `coin_transaction_ibfk_5` (`orderid`),
  KEY `deliveryid` (`deliveryid`),
  KEY `withdrawalid` (`withdrawalid`),
  CONSTRAINT `coin_transaction_ibfk_1` FOREIGN KEY (`createdby`) REFERENCES `user` (`id`),
  CONSTRAINT `coin_transaction_ibfk_2` FOREIGN KEY (`updatedby`) REFERENCES `user` (`id`),
  CONSTRAINT `coin_transaction_ibfk_3` FOREIGN KEY (`transactionid`) REFERENCES `transaction` (`id`),
  CONSTRAINT `coin_transaction_ibfk_4` FOREIGN KEY (`userid`) REFERENCES `user` (`id`),
  CONSTRAINT `coin_transaction_ibfk_5` FOREIGN KEY (`orderid`) REFERENCES `order` (`id`),
  CONSTRAINT `coin_transaction_ibfk_6` FOREIGN KEY (`deliveryid`) REFERENCES `delivery` (`id`),
  CONSTRAINT `coin_transaction_ibfk_7` FOREIGN KEY (`withdrawalid`) REFERENCES `withdrawal` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72643 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `coin_transaction` */

insert  into `coin_transaction`(`id`,`createdby`,`updatedby`,`createddate`,`isactive`,`transactionid`,`userid`,`type`,`amount`,`cointransactionid`,`orderid`,`deliveryid`,`withdrawalid`,`lastupdated`) values 
(396,37,37,'2020-08-12 18:18:46',1,107,37,'D',3000.00,NULL,NULL,NULL,NULL,""),
(397,37,37,'2020-08-12 18:18:46',1,NULL,147,'C',-3000.00,396,NULL,NULL,NULL,"");