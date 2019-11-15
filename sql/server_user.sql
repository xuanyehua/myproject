/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50727
Source Host           : localhost:3306
Source Database       : server_user

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2019-11-15 11:09:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rabc_group
-- ----------------------------
DROP TABLE IF EXISTS `rabc_group`;
CREATE TABLE `rabc_group` (
  `g_id` int(11) NOT NULL DEFAULT '0' COMMENT '组id',
  `group_name` varchar(64) NOT NULL DEFAULT '' COMMENT '组名',
  `parent_g_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `desc` varchar(128) NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`g_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rabc_user
-- ----------------------------
DROP TABLE IF EXISTS `rabc_user`;
CREATE TABLE `rabc_user` (
  `u_id` int(11) NOT NULL COMMENT '用户id',
  `g_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组id',
  `login_name` varchar(64) NOT NULL DEFAULT '' COMMENT '登录名',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `last_login_time` datetime NOT NULL,
  `login_count` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  PRIMARY KEY (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rbac_group_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_group_role`;
CREATE TABLE `rbac_group_role` (
  `gr_id` int(11) NOT NULL,
  `g_id` int(11) NOT NULL DEFAULT '0' COMMENT '组id',
  `r_id` int(11) DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`gr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rbac_log
-- ----------------------------
DROP TABLE IF EXISTS `rbac_log`;
CREATE TABLE `rbac_log` (
  `log_id` int(11) NOT NULL,
  `content` varchar(128) NOT NULL DEFAULT '' COMMENT '操作内容',
  `u_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人',
  `create_time` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rbac_right
-- ----------------------------
DROP TABLE IF EXISTS `rbac_right`;
CREATE TABLE `rbac_right` (
  `ri_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限id',
  `parent_ri_id` int(11) NOT NULL DEFAULT '0' COMMENT '父权限',
  `right_name` varchar(64) NOT NULL DEFAULT '' COMMENT '权限名称',
  `right_desc` varchar(128) NOT NULL DEFAULT '' COMMENT '权限描述',
  PRIMARY KEY (`ri_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role`;
CREATE TABLE `rbac_role` (
  `r_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `parent_r_id` int(11) NOT NULL DEFAULT '0' COMMENT '父角色id',
  `role_name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL,
  `desc` varchar(128) NOT NULL DEFAULT '' COMMENT '角色描述',
  PRIMARY KEY (`r_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rbac_role_right
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_right`;
CREATE TABLE `rbac_role_right` (
  `rr_id` int(11) NOT NULL DEFAULT '0',
  `r_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `ri_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限id',
  PRIMARY KEY (`rr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
