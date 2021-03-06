package database

// UserStore implements database operations on `users`
type UserStore struct {
	*DataStore
}

// GetByID populates `dest` with the user identified by userID
func (s *UserStore) GetByID(userID int64, dest interface{}) error {
	return s.ByID(userID).Scan(dest).Error
}

// ByID returns a composable query of users filtered by userID
func (s *UserStore) ByID(userID int64) *DB {
	return &DB{s.All().Where("ID = ?", userID)}
}

// All creates a composable query without filtering
func (s *UserStore) All() *DB {
	return &DB{s.db.Table("users")}
}
