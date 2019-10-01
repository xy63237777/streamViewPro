create table `comments` (
    `id` varchar(255) primary key not null,
    `video_id` varchar(255),
    author_id int unsigned,
    content text,
    time datetime
)charset=utf8;