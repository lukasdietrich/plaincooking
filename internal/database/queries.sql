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
from "recipes" r
order by r.title asc ;
