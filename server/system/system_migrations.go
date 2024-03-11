package system

func Migrations() error {
	var err error
	_, err = Db.Exec(`
         create table if not exists favorites (
             id text primary key not null,
             created datetime not null default current_timestamp,
             updated datetime not null default current_timestamp
         )`)
	if err != nil {
		return err
	}
	return nil
}
