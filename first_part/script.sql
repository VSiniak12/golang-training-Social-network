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
    state_id   serial  not null
        constraint states_pk
            primary key,
    name       varchar not null,
    country_id integer not null
        constraint states_countries_country_id_fk
            references social_network.countries
);

alter table social_network.states
    owner to postgres;

create unique index states_state_id_uindex
    on social_network.states (state_id);

--create cities
create table social_network.cities
(
    city_id  serial  not null
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

create unique index groups_id_group_uindex
    on social_network.groups (id_group);

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
        constraint "user_cities _city_id_fk"
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

create unique index post_id_post_uindex
    on social_network.posts (id_post);

---create messages
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

create unique index message_id_message_uindex
    on social_network.messages (id_message);

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

create unique index comments_posts_comment_id_uindex
    on social_network.comments_posts (comment_id);

--create lu_members_group
create table social_network.lu_members_group
(
    group_id integer not null
        constraint lu_members_group_groups_id_group_fk
            references social_network.groups,
    user_id  integer not null
        constraint lu_members_group_user_id_user_fk
            references social_network.users
);

alter table social_network.lu_members_group
    owner to postgres;

--create lu_friends
create table social_network.lu_friends
(
    profile integer not null
        constraint lu_friends_user_id_user_fk
            references social_network.users,
    friend  integer not null
        constraint lu_friends_user_id_user_fk_2
            references social_network.users
);

alter table social_network.lu_friends
    owner to postgres;

--create lu_administrator_group
create table social_network.lu_administrator_group
(
    group_id         integer not null
        constraint lu_administrator_group_groups_id_group_fk
            references social_network.groups,
    administrator_id integer not null
        constraint lu_administrator_group_user_id_user_fk
            references social_network.users
);

alter table social_network.lu_administrator_group
    owner to postgres;

--Insert in tables

--Insert countires
INSERT INTO social_network.countries (name) VALUES ('Belarus');
INSERT INTO social_network.countries (name) VALUES ('Russia');
INSERT INTO social_network.countries (name) VALUES ('Ukraina');
INSERT INTO social_network.countries (name) VALUES ('Czech');
INSERT INTO social_network.countries (name) VALUES ('Poland');

--Insert states
INSERT INTO social_network.states (name, country_id) VALUES ('Gomelskaia', 1);
INSERT INTO social_network.states (name, country_id) VALUES ('Minskaaia', 1);
INSERT INTO social_network.states (name, country_id) VALUES ('Moskovskaia', 2);
INSERT INTO social_network.states (name, country_id) VALUES ('Kievskaia', 3);
INSERT INTO social_network.states (name, country_id) VALUES ('Xarkovskaia', 3);

--Insert cities
INSERT INTO social_network.cities (name, state_id) VALUES ('Gomel', 1);
INSERT INTO social_network.cities (name, state_id) VALUES ('Mozyr', 1);
INSERT INTO social_network.cities (name, state_id) VALUES ('Minsk', 2);
INSERT INTO social_network.cities (name, state_id) VALUES ('Moskva', 3);
INSERT INTO social_network.cities (name, state_id) VALUES ('Kiev', 4);

--Insert groups
INSERT INTO social_network.groups (name, date_create) VALUES ('epam', '2021-04-08');
INSERT INTO social_network.groups (name, date_create) VALUES ('Max Korzh', '2020-02-03');
INSERT INTO social_network.groups (name, date_create) VALUES ('Footbal', '2021-01-03');
INSERT INTO social_network.groups (name, date_create) VALUES ('News', '2021-03-21');
INSERT INTO social_network.groups (name, date_create) VALUES ('MDK', '2021-01-30');

--Insert users
INSERT INTO social_network.users(login, password, gender, email, last_name, first_name, birthday, city_id)
VALUES ('login1', '1234', true, 'e1@mail.ru', 'Ivanov', 'Ivan', '1997-03-01', 1);
INSERT INTO social_network.users(login, password, gender, email, last_name, first_name, birthday, city_id)
VALUES ('login2', '1234', true, 'e2@mail.ru', 'Sidorov', 'Andrei', '2000-12-03', 1);
INSERT INTO social_network.users(login, password, gender, email, last_name, first_name, birthday, city_id)
VALUES ('login3', '1234', false, 'e3@mail.ru', 'Sokolova', 'Svetlana', '1990-05-01', 3);
INSERT INTO social_network.users(login, password, gender, email, last_name, first_name, birthday, city_id)
VALUES ('login4', '1234', false, 'e5@mail.ru', 'Orexova', 'Anna', '2002-01-22', 5);
INSERT INTO social_network.users(login, password, gender, email, last_name, first_name, birthday, city_id)
VALUES ('login5', '1234', true, 'e5@mail.ru', 'Bulkin', 'Slava', '1980-06-03', 4);

--Insert posts
INSERT INTO social_network.posts (text, date_post, user_id) VALUES ('My holiday', '2021-04-12', 1);
INSERT INTO social_network.posts (text, date_post, user_id) VALUES ('My heart is broken', '2021-04-10', 3);
INSERT INTO social_network.posts (text, date_post, user_id) VALUES ('My mood...', '2021-02-01', 2);
INSERT INTO social_network.posts (text, date_post, user_id) VALUES ('Hello, friends!', '2021-03-22', 1);
INSERT INTO social_network.posts (text, date_post, user_id) VALUES ('I will get offer from epam', '2021-04-09', 4);

--Insert messages
INSERT INTO social_network.messages (text, from_user_id, to_user_id) VALUES ('Hello, how are you?', 1, 2);
INSERT INTO social_network.messages (text, from_user_id, to_user_id) VALUES ('Hello, i am pretty good, and you?', 2, 1);
INSERT INTO social_network.messages (text, from_user_id, to_user_id) VALUES ('Me too', 1, 2);
INSERT INTO social_network.messages (text, from_user_id, to_user_id) VALUES ('Call me, please', 3, 4);
INSERT INTO social_network.messages (text, from_user_id, to_user_id) VALUES ('When do you need an answer?', 4, 1);

--Insert comments_posts
INSERT INTO social_network.comments_posts (text, date_write, user_write, post_id)
VALUES ('Why??', '2021-04-12', 4, 2);
INSERT INTO social_network.comments_posts (text, date_write, user_write, post_id)
VALUES ('is best?', '2021-02-12', 5, 3);
INSERT INTO social_network.comments_posts (text, date_write, user_write, post_id)
VALUES ('Hello, how are you?', '2021-04-12', 2, 4);
INSERT INTO social_network.comments_posts (text, date_write, user_write, post_id)
VALUES ('Yes, of course', '2021-04-12', 2, 5);
INSERT INTO social_network.comments_posts (text, date_write, user_write, post_id)
VALUES ('You are hard-working', '2021-04-10', 3, 5);

--Insert lu_members_group
INSERT INTO social_network.lu_members_group VALUES (1,2);
INSERT INTO social_network.lu_members_group VALUES (1,5);
INSERT INTO social_network.lu_members_group VALUES (3,2);
INSERT INTO social_network.lu_members_group VALUES (2,1);
INSERT INTO social_network.lu_members_group VALUES (5,3);

--Insert lu_friends
INSERT INTO social_network.lu_friends VALUES (1,2);
INSERT INTO social_network.lu_friends VALUES (1,5);
INSERT INTO social_network.lu_friends VALUES (3,1);
INSERT INTO social_network.lu_friends VALUES (5,2);
INSERT INTO social_network.lu_friends VALUES (4,3);

--Insert lu_administrator_group
INSERT INTO social_network.lu_administrator_group VALUES (1,2);
INSERT INTO social_network.lu_administrator_group VALUES (2,3);
INSERT INTO social_network.lu_administrator_group VALUES (3,4);
INSERT INTO social_network.lu_administrator_group VALUES (4,1);
INSERT INTO social_network.lu_administrator_group VALUES (5,5);

