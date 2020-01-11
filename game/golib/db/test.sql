CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nickName` varchar(50) NOT NULL DEFAULT '' ,
  `createTime` datetime DEFAULT '0000-00-00 00:00:00',
  `sex` char DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert into
INSERT INTO `user` (`nickname`,`createTime`,`sex`)VALUES('rs',NOW(),'0');
INSERT INTO `user` (`nickname`,`createTime`,`sex`)VALUES('bs',NOW(),'0');
INSERT INTO `user` (`nickname`,`createTime`,`sex`)VALUES('ds',NOW(),'0');
INSERT INTO `user` (`nickname`,`createTime`,`sex`)VALUES('ns',NOW(),'0');

SELECT `id`,`nickName`,`createTime` FROM `user`;


CREATE TABLE `stu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sex` char DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert into
INSERT INTO `stu` (`name`,`sex`)VALUES('rs','0');
INSERT INTO `stu` (`name`,`sex`)VALUES('bs','0');
INSERT INTO `stu` (`name`,`sex`)VALUES('ds','0');
INSERT INTO `stu` (`name`,`sex`)VALUES('ns','0');

SELECT `id`,`name`,`sex` FROM `stu`;
