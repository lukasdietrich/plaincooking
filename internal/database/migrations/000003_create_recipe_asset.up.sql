create table "recipe_asset" (
	"recipe_id" varchar not null ,
	"asset_id"  varchar not null ,

	primary key ( "recipe_id", "asset_id" ) ,
	foreign key ( "recipe_id" ) references "recipe" ( "id" )
		on update restrict
		on delete cascade ,
	foreign key ( "asset_id" ) references "asset" ( "id" )
		on update restrict
		on delete cascade
) ;
