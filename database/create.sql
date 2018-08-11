CREATE DATABASE  IF NOT EXISTS `restapi`;

USE `restapi`


DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`username` varchar(150) NOT NULL,
`password` varchar(128) NOT NULL,
PRIMARY KEY (`id`),
UNIQUE KEY `username` (`username`)
) 
ENGINE=InnoDB 
AUTO_INCREMENT=2 
DEFAULT CHARSET=latin1;

LOCK TABLES `auth_user` WRITE;
INSERT INTO `auth_user` VALUES
(1,'admin','$2a$08$685B.DXF0SUOpgkMzNWzw.T8gt/wDSjanU/nVArOoAkL42SjYT3oG');
UNLOCK TABLES;


DROP TABLE IF EXISTS `announcement_announcement`;
CREATE TABLE `announcement_announcement` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`text` longtext NOT NULL,
`datefrom` date DEFAULT NULL,
`dateto` date DEFAULT NULL,
PRIMARY KEY (`id`)
) 
ENGINE=InnoDB 
AUTO_INCREMENT=23 
DEFAULT CHARSET=latin1 
COMMENT='Announcements';

LOCK TABLES `announcement_announcement` WRITE;
INSERT INTO `announcement_announcement` VALUES 
(1,'Example 1','2018-08-01','2018-12-31'),
(2,'Example 2','2018-08-15','2018-09-30');
UNLOCK TABLES;
