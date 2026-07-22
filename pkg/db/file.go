package db

import "time"

type File struct {
	Sha512sum   string
	Filename    string
	StorePath   string
	FilePath    string
	LastCheck   time.Time
	IsValid     bool
	FileNameSet bool
}

type DB struct {
}

func Open(path string) (*DB, error) {
	return &DB{}, nil
}

func (db *DB) Set(store string) error {
	return nil
}
