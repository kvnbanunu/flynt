package database

// only pings database for now will change later
func (db *DB) CheckHealth() error {
	return db.Ping()
}
