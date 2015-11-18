CREATE TABLE `auth_session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11),
  `sid` char(128) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (user_id) REFERENCES auth_user(id)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
