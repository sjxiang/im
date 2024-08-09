

-- 好友关系表;
CREATE TABLE `friend` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
    `friend_uid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '好友用户id',
    `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '好友验证',
    `add_source` tinyint NOT NULL DEFAULT '1' COMMENT '添加来源',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='好友关系表';


-- 好友申请表，女神 vs. 舔狗
CREATE TABLE `friend_apply` (
    `id`            int(11) UNSIGNED  NOT NULL AUTO_INCREMENT        COMMENT '自增id',
    `user_id`       varchar(64)       NOT NULL DEFAULT ''            COMMENT '用户id',
    `apply_uid`     varchar(64)       NOT NULL DEFAULT ''            COMMENT '申请用户id',
    `apply_msg`     varchar(255)      NOT NULL DEFAULT ''            COMMENT '申请信息',
    `apply_at`      timestamp         NOT NULL                       COMMENT '申请时间',
    `handle_result` tinyint           NOT NULL DEFAULT '0'           COMMENT '处理结果',
    `handle_msg`    varchar(255)      NOT NULL DEFAULT ''            COMMENT '处理回复',
    `handle_at`     timestamp         NULL     DEFAULT NULL          COMMENT '添加时间', 
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='好友申请表';



-- 群信息表
CREATE TABLE `group` (
    `id` varchar(24) NOT NULL DEFAULT '',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '群名',
    `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '群图标',
    `status`  tinyint DEFAULT NULL COMMENT '是否',
    `creator_uid` varchar(64) NOT NULL COMMENT '创建群组的用户id',
    `group_type` int(11) NOT NULL,
    `is_verify` boolean NOT NULL COMMENT '是否开启验证',
    `notification` varchar(255) DEFAULT NULL COMMENT '群公告',
    `notification_uid` varchar(64) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群信息表';


-- 群与用户关联表
CREATE TABLE `group_member` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `group_id` varchar(64) NOT NULL,
    `user_id` varchar(64) NOT NULL,
    `role_level` tinyint NOT NULL COMMENT '群聊用户等级',
    `join_at` timestamp NULL DEFAULT NULL,
    `join_source` tinyint DEFAULT NULL,
    `inviter_uid` varchar(64) DEFAULT NULL COMMENT '邀请人',
    `operator_uid` varchar(64) DEFAULT NULL COMMENT '审核员',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群与用户关联表';


-- 群申请表
CREATE TABLE `group_apply` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `apply_id` varchar(64) NOT NULL,
    `group_id` varchar(64) NOT NULL,
    `apply_msg` varchar(255) DEFAULT NULL,
    `apply_at` timestamp NULL DEFAULT NULL,
    `join_source`  tinyint DEFAULT NULL,
    `inviter_user_id` varchar(64) DEFAULT NULL,
    `handle_user_id` varchar(64) DEFAULT NULL,
    `handle_at` timestamp NULL DEFAULT NULL,
    `handle_result` tinyint DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群申请表';

-- 脚本生成
-- goctl model mysql ddl -src *.sql -dir . -cache=false --style=go_zero