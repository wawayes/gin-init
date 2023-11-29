use db;

-- database.db_user definition

CREATE TABLE `db_user` (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                            `user_id` varchar(128) NOT NULL COMMENT '用户ID',
                           `user_account` varchar(128) DEFAULT NULL COMMENT '账户名',
                           `user_password` varchar(128) DEFAULT NULL COMMENT '密码',
                           `nickname` varchar(256) DEFAULT 'gogogo' COMMENT '用户昵称',
                           `user_role` varchar(128) NOT NULL DEFAULT 'user' COMMENT '用户角色',
                           `avatar_url` varchar(1024) DEFAULT 'https://pkg.go.dev/static/shared/icon/favicon.ico',
                           `phone_number` varchar(128) DEFAULT NULL COMMENT '电话',
                           `email` varchar(512) DEFAULT NULL COMMENT '邮箱地址',
                           `status` int NOT NULL DEFAULT '1' COMMENT '账户状态 0-正常 1-冻结',
                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` timestamp NULL COMMENT '删除时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_userAccount` (`user_account`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;