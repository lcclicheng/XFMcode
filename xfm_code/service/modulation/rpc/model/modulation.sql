CREATE TABLE `orderDetails`
(
    `PayStatus`       varchar(255) default '',
    `PayDate`         DATE         DEFAULT NULL,
    `PayTime`         DATETIME     DEFAULT NULL,
    `TotalFee`        varchar(255) default '',
    `PayCouponFee`    varchar(255) default '',
    `PayOutTradeNo`   varchar(255) default '',
    `PayErrDesc`      varchar(255) default '',
    `Uid`             varchar(255) default '',
    `PayType`         varchar(255) default '',
    `PayTypeTradeNo`  varchar(255) default '',
    `OutRequestNo`    varchar(255) default '',
    `DimensionalCode` varchar(255) default '',
    `BarCode`         varchar(255) default '',
    INDEX             IDX1 (`Uid`,`PayTime`),
    PRIMARY KEY (`Uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;