CREATE TABLE IF NOT EXISTS `products` (
    `product_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(150) NOT NULL,
    `weight` float NOT NULL,
    `size` float NOT NULL,
    `description` text NOT NULL,
    `photo_link` varchar(100) DEFAULT NULL,
    `price` float DEFAULT NULL,
    `status` tinyint(1) DEFAULT '1',
    `category_id` int DEFAULT NULL,
    PRIMARY KEY (`product_id`)
);

INSERT INTO `products` (`product_id`, `title`, `weight`, `size`, `description`, `photo_link`, `price`, `status`, `category_id`) VALUES (1, 'Coca Cola 0.5', 0.5, 0.5, 'The Coca-Cola Company — американская пищевая компания, крупнейший мировой производитель и поставщик концентратов, сиропов и безалкогольных напитков.', 'Coca-Cola-Enjoy-Beverage-Industry.jpg', 59.99, 1, 1);
INSERT INTO `products` (`product_id`, `title`, `weight`, `size`, `description`, `photo_link`, `price`, `status`, `category_id`) VALUES (2, 'Sprite 1.5', 1.5, 1.5, 'Sprite — газированный безалкогольный напиток, со вкусом лайма и лимона, принадлежащий американской компании The Coca-Cola Company.', 'Sprite-Lymonade.jpg', 89.99, 0, 1);
INSERT INTO `products` (`product_id`, `title`, `weight`, `size`, `description`, `photo_link`, `price`, `status`, `category_id`) VALUES (3, 'Горький шоколад «OZera»', 0.09, 0.09, 'Горький шоколад без сахара и с пребиотиками – идеальное лакомство для сторонников здорового и сбалансированного питания.', 'ozera.jpg', 99.99, 1, 3);
