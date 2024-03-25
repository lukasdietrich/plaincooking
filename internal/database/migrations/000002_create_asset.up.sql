create table "asset" (
	"id"         varchar   not null ,
	"created_at" timestamp not null ,
	"filename"   varchar   not null ,
	"media_type" varchar   not null ,

	primary key ( "id" )
) ;

create table "asset_chunk" (
	"id"       varchar not null ,
	"asset_id" varchar not null ,
	"content"  blob    not null ,

	primary key ( "id" ) ,
	foreign key ( "asset_id" ) references "asset" ( "id" )
		on update restrict
		on delete cascade
) ;
