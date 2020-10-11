package database

type Database struct {
}

func New(url string) (*Database, error) {
	db := Database{}
	return &db, nil
}
