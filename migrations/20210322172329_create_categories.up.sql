CREATE TABLE IF NOT EXISTS `categories` (
    `category_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(150) NOT NULL,
    `status` tinyint(1) DEFAULT '1',
    PRIMARY KEY (`category_id`),
    UNIQUE KEY `title` (`title`)
);

-- Тестовые данные
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (1, 'Напитки', 1);
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (2, 'Одежда', 1);
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (3, 'Шоколадки', 1);
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (4, 'Бакалея', 1);
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (5, 'Мясо', 1);
INSERT INTO `categories` (`category_id`, `title`, `status`) VALUES (6, 'Яйца', 0);