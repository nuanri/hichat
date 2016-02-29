CREATE TABLE `auth_code` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `authcode_key` char(64) DEFAULT NULL,
  `authcode` char(6) DEFAULT NULL,
  `email` char(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

CREATE TABLE `auth_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` char(60) NOT NULL,
  `password` char(128) NOT NULL,
  `email` char(60) NOT NULL,
  `online` tinyint(1) NOT NULL,
  `last_msg_time` datetime DEFAULT NULL,
  `last_act_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ;

CREATE TABLE `auth_session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11),
  `sid` char(128) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (user_id) REFERENCES auth_user(id)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

CREATE TABLE `msg_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` char(20) NOT NULL,
  `msg` char(255) NOT NULL,
  `add_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=324 DEFAULT CHARSET=utf8;


