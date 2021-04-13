--Create SCHEMA query

CREATE SCHEMA social_network;

--Create tables

--create countries
create table social_network.countries
(
    country_id serial  not null
        constraint country_pk
            primary key,
    name       varchar not null
);

alter table social_network.countries
    owner to postgres;

--create states
create table social_network.states
(
    id_state   serial  not null
        constraint states_pk
            primary key,
    name       varchar not null,
    country_id integer not null
        constraint states_countries_country_id_fk
            references social_network.countries
);

alter table social_network.states
    owner to postgres;

--create cities
create table social_network.cities
(
    id_city  serial  not null
        constraint city_pk
            primary key,
    name     varchar not null,
    state_id integer not null
        constraint "cities _states_state_id_fk"
            references social_network.states
);

alter table social_network.cities
    owner to postgres;

--create groups
create table social_network.groups
(
    id_group    serial  not null
        constraint groups_pk
            primary key,
    name        varchar not null,
    date_create date    not null
);

alter table social_network.groups
    owner to postgres;

--create users
create table social_network.users
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
    city_id    integer not null
        constraint user_cities_city_id_fk
            references social_network.cities
);

alter table social_network.users
    owner to postgres;

--create posts
create table social_network.posts
(
    id_post   serial  not null
        constraint post_pk
            primary key,
    text      varchar not null,
    date_post date,
    user_id   integer not null
        constraint posts_user_id_user_fk
            references social_network.users
);

alter table social_network.posts
    owner to postgres;

--create messages
create table social_network.messages
(
    id_message   serial  not null
        constraint message_pk
            primary key,
    text         varchar not null,
    from_user_id integer not null
        constraint messages_user_id_user_fk
            references social_network.users,
    to_user_id   integer not null
        constraint messages_user_id_user_fk_2
            references social_network.users
);

alter table social_network.messages
    owner to postgres;

--create comments_posts
create table social_network.comments_posts
(
    comment_id serial  not null
        constraint comments_posts_pk
            primary key,
    text       text    not null,
    date_write date    not null,
    user_write integer not null
        constraint comments_posts_user_id_user_fk
            references social_network.users,
    post_id    integer not null
        constraint comments_posts_posts_id_post_fk
            references social_network.posts
);

alter table social_network.comments_posts
    owner to postgres;

--create members_group
create table social_network.members_group
(
    id serial not null
        constraint members_group_pk
            primary key,
    group_id integer not null
        constraint members_group_groups_id_group_fk
            references social_network.groups,
    user_id  integer not null
        constraint members_group_user_id_user_fk
            references social_network.users
);

alter table social_network.members_group
    owner to postgres;

--create friends
create table social_network.friends
(
    id serial not null
        constraint friends_pk
            primary key,
    profile_id integer not null
        constraint friends_user_id_user_fk
            references social_network.users,
    friend_id  integer not null
        constraint friends_user_id_user_fk_2
            references social_network.users
);

alter table social_network.friends
    owner to postgres;

--create administrator_group
create table social_network.administrator_group
(
    id serial not null
        constraint administrator_group_pk
            primary key,
    group_id integer not null
        constraint administrator_group_groups_id_group_fk
            references social_network.groups,
    administrator_id integer not null
        constraint administrator_group_user_id_user_fk
            references social_network.users
);

alter table social_network.administrator_group
    owner to postgres;
