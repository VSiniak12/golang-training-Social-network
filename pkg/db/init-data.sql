
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

--Insert members_group
INSERT INTO social_network.members_group (group_id, user_id) VALUES (1,2);
INSERT INTO social_network.members_group (group_id, user_id) VALUES (1,5);
INSERT INTO social_network.members_group (group_id, user_id) VALUES (3,2);
INSERT INTO social_network.members_group (group_id, user_id) VALUES (2,1);
INSERT INTO social_network.members_group (group_id, user_id) VALUES (5,3);

--Insert friends
INSERT INTO social_network.friends (profile_id, friend_id) VALUES (1,2);
INSERT INTO social_network.friends (profile_id, friend_id) VALUES (1,5);
INSERT INTO social_network.friends (profile_id, friend_id) VALUES (3,1);
INSERT INTO social_network.friends (profile_id, friend_id) VALUES (5,2);
INSERT INTO social_network.friends (profile_id, friend_id) VALUES (4,3);

--Insert administrator_group
INSERT INTO social_network.administrator_group (group_id, administrator_id) VALUES (1,2);
INSERT INTO social_network.administrator_group (group_id, administrator_id) VALUES (2,3);
INSERT INTO social_network.administrator_group (group_id, administrator_id) VALUES (3,4);
INSERT INTO social_network.administrator_group (group_id, administrator_id) VALUES (4,1);
INSERT INTO social_network.administrator_group (group_id, administrator_id) VALUES (5,5);