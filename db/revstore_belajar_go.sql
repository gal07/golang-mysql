/*
SQLyog Community v13.1.1 (64 bit)
MySQL - 8.0.30 : Database - revstore_belajar_go
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`revstore_belajar_go` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `revstore_belajar_go`;

/*Table structure for table `tb_lesson` */

DROP TABLE IF EXISTS `tb_lesson`;

CREATE TABLE `tb_lesson` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `teacher` int DEFAULT NULL,
  `status` int DEFAULT '1',
  `isdelete` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uq_teacher` (`teacher`),
  CONSTRAINT `fk_teacher` FOREIGN KEY (`teacher`) REFERENCES `tb_teacher` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tb_lesson` */

LOCK TABLES `tb_lesson` WRITE;

insert  into `tb_lesson`(`id`,`name`,`teacher`,`status`,`isdelete`,`created_at`,`deleted_at`) values 
(1,'Algebra',49,1,1,'2024-10-19 01:07:39','2024-10-19 01:07:39'),
(2,'Math',2,1,0,'2024-10-19 01:09:44','2024-10-19 01:09:44'),
(9,'Monolith',23,1,0,'2024-10-20 02:28:45','2024-10-20 02:28:45'),
(10,'Sains',51,1,0,'2024-10-20 02:39:01','2024-10-20 02:39:01'),
(31,'Gravity',25,1,0,'2024-10-22 22:47:51','2024-10-22 22:47:51'),
(32,'Algorithm',26,1,0,'2024-10-22 22:48:02','2024-10-22 22:48:02'),
(33,'Pyhsich',27,1,0,'2024-10-22 22:48:21','2024-10-22 22:48:21'),
(35,'Structure Data',48,1,0,'2024-10-22 22:49:01','2024-10-22 22:49:01'),
(37,'Networking',50,1,0,'2024-10-22 23:54:19','2024-10-22 23:54:19'),
(38,'Finance',1,1,0,'2024-10-23 00:09:41','2024-10-23 00:09:41'),
(41,'Optimize',49,1,0,'2024-10-23 00:20:15','2024-10-23 00:20:15');

UNLOCK TABLES;

/*Table structure for table `tb_report_card` */

DROP TABLE IF EXISTS `tb_report_card`;

