CREATE TABLE `wecart`
(
    `status` varchar(255) NOT NULL COMMENT 'status',
    `url` varchar(255) NOT NULL COMMENT 'original url',
    PRIMARY KEY(`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;