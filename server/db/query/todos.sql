-- name: ListTodos :many
select * from todos;

-- name: SetTodo :exec
insert into todos(name)values($1);