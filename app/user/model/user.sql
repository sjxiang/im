

-- 创建数据库
-- CREATE DATABASE IF NOT EXISTS `im` DEFAULT CHARACTER SET = 'utf8mb4';
    
-- 使用数据库
-- USE `im`;

-- 查看建表语句
-- SHOW CREATE TABLE `users`;

-- 删除表
-- DROP TABLE IF EXISTS `users`;

-- 脚本生成
-- goctl model mysql ddl -src *.sql -dir . -cache=false --style=go_zero

CREATE TABLE `user` (
    `id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'uuid',
    `avatar` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default.jpeg' COMMENT '头像',
    `nickname` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
    `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码',
    `password` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
    `status` tinyint COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '是否锁住，0 active、1 forbidden',
    `sex` tinyint COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '性别，0 man、1 woman',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';   