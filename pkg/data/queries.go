package data

//users
const joinNameCity = "left join social_network.cities on social_network.users.city_id = social_network.cities.id_city"
const readAllUsersQuery = `SELECT users.id_user, users.login, users.password, users.gender,users.email, 
users.last_name, users.first_name, users.birthday, cities.name FROM "users" left join social_network.cities on social_network.users.city_id = social_network.cities.id_city `
const readAllUsersUpdateQuery = `SELECT * FROM "users" WHERE (id_user = $1)`
const insertUsersQuery = `INSERT INTO "users" ("id_user","login","password","gender","email","last_name",
"first_name", "birthday", "name") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "users".*`
const updateUsersQuery = `UPDATE "users" SET "login" = $1 WHERE "users"."id_user" = $2`
const deleteUsersQuery = `DELETE FROM "users" WHERE (id_user = $1)`
const readUsersQuery = `SELECT users.id_user, users.login, users.password, users.gender,users.email, users.last_name, users.first_name, users.birthday, cities.name FROM "users" left join social_network.cities on social_network.users.city_id = social_network.cities.id_city WHERE (id_user = $1)`

//states
const readAllStatesQuery = `SELECT * FROM "states"`
const insertStatesQuery = `INSERT INTO "states" ("id_state","name","country_id") VALUES ($1,$2,$3) RETURNING "states".*`
const updateStatesQuery = `UPDATE "states" SET "name" = $1 WHERE "states"."id_state" = $2`
const selectUpdateStatesQuery = `SELECT * FROM "states" WHERE (id_state = $1)`
const deleteStatesQuery = `DELETE FROM "states" WHERE (id_state = $1)`
