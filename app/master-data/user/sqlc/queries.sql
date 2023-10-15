-- name: SelectAllUser :many
select * from master_user;
-- name: SelectOneUser :one
select * from master_user where username = $1;