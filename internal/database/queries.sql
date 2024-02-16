-- name: CreateRecipe :one
insert into "recipes" (
	"id" ,
	"created_at" ,
	"updated_at" ,
	"slug" ,
	"title"
) values (
	sqlc.arg(id) ,
	sqlc.arg(created_at) ,
	sqlc.arg(created_at) ,
	sqlc.arg(slug) ,
	sqlc.arg(title)
) returning * ;

-- name: ListRecipes :many
select *
from "recipes" "r"
order by "r"."title" asc ;

-- name: ReadRecipe :one
select *
from "recipes" "r"
where "r"."id" = sqlc.arg(id) ;

-- name: UpdateRecipe :one
update "recipes"
set "updated_at" = sqlc.arg(updated_at) ,
	"slug" = sqlc.arg(slug) ,
	"title" = sqlc.arg(title)
where "id" = sqlc.arg(id)
returning * ;

-- name: DeleteRecipe :one
delete from "recipes"
where "id" = sqlc.arg(id)
returning * ;
