-- db.db_user definition

CREATE TABLE `db_user` (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                           `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `isDeleted` tinyint NOT NULL DEFAULT '0' COMMENT '软删除 0-未删除 1-已删除',
                           `userAccount` varchar(128) DEFAULT NULL COMMENT '账户名',
                           `userPassword` varchar(128) DEFAULT NULL COMMENT '密码',
                           `nickname` varchar(256) DEFAULT 'gogogo' COMMENT '用户昵称',
                           `userRole` varchar(128) NOT NULL DEFAULT '普通用户' COMMENT '用户角色',
                           `avatarUrl` varchar(1024) DEFAULT 'https://pkg.go.dev/static/shared/icon/favicon.ico',
                           `phoneNumber` varchar(128) DEFAULT NULL COMMENT '电话',
                           `email` varchar(512) DEFAULT NULL COMMENT '邮箱地址',
                           `status` int NOT NULL DEFAULT '1' COMMENT '账户状态 0-正常 1-冻结',
                           PRIMARY KEY (`id`),
                           KEY `idx_userAccount` (`userAccount`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;