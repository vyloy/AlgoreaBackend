package database

import (
	t "github.com/France-ioi/AlgoreaBackend/app/types"
)

// ItemStringStore implements database operations on `items_strings`
type ItemStringStore struct {
	*DataStore
}

// ItemString matches the content the `items_strings` table
type ItemString struct {
	ID         t.Int64  `sql:"column:ID"`
	ItemID     t.Int64  `sql:"column:idItem"`
	LanguageID t.Int64  `sql:"column:idLanguage"`
	Title      t.String `sql:"column:sTitle"`
	Version    int64    `sql:"column:iVersion"` // use Go default in DB (to be fixed)
}

func (s *ItemStringStore) tableName() string {
	return "items_strings"
}

// Insert does a INSERT query in the given table with data that may contain types.* types
func (s *ItemStringStore) Insert(data *ItemString) error {
	return s.db.insert(s.tableName(), data)
}

// All creates a composable query without filtering
func (s *ItemStringStore) All() *DB {
	return &DB{s.db.Table(s.tableName())}
}
