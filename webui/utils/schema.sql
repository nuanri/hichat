CREATE TABLE IF NOT EXISTS `auth_user` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` INTEGER NOT NULL,
    `username` VARCHAR(128) NULL,
    `password` VARCHAR(512) NULL,
    `email` VARCHAR(128) NULL,
    `online` BOOLEAN NULL CHECK (online IN (0,1)),
    `last_activity_time` datetime DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `auth_session` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `sid` VARCHAR(128) NULL,
    `user_id` INTEGER NOT NULL
);
