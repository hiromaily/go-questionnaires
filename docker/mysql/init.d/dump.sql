-- MySQL dump 10.13  Distrib 5.7.12, for osx10.11 (x86_64)
--
-- Host: 127.0.0.1    Database: questionnaire
-- ------------------------------------------------------
-- Server version	5.7.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


--
-- DATABASE questionnaire
--
DROP DATABASE IF EXISTS `questionnaire`;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `questionnaire` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `questionnaire`;


--
-- Table structure for table `t_questionnaires`
--

DROP TABLE IF EXISTS `t_questionnaires`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_questionnaires` (
  `questionnaire_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'questionnaire ID',
  `title` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'title',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `created` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `updated` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`questionnaire_id`),
  KEY `idx_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Questionnaires Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_questionnaires` WRITE;
/*!40000 ALTER TABLE `t_questionnaires` DISABLE KEYS */;
INSERT INTO `t_questionnaires` VALUES (1,'title1','0','2016-09-24 21:43:15','2016-09-24 21:43:15'),(2,'title2','0','2016-09-24 21:43:15','2016-09-24 21:43:15'),(3,'title3','0','2016-09-24 21:43:15','2016-09-24 21:43:15');
/*!40000 ALTER TABLE `t_questionnaires` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_questions`
--

DROP TABLE IF EXISTS `t_questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_questions` (
  `questions_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'question ID',
  `questionnaire_id` int(11) NOT NULL COMMENT 'questionnaire ID',
  `order_no` smallint NOT NULL COLLATE utf8_unicode_ci NOT NULL COMMENT 'question order no',
  `question` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'question',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`questions_id`),
  KEY `idx_questionnaire` (`questionnaire_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Questions Table';
/*!40101 SET character_set_client = @saved_cs_client */;


LOCK TABLES `t_questions` WRITE;
/*!40000 ALTER TABLE `t_questions` DISABLE KEYS */;
INSERT INTO `t_questions` VALUES (1,1,1, 'q1','0','2016-09-24 21:43:15','2016-09-24 21:43:15'),(2,1,2,'q2','0','2016-09-24 21:43:15','2016-09-24 21:43:15'),(3,1,3,'q3','0','2016-09-24 21:43:15','2016-09-24 21:43:15');
/*!40000 ALTER TABLE `t_questions` ENABLE KEYS */;
UNLOCK TABLES;


--
-- Table structure for table `t_answers`
--

DROP TABLE IF EXISTS `t_answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_answers` (
  `answers_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'question ID',
  `questionnaire_id` int(11) NOT NULL COMMENT 'questionnaire ID',
  `user_email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'question',
  `answers` json DEFAULT NULL COMMENT 'answers by json data',
  `delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
  `create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  PRIMARY KEY (`answers_id`),
  KEY `idx_questionnaire` (`questionnaire_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Questions Table';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `t_answers` WRITE;
/*!40000 ALTER TABLE `t_answers` DISABLE KEYS */;
INSERT INTO `t_answers` VALUES (1,1,'abc@gmail.com', '{"answer": [{"order_no": 1, "answer": "answer1"},{"order_no": 2, "answer": "answer2"},{"order_no": 3, "answer": "answer3"}]}','0','2016-09-24 21:43:15','2016-09-24 21:43:15'),(2,1,'xxxx@gmail.com','{"answer": [{"order_no": 1, "answer": "answer--1"},{"order_no": 2, "answer": "answer--2"},{"order_no": 3, "answer": "answer--3"}]}','0','2016-09-24 21:43:15','2016-09-24 21:43:15');
/*!40000 ALTER TABLE `t_answers` ENABLE KEYS */;
UNLOCK TABLES;



/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

