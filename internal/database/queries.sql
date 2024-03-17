-- name: CreateRecipe :exec
insert into "recipe" (
	"id" ,
	"created_at" ,
	"updated_at" ,
	"content"
) values (
	sqlc.arg(id) ,
	sqlc.arg(created_at) ,
	sqlc.arg(created_at) ,
	sqlc.arg(content)
) ;

-- name: CreateRecipeMetadata :exec
insert into "recipe_metadata" (
	"recipe_id" ,
	"title"
) values (
	sqlc.arg(recipe_id) ,
	sqlc.arg(title)
) ;

-- name: ListRecipeMetadata :many
select *
from "recipe_metadata"
order by "title" asc ;

-- name: ReadRecipe :one
select *
from "recipe"
where "id" = sqlc.arg(id) ;

-- name: UpdateRecipe :one
update "recipe"
set "updated_at" = sqlc.arg(updated_at) ,
	"content" = sqlc.arg(content)
where "id" = sqlc.arg(id)
returning * ;

-- name: UpdateRecipeMetadata :one
update "recipe_metadata"
set "title" = sqlc.arg(title)
where "recipe_id" = sqlc.arg(recipe_id)
returning * ;

-- name: DeleteRecipe :one
delete from "recipe"
where "id" = sqlc.arg(id)
returning * ;
