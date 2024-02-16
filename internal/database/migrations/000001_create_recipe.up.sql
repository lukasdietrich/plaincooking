create table "recipes" (
	"id"         varchar   not null ,
	"created_at" timestamp not null ,
	"updated_at" timestamp not null ,

	"slug"       varchar   not null ,
	"title"      varchar   not null ,

	primary key ( "id" )
) ;
