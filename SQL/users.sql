create table `users` (
    `id` int unsigned primary key  auto_increment,
    `login_name` varchar(255) unique key,
    `pwd` text
)charset=utf8;