package system

func Migrations() error {
	var err error
	_, err = Db.Exec(`
         create table if not exists anime (
             id integer primary key not null,
             created datetime not null default current_timestamp,
             updated datetime not null default current_timestamp,
             title text not null,
             image text not null
         );
		 create table if not exists favorites (
			created datetime not null default current_timestamp,
			user_id text not null,
			anime_id integer not null,
			primary key (user_id, anime_id),
			foreign key (anime_id) references anime(id) on delete cascade on update cascade
		);
         create index if not exists idx_favorites_user_id on favorites(user_id);
    `)
	if err != nil {
		return err
	}

	return nil
}