CREATE TABLE `tb_report_card` (
  `id` int NOT NULL AUTO_INCREMENT,
  `student` int DEFAULT NULL,
  `lesson` int DEFAULT NULL,
  `grade` int DEFAULT NULL,
  `status` int DEFAULT '1',
  `isdelete` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `SECOND` (`student`),
  KEY `THIRD` (`lesson`),
  CONSTRAINT `fk_lesson` FOREIGN KEY (`lesson`) REFERENCES `tb_lesson` (`id`),
  CONSTRAINT `fk_student` FOREIGN KEY (`student`) REFERENCES `tb_student` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tb_report_card` */

LOCK TABLES `tb_report_card` WRITE;

UNLOCK TABLES;

/*Table structure for table `tb_student` */

DROP TABLE IF EXISTS `tb_student`;

CREATE TABLE `tb_student` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `age` int NOT NULL,
  `grade` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `isdelete` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=640159 DEFAULT CHARSET=latin1;

/*Data for the table `tb_student` */

LOCK TABLES `tb_student` WRITE;

UNLOCK TABLES;

/*Table structure for table `tb_teacher` */

DROP TABLE IF EXISTS `tb_teacher`;

CREATE TABLE `tb_teacher` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `status` int DEFAULT '1',
  `isdelete` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tb_teacher` */

LOCK TABLES `tb_teacher` WRITE;

insert  into `tb_teacher`(`id`,`name`,`age`,`status`,`isdelete`,`created_at`,`deleted_at`) values 
(1,'Mr Hendro Soedjono',24,1,0,'2024-10-18 20:22:39','2024-10-18 20:22:39'),
(2,'Mr Hamdan',23,1,0,'2024-10-18 20:24:42','2024-10-18 20:24:42'),
(23,'Mr Amdan',21,1,0,'2024-10-20 02:28:16','2024-10-20 02:28:16'),
(24,'Mr Bani Aman',23,1,0,'2024-10-20 02:38:39','2024-10-20 02:38:39'),
(25,'Mr Anwar ibrahim',45,1,0,'2024-10-20 16:58:48','2024-10-20 16:58:48'),
(26,'Mr Sakda ibrahim',31,1,0,'2024-10-20 17:16:58','2024-10-20 17:16:58'),
(27,'Mrs Nami Namn',21,1,0,'2024-10-20 19:31:43','2024-10-20 19:31:43'),
(48,'Ann doberson',34,1,0,'2024-10-22 21:43:56','2024-10-22 21:43:56'),
(49,'Stanley Anedot',23,1,0,'2024-10-22 21:44:11','2024-10-22 21:44:11'),
(50,'Brendan Nosel',45,1,0,'2024-10-22 21:58:28','2024-10-22 21:58:28'),
(51,'Brendan Ager',33,1,0,'2024-10-22 21:58:36','2024-10-22 21:58:36');

UNLOCK TABLES;

/*Table structure for table `tb_token_whitelist` */

DROP TABLE IF EXISTS `tb_token_whitelist`;

CREATE TABLE `tb_token_whitelist` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT NULL,
  `token` text,
  `type` varchar(50) DEFAULT NULL,
  `expires` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tb_token_whitelist` */

LOCK TABLES `tb_token_whitelist` WRITE;

insert  into `tb_token_whitelist`(`id`,`email`,`token`,`type`,`expires`,`created_at`) values 
(36,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODI4NzUsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.fzduvllpjpp818Bf0xqdK52fwI6nZ3hogEPuyyf19SQ','refresh-token',1730382875,'2024-10-30 20:54:35'),
(37,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDAwNzUsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.TwBHDgvNtW60HaU1PkSa3IYwqwyJ0qvf27TnBCT2cRM','primary-token',1730382875,'2024-10-30 20:54:35'),
(38,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODI5OTgsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.StdrzrGTOf2fd4dDn_THfZUQu0SkzkerP1fTKh_dlmY','refresh-token',1730382998,'2024-10-30 20:56:38'),
(39,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDAxOTgsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.avtZ9IaNuPEqesGakw1825we011zdao3oHzruRdv_NA','primary-token',1730382998,'2024-10-30 20:56:38'),
(40,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODQzMDUsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.3DJciNyW_z3W27PwOazMVrtIRhLFLyn5UOHnY6Zm8Bs','refresh-token',1730384305,'2024-10-30 21:18:25'),
(41,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDE1MDUsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.J3dUJLoEZaRXhB2H1izsvJBaHJj1v4tp5vvDX_MO--E','primary-token',1730384305,'2024-10-30 21:18:25'),
(42,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODQ0MDUsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.gXA5qpoFwKxukQ8IevjZkfvywh63HlOdciSa0h8SL4M','refresh-token',1730384405,'2024-10-30 21:20:05'),
(43,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDE2MDUsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.dFn9KQ83PE0d_wBJarwEAYGjhU34Zf0Lzpxa2iQ9nJ8','primary-token',1730384405,'2024-10-30 21:20:05'),
(44,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODQ0NTksInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.2TDmtpTHiegYBnGtWb3zsr3Pr4SiWZ9xHH4f5592DYc','refresh-token',1730384459,'2024-10-30 21:20:59'),
(45,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDE2NTksInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.FMwrQeQN8JnMb2OGl6HFhtuCaxScxJJqJGMYPQqhJa0','primary-token',1730384459,'2024-10-30 21:20:59'),
(46,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODQ2MTEsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.NrbA7dmWUrAeWoxqpilua7o5cz68MpuO1GixHv_2umo','refresh-token',1730384611,'2024-10-30 21:23:31'),
(47,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDE4MTEsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.HSguOfGxJXYi4_AVhprQb2dy7kaoGAVmNKqOWRflC8A','primary-token',1730384611,'2024-10-30 21:23:31'),
(48,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODUxOTIsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.7dZ1H4RqKoUBA0YMPiE7KbywIYNxA8eI6LVTgQi5d00','refresh-token',1730385192,'2024-10-30 21:33:12'),
(49,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDIzOTIsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.lFLbcGOcrEIQkeXNI0ni8uBqyKojvrNTETbBcd_C3HU','primary-token',1730385192,'2024-10-30 21:33:12'),
(50,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODU4NzUsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.DTJisR2ovEXaUNqn0C2dIWNssSN09P5-VgjYf74aHA8','refresh-token',1730385875,'2024-10-30 21:44:35'),
(51,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzMDMwNzUsInR5cGUiOiJwcmltYXJ5LXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.52zzSCFNu3Nhs4i9c5c6vVvFB3_KwgW_IRpq3_t_jIA','primary-token',1730385875,'2024-10-30 21:44:35'),
(52,NULL,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzAzODg1NTMsInR5cGUiOiJyZWZyZXNoLXRva2VuIiwidXNlcm5hbWUiOiJnYWxpaGt1cjI5QGdtYWlsLmNvbSJ9.yUv8S2GLnjyQvSq4lz3qSpSjBUWtY340nqjtMA1zvwI','primary-token',1730388553,'2024-10-30 22:29:13');

UNLOCK TABLES;

/*Table structure for table `tb_user` */

DROP TABLE IF EXISTS `tb_user`;

CREATE TABLE `tb_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int DEFAULT '1',
  `isdelete` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_unique` (`email`),
  UNIQUE KEY `username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tb_user` */

LOCK TABLES `tb_user` WRITE;

insert  into `tb_user`(`id`,`username`,`email`,`password`,`salt`,`status`,`isdelete`,`created_at`,`deleted_at`) values 
(4,'gg07','galihkur29@gmail.com','$2a$04$KHR7id2zSa9HDqju1ucBL.F0ghL.4zFJw8Dw5SKtj8lwKYelSx4KG','',1,0,'2024-10-22 09:35:48','2024-10-22 09:35:48');

UNLOCK TABLES;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
