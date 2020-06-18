CREATE DATABASE `library` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE library;
CREATE TABLE `authors` (
  `id` int NOT NULL DEFAULT '1',
  `firstname` varchar(45) NOT NULL,
  `lastname` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `books` (
  `id` int NOT NULL,
  `title` varchar(45) NOT NULL,
  `fkAuthor` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fkAuthor_idx` (`fkAuthor`),
  CONSTRAINT `fkAuthor` FOREIGN KEY (`fkAuthor`) REFERENCES `authors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `library`.`authors` (`id`,`firstname`,`lastname`) VALUES (1,"George","Washington");
INSERT INTO `library`.`authors` (`id`,`firstname`,`lastname`) VALUES (2,"Penelope","Penultimate");
INSERT INTO `library`.`authors` (`id`,`firstname`,`lastname`) VALUES (3,"Magnus","Carlson");
INSERT INTO `library`.`books` (`id`,`title`,`fkAuthor`) VALUES (1,"My Journey With Jonny Appleseed",1);
INSERT INTO `library`.`books` (`id`,`title`,`fkAuthor`) VALUES (2,"The Secret to Super Powers",2);
INSERT INTO `library`.`books` (`id`,`title`,`fkAuthor`) VALUES (3,"My 'PEN' Ultimate Book on Screenwriting",2);
INSERT INTO `library`.`books` (`id`,`title`,`fkAuthor`) VALUES (4,"Power Chess: A Guide to Rage Quiting",3);
