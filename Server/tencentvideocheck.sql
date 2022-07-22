-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2022-07-22 12:06:26
-- 服务器版本： 10.4.24-MariaDB
-- PHP 版本： 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `tencentvideocheck`
--
CREATE DATABASE IF NOT EXISTS `tencentvideocheck` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `tencentvideocheck`;

-- --------------------------------------------------------

--
-- 表的结构 `user`
--
-- 创建时间： 2022-07-22 08:37:11
-- 最后更新： 2022-07-22 10:03:44
--

CREATE TABLE `user` (
  `Cookie` mediumtext NOT NULL COMMENT 'Cookie',
  `Barrage` int(11) DEFAULT NULL COMMENT '弹幕签到',
  `Check` int(11) DEFAULT NULL COMMENT '签到',
  `Download` int(11) DEFAULT NULL COMMENT '下载签到',
  `Giving` int(11) DEFAULT NULL COMMENT '赠送签到',
  `Watch` int(11) DEFAULT NULL COMMENT '观看60分钟签到',
  `Obtained` varchar(255) DEFAULT NULL COMMENT '本月获得V力值',
  `UserInfo` varchar(255) DEFAULT NULL COMMENT '用户信息',
  `Notice` varchar(255) DEFAULT NULL COMMENT '通知类型',
  `NoticeToken` varchar(255) DEFAULT NULL COMMENT '通知Token'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
