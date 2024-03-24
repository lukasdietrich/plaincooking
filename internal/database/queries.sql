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
select "r".*, (
	select "a"."asset_id"
	from "recipe_asset" "a"
	where "a"."recipe_id" = "r"."recipe_id"
	order by "a"."asset_id" desc
	limit 1
)
from "recipe_metadata" "r"
order by "r"."title" asc ;

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

-- name: CreateAsset :one
insert into "asset" (
	"id" ,
	"created_at" ,
	"filename" ,
	"media_type"
) values (
	sqlc.arg(id) ,
	sqlc.arg(created_at) ,
	sqlc.arg(filename) ,
	sqlc.arg(media_type)
) returning * ;

-- name: CreateAssetChunk :exec
insert into "asset_chunk" (
	"id" ,
	"asset_id" ,
	"content"
)  values (
	sqlc.arg(id) ,
	sqlc.arg(asset_id) ,
	sqlc.arg(content)
) ;

-- name: ReadAsset :one
select "a".*, (
	select sum(octet_length("c"."content")) "total_size"
	from "asset_chunk" "c"
	where "c"."asset_id" = "a"."id"
)
from "asset" "a"
where "a"."id" = sqlc.arg(id) ;

-- name: ReadAssetChunk :one
select *
from "asset_chunk"
where "asset_id" = sqlc.arg(asset_id)
  and (sqlc.arg(id_offset) is null or "id" > sqlc.arg(id_offset))
order by "id"
limit 1 ;

-- name: CreateRecipeAsset :exec
insert into "recipe_asset" (
	"recipe_id" ,
	"asset_id"
) values (
	sqlc.arg(recipe_id) ,
	sqlc.arg(asset_id)
) ;

-- name: ListRecipeAssets :many
select "a".*
from "asset" "a"
	inner join "recipe_asset" "r" on "a"."id" = "r"."asset_id"
where "r"."recipe_id" = sqlc.arg(recipe_id)
order by "a"."id" desc ;

-- name: CreateAssetThumbnail :exec
insert into "asset_thumbnail" (
	"asset_id" ,
	"thumbnail_asset_id" ,
	"size"
) values (
	sqlc.arg(asset_id) ,
	sqlc.arg(thumbnail_asset_id) ,
	sqlc.arg(size)
) ;

-- name: ReadAssetThumbnail :one
select *
from "asset_thumbnail"
where "asset_id" = sqlc.arg(asset_id)
  and "size" = sqlc.arg(size) ;
