CREATE TABLE users
(
    `id`                bigint(20)              NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`              varchar(64) UNIQUE      NOT NULL COMMENT '用户名称',
    `password`          varchar(200)            NOT NULL,
    `follow_count`      int                     NOT NULL DEFAULT 0 comment '关注总数',
    `follower_count`    int                     NOT NULL DEFAULT 0 comment '粉丝总数',
    `is_follow`         tinyint(1)              NOT NULL DEFAULT 0 comment '是否关注',
    `create_time`       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time`       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `salt`              varchar(255)            NOT NULL COMMENT '加密用的salt',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE videos
(
    `id`                bigint(20)              NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`           bigint(20)              NOT NULL  COMMENT '用户id',
#     `name`              varchar(64) UNIQUE      NOT NULL COMMENT '用户名称',
#     `follow_count`      int                     NOT NULL DEFAULT 0 COMMENT '关注总数',
#     `follower_count`    int                     NOT NULL DEFAULT 0 COMMENT '粉丝总数',
#     `is_follow`         tinyint(1)              NOT NULL DEFAULT 0 COMMENT '是否关注',
    `play_url`          varchar(255)            NOT NULL,
    `cover_url`         varchar(255)            NOT NULL,
    `favorite_count`    int                     NOT NULL DEFAULT 0 COMMENT '',
    `comment_count`     INT                     NOT NULL DEFAULT 0 COMMENT '',
    `is_favorite`       tinyint(1)              not null default 0 COMMENT '是否点赞',
    `title`             varchar(255)            NOT NULL DEFAULT '' COMMENT '视频标题',
    PRIMARY KEY (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4;

CREATE TABLE comments
(
    `id`                bigint(20)              NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`           bigint(20)              NOT NULL  COMMENT '用户id',
    `name`              varchar(64) UNIQUE      NOT NULL  collate utf8mb4_general_ci  COMMENT '用户名称',
    `follow_count`      int                     NOT NULL DEFAULT 0 comment '关注总数',
    `follower_count`    int                     NOT NULL DEFAULT 0 comment '粉丝总数',
    `is_follow`         tinyint(1)              NOT NULL DEFAULT 0 comment '是否关注',
    `content`           mediumtext              NOT NULL,
    `create_date`       timestamp               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4;