create table `video_info`(
    `id` varchar(255) primary key not null,
    `author_id` int unsigned,
    `name` text,
    `display_ctime` text,
    `create_time` datetime
)charset=utf8;