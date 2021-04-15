--Create SCHEMA query

--CREATE SCHEMA social_network;

--Create tables

create table countries
(
    country_id serial  not null
        constraint country_pk
            primary key,
    name       varchar not null
);

create table states
(
    id_state   serial  not null
        constraint states_pk
            primary key,
    name       varchar not null,
    country_id integer not null
        constraint states_countries_country_id_fk
            references countries
            on update cascade on delete cascade
);

create table cities
(
    id_city  serial  not null
        constraint city_pk
            primary key,
    name     varchar not null,
    state_id integer not null
        constraint "cities _states_state_id_fk"
            references states
            on update cascade on delete cascade
);

create table groups
(
    id_group    serial  not null
        constraint groups_pk
            primary key,
    name        varchar not null,
    date_create date    not null
);

create table users
(
    id_user    serial  not null
        constraint user_pk
            primary key,
    login      varchar,
    password   varchar,
    gender     boolean not null,
    email      varchar not null,
    last_name  varchar not null,
    first_name varchar not null,
    birthday   date    not null,
    city_id    integer
        constraint user_cities_city_id_fk
            references cities
            on update set null on delete set null
);

create table posts
(
    id_post   serial  not null
        constraint post_pk
            primary key,
    text      varchar not null,
    date_post date,
    user_id   integer not null
        constraint posts_user_id_user_fk
            references users
            on delete cascade
);

create table messages
(
    id_message   serial  not null
        constraint message_pk
            primary key,
    text         varchar not null,
    from_user_id integer
        constraint messages_user_id_user_fk
            references users
            on delete set null,
    to_user_id   integer
        constraint messages_user_id_user_fk_2
            references users
            on delete set null
);

create table comments_posts
(
    comment_id serial  not null
        constraint comments_posts_pk
            primary key,
    text       text    not null,
    date_write date    not null,
    user_write integer
        constraint comments_posts_user_id_user_fk
            references users
            on delete set null,
    post_id    integer not null
        constraint comments_posts_posts_id_post_fk
            references posts
            on delete cascade
);

create table members_group
(
    id       serial  not null
        constraint members_group_pk
            primary key,
    group_id integer not null
        constraint members_group_groups_id_group_fk
            references groups
            on delete cascade,
    user_id  integer not null
        constraint members_group_user_id_user_fk
            references users
            on delete cascade
);

create table friends
(
    id         serial  not null
        constraint friends_pk
            primary key,
    profile_id integer not null
        constraint friends_user_id_user_fk
            references users
            on delete cascade,
    friend_id  integer not null
        constraint friends_user_id_user_fk_2
            references users
            on delete cascade
);

create table administrator_group
(
    id               serial  not null
        constraint administrator_group_pk
            primary key,
    group_id         integer not null
        constraint administrator_group_groups_id_group_fk
            references groups
            on delete cascade,
    administrator_id integer not null
        constraint administrator_group_user_id_user_fk
            references users
            on delete cascade
);