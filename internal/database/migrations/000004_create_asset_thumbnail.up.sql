create table "asset_thumbnail" (
	"asset_id"           varchar not null ,
	"thumbnail_asset_id" varchar not null ,
	"size"               varchar not null ,

	primary key ( "asset_id", "thumbnail_asset_id" ) ,
	foreign key ( "asset_id" ) references "asset" ( "id" )
		on update restrict
		on delete cascade ,
	foreign key ( "thumbnail_asset_id" ) references "asset" ( "id" )
		on update restrict
		on delete cascade
) ;

create unique index "asset_thumbnail_size_unique_idx"
	on "asset_thumbnail" ( "asset_id", "size" ) ;
