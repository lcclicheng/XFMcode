CREATE TABLE `OrderDetails`
(
    `PayStatus`      varchar(255) default '',
    `PayDate`        varchar(255) default '',
    `PayTime`        varchar(255) default '',
    `TotalFee`       varchar(255) default '',
    `PayCouponFee`   varchar(255) default '',
    `PayOutTradeNo`  varchar(255) default '',
    `PayErrDesc`     varchar(255) default '',
    `Uid`            varchar(255) default '',
    `PayType`        varchar(255) default '',
    `PayTypeTradeNo` varchar(255) default '',
    INDEX            IDX1 (`Uid`,`PayTime`),
    PRIMARY KEY (`Uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;