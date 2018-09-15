USE gdisid;

CREATE TABLE `user` (
   `uid` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
   `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户名',
   `sequence` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '序列号',
   `group_id` BIGINT NOT NULL DEFAULT 0 COMMENT '组别主键',

   PRIMARY KEY (`uid`),
   INDEX `user_group` (`group_id`) -- 二级索引为 uid
) ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户表';

CREATE TABLE `user_group` (
    `group_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `max_sequence` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '组别最大序列号',

    PRIMARY KEY (`group_id`)
) ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户组别表';

CREATE TABLE `section` (
    `section_id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `group_id_left_interval` INT(11) NOT NULL DEFAULT 0 COMMENT '组别段号左区间',
    `group_id_right_interval` INT(11) NOT NULL DEFAULT 0 COMMENT '组别段号右区间',

    -- use INET_ATON or INET_NTOA
    `is_bind` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已被绑定',
    `ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册段号的服务器IP',
    `bind_date` BIGINT NOT NULL DEFAULT 0 COMMENT '绑定时间',

    PRIMARY KEY (`section_id`),
    INDEX idx_group_interval (`group_id_left_interval`, `group_id_right_interval`),
    INDEX  idx_section_machine (`ip`, `bind_date`)
) ENGINE InnoDB CHARSET utf8mb4 COMMENT '段号';
