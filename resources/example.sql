/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50721
Source Host           : localhost:3306
Source Database       : example

Target Server Type    : MYSQL
Target Server Version : 50721
File Encoding         : 65001

Date: 2021-07-23 23:13:17
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(50) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '用户名',
  `nick_name` varchar(150) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(150) DEFAULT NULL COMMENT '头像',
  `password` varchar(100) DEFAULT NULL COMMENT '密码',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
  `create_time` bigint(50) DEFAULT NULL COMMENT '更新时间',
  `del_status` tinyint(4) DEFAULT '0' COMMENT '是否删除 -1：已删除   0：正常',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('4', 'admin', '', '', '123456', '924931408@qq.com', '15751004355', '0', '0');
INSERT INTO `sys_user` VALUES ('7', 'admin1', 'admin', '', '123456', '924931408@qq.com', '15751004355', '0', '0');
INSERT INTO `sys_user` VALUES ('8', 'admin13', 'admin', '', '123456', '924931408@qq.com', '15751004355', '0', '0');
INSERT INTO `sys_user` VALUES ('10', 'admin1e3', 'admin', '', '123456', '924931408@qq.com', '15751004355', '0', '0');
