create table book
(
    id        bigint auto_increment
        primary key,
    name      varchar(128) not null,
    author_id bigint       not null
);

create table user
(
    id   bigint auto_increment
        primary key,
    name varchar(128) not null,
    age  int          not null
);