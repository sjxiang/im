

-- 创建数据库
-- CREATE DATABASE IF NOT EXISTS `im` DEFAULT CHARACTER SET = 'utf8mb4';
    
-- 使用数据库
-- USE `im`;

-- 查看建表语句
-- SHOW CREATE TABLE `users`;

-- 删除表
-- DROP TABLE IF EXISTS `users`;


CREATE TABLE `users` (
    `id`           bigint(20)    NOT NULL AUTO_INCREMENT            COMMENT '自增id',
    `nickname`     varchar(64)   NOT NULL DEFAULT ''                COMMENT '昵称',
    `mobile`       varchar(64)   NOT NULL DEFAULT ''                COMMENT '手机号码',
    `password`     varchar(64)            DEFAULT NULL              COMMENT '密码',
    `avatar`       varchar(1024) NOT NULL DEFAULT ''                COMMENT '头像',
    `sex`          tinyint       NOT NULL DEFAULT '1'               COMMENT '性别',
    `status`       tinyint       NOT NULL DEFAULT '0'               COMMENT '是否锁住',
    `intro`        varchar(1024) NOT NULL DEFAULT ''                COMMENT '自我介绍',
    `created_at`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `updated_at`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- goctl model mysql ddl -src *.sql -dir . -cache=true --style=go_zero