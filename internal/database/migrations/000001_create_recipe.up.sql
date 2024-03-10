create table "recipe" (
	"id"         varchar   not null ,
	"created_at" timestamp not null ,
	"updated_at" timestamp not null ,
	"content"    blob      not null ,

	primary key ( "id" )
) ;

create table "recipe_metadata" (
	"recipe_id"  varchar   not null ,
	"title"      varchar   not null ,

	primary key ( "recipe_id" ) ,
	foreign key ( "recipe_id" ) references "recipe" ( "id" )
		on update restrict
		on delete cascade
) ;
