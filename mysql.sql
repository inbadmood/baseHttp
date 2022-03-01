CREATE DATABASE  IF NOT EXISTS `example` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `example`;
--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
CREATE TABLE `User` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                          `name` varchar(200) COLLATE utf8_unicode_ci DEFAULT '""',
                          `serial` varchar(50) NOT NULL DEFAULT '""',
                          `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
                          `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
INSERT INTO `User` (`id`,`name`,`serial`) VALUES (1,'Mike','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (2,'Sean','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (3,'Aries','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (4,'Roy','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (5,'Ryan','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (6,'Ellen','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (7,'Noel','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (8,'Jerry','0');
INSERT INTO `User` (`id`,`name`,`serial`)  VALUES (9,'James','0');
UNLOCK TABLES;
