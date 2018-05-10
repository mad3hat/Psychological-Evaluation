/*
drop table users;
drop table managers;
drop table article;
drop table posts;
drop table resource;
drop table questions;
*/

/*
生成用户数据表
*/
create table users (
    id int(20) auto_increment primary key,
    name varchar(255) not null unique,
    age  varchar(64),
    sex  varchar(64),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);

/*
生成管理员数据表
*/
create table managers (
    id int(20) auto_increment primary key,
    name varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);

/*
心理阅读文章表
*/
create table article (
    id int(20) auto_increment primary key,
    topic text,
    body text,
    created_at timestamp not null
);

/*
用户留言表
*/
create table posts (
    id int(20) auto_increment primary key,
    body text,
    respond text,
    user_id integer references users(id),
    created_at timestamp not null
);

/*
试题管理表
*/
create table questions (
    id int(20) auto_increment primary key,
    body text,
    a text,
    b text,
    c text,
    d text,
    created_at timestamp not null
);

/*
友情链接与热门推荐管理表
*/
create table resource (
    id int(20) auto_increment primary key,
    link text,
    url text,
    created_at timestamp not null
);
