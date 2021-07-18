-- MySQL dump 10.13  Distrib 8.0.18, for osx10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: tasa
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `project_post_images`
--

DROP TABLE IF EXISTS `project_post_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project_post_images` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_post_id` int(11) DEFAULT NULL,
  `url` varchar(1000) DEFAULT NULL,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `is_default` tinyint(1) DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `project_post_images_project_posts_id_fk` (`project_post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_post_images`
--

LOCK TABLES `project_post_images` WRITE;
/*!40000 ALTER TABLE `project_post_images` DISABLE KEYS */;
INSERT INTO `project_post_images` (`id`, `project_post_id`, `url`, `updated_at`, `created_at`, `is_default`, `deleted_at`) VALUES (2,4,'https://i.imgur.com/PCXtnvX.jpg','2019-11-24 13:52:50','2019-11-17 20:00:55',0,'2019-11-24 13:52:51'),(3,2,'','2019-12-08 10:13:17','2019-11-17 20:00:55',1,NULL),(4,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 12:30:15','2019-11-24 12:30:15',0,NULL),(5,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-19 23:35:46','2019-11-17 20:00:34',1,NULL),(6,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-19 23:35:46','2019-11-17 20:00:34',1,NULL),(7,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-19 23:35:46','2019-11-17 20:00:34',1,NULL),(8,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 12:33:14','2019-11-24 12:33:14',0,NULL),(9,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 12:34:59','2019-11-24 12:34:59',0,NULL),(10,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 12:51:08','2019-11-24 12:51:08',0,NULL),(11,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 13:50:27','2019-11-24 13:50:27',0,NULL),(12,4,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 13:51:14','2019-11-24 13:51:14',0,NULL),(13,5,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 13:53:04','2019-11-24 13:53:04',0,NULL),(14,5,'https://i.imgur.com/Fj3nnwZ.jpg','2019-11-24 13:53:20','2019-11-24 13:53:07',0,'2019-11-24 13:53:21');
/*!40000 ALTER TABLE `project_post_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_posts`
--

DROP TABLE IF EXISTS `project_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project_posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `body` varchar(5000) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `subtitle` varchar(255) DEFAULT NULL,
  `section` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_posts`
--

LOCK TABLES `project_posts` WRITE;
/*!40000 ALTER TABLE `project_posts` DISABLE KEYS */;
INSERT INTO `project_posts` (`id`, `title`, `body`, `created_at`, `updated_at`, `deleted_at`, `subtitle`, `section`) VALUES (4,'updated','fuck body','2019-11-19 16:19:41','2019-12-08 08:51:25',NULL,'Hanooi','project'),(5,'whatsup','insert body','2019-12-08 09:01:58','2019-12-08 09:21:01',NULL,'','inspiratio'),(6,'whatsup','insert body','2019-12-08 09:01:58','2019-12-08 09:21:01',NULL,'','inspiratio'),(7,'whatsup','insert body','2019-12-08 10:05:41','2019-12-08 10:05:41',NULL,'','inspiratiomn');
/*!40000 ALTER TABLE `project_posts` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-03 18:36:20
