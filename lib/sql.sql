
------------ 用户表 --------------
CREATE TABLE `gopj`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(100) NULL,
  `password` VARCHAR(225) NULL,
  `isRight` VARCHAR(45) NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`));

-------------- admin ------------
CREATE TABLE `gopj`.`admin` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(100) NULL,
`password` VARCHAR(255) NULL,
`isRight` VARCHAR(45) NULL,
`created_at` DATETIME NULL,
`updated_at` DATETIME NULL,
PRIMARY KEY (`id`));

------------- 投诉表 -------------
CREATE TABLE `gopj`.`complaint` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `phone` VARCHAR(45) NULL,
  `be_complainted` VARCHAR(45) NULL, ----被投诉人或群里
  `address` VARCHAR(255) NULL,       ----被投诉人地址 
  `reason` VARCHAR(255) NULL,
  `deal_person` VARCHAR(45) NULL,    ----处理人
  `status` VARCHAR(45) NULL,         ----状态
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  PRIMARY KEY (`id`));
