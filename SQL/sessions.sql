create table `sessions` (
    `session_id` varchar(255) primary key not null,
    `TTL` tinytext,
    `login_name` varchar(255)
) charset=utf8;