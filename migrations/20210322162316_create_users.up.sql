CREATE TABLE IF NOT EXISTS `users` (
    `user_id` int NOT NULL AUTO_INCREMENT,
    `login` varchar(50) NOT NULL,
    `password` varchar(100) NOT NULL,
    `phone` bigint NOT NULL,
    `name` varchar(150) NOT NULL,
    `access` tinyint(1) DEFAULT '0',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `login` (`login`),
    UNIQUE KEY `phone` (`phone`)
);

-- login: admin
-- password: admin000
INSERT INTO `users` (`login`,`password`,`phone`,`name`,`access`) VALUES (
    'admin',
    '$2a$04$tYq8VY53DYUJNes6BG3jC.W8pLxGRPbkJzcpVglbwZoaHFdoobYGO',
    '88000000000',
    'Администратор',
    1
)

