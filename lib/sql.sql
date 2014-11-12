
------------ 用户表 --------------
CREATE TABLE `gopj`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(100) NULL,
  `password` VARCHAR(225) NULL,
  `isRight` VARCHAR(45) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`));

------------ admin ------------
CREATE TABLE `gopj`.`admin` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(100) NULL,
`password` VARCHAR(255) NULL,
`isRight` VARCHAR(45) NULL,
`created_at` DATETIME NULL,
`updated_at` DATETIME NULL,
PRIMARY KEY (`id`));