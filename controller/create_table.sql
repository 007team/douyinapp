CREATE DATABASE IF NOT EXISTS `douyin`;
USE `douyin`;
CREATE TABLE user
(
    `id`                bigint(20) unsigned     NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`              varchar(64) UNIQUE      NOT NULL  collate utf8mb4_general_ci NOT NULL COMMENT '用户名称',
    `follow_count`      integer                 NOT NULL DEFAULT 0 comment '关注总数',
    `follower_count`    integer                 NOT NULL DEFAULT 0 comment '粉丝总数',
    `is_follow`         boolean                 NOT NULL DEFAULT FALSE comment '是否关注',
    `user_id`           bigint(20)              NOT NULL COMMENT 'user_id',
    `password`          varchar(64)             NOT NULL collate utf8mb4_general_ci NOT NULL,
    `salt`              integer                 not null comment '加密 salt',
    `create_time`       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY 'idx_name' (`name`),
    UNIQUE KEY 'idx_user_id' (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

