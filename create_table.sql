CREATE TABLE `user` (
   `uid` BIGINT NOT NULL AUTO_INCREMENT COMMIT '主键',
   `name` VARCHAR(255) NOT NULL DEFAULT '' COMMIT '用户名',
   `sequence` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMIT '序列号',
   `group_id` BIGINT NOT NULL DEFAULT 0 COMMIT '组别主键',

   PRIMARY KEY (`uid`),
   INDEX `user_group` (`group_id`) -- 二级索引为 uid
) COMMIT '用户表';

CREATE TABLE `group` (
    `gid` BIGINT NOT NULL AUTO_INCREMENT COMMIT '主键',
    `max_sequence` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMIT '组别最大序列号'

    PRIMARY KEY (`gid`)
) COMMIT '用户组别表';

CREATE TABLE `section`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMIT '主键',
    `number` INT(11) NOT NULL DEFAULT 0 COMMIT '段号',

    -- use INET_ATON or INET_NTOA
    `ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMIT '注册段号的服务器IP',
    `bind_date` BIGINT NOT NULL DEFAULT 0 COMMIT '绑定时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY uk_section_number (`number`),
    INDEX KEY idx_section_machine (`number`, `ip`, `bind_date`)
) COMMIT '段号';