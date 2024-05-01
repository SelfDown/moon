-- backend_db.attendance_data definition

CREATE TABLE `attendance_data`
(
    `attendance_data_id`            varchar(100) NOT NULL,
    `attendance_day`                varchar(100) NOT NULL,
    `user_id`                       varchar(50)  NOT NULL,
    `attendance_type`               varchar(50)  NOT NULL,
    `attendance_morning_type`       varchar(100) DEFAULT NULL,
    `attendance_morning_time`       varchar(50)  DEFAULT NULL,
    `attendance_morning_location`   varchar(50)  DEFAULT NULL,
    `attendance_afternoon_type`     varchar(50)  DEFAULT NULL,
    `attendance_afternoon_time`     varchar(100) DEFAULT NULL,
    `attendance_afternoon_location` varchar(50)  DEFAULT NULL,
    `day_error_info`                varchar(100) DEFAULT NULL,
    `create_time`                   varchar(50)  DEFAULT NULL,
    PRIMARY KEY (`attendance_data_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;


-- backend_db.attendance_rule definition

CREATE TABLE `attendance_rule`
(
    `rule_id`       varchar(255) NOT NULL,
    `rule_group_id` varchar(50)  DEFAULT NULL,
    `name`          varchar(255) DEFAULT NULL,
    `rule_code`     varchar(255) DEFAULT NULL,
    `rule_type`     varchar(255) DEFAULT NULL,
    `rule_time`     varchar(255) DEFAULT NULL,
    PRIMARY KEY (`rule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;


-- backend_db.attendance_rule_group definition

CREATE TABLE `attendance_rule_group`
(
    `rule_group_id` varchar(50) NOT NULL,
    `name`          varchar(255) DEFAULT NULL,
    `weight`        int(11) DEFAULT NULL,
    `default`       varchar(100) DEFAULT NULL,
    PRIMARY KEY (`rule_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;


-- backend_db.attendance_type definition

CREATE TABLE `attendance_type`
(
    `attendance_type` varchar(50) NOT NULL,
    `name`            varchar(50) DEFAULT NULL,
    `short_name`      varchar(50) DEFAULT NULL,
    PRIMARY KEY (`attendance_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;


-- backend_db.check_in_data definition

CREATE TABLE `check_in_data`
(
    `check_in_id`      bigint(20) NOT NULL,
    `checkin_time`     bigint(20) DEFAULT NULL,
    `checkin_type`     varchar(50)  DEFAULT NULL,
    `deviceid`         varchar(50)  DEFAULT NULL,
    `exception_type`   varchar(100) DEFAULT NULL,
    `groupid`          bigint(20) DEFAULT NULL,
    `groupname`        varchar(50)  DEFAULT NULL,
    `lat`              bigint(20) DEFAULT NULL,
    `lng`              bigint(20) DEFAULT NULL,
    `location_detail`  varchar(100) DEFAULT NULL,
    `location_title`   varchar(100) DEFAULT NULL,
    `notes`            varchar(100) DEFAULT NULL,
    `sch_checkin_time` bigint(20) DEFAULT NULL,
    `timeline_id`      bigint(20) DEFAULT NULL,
    `userid`           varchar(100) DEFAULT NULL,
    `wifimac`          varchar(100) DEFAULT NULL,
    `wifiname`         varchar(100) DEFAULT NULL,
    `create_time`      varchar(100) DEFAULT NULL,
    PRIMARY KEY (`check_in_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;