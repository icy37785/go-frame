package db

type Test struct {
	ID    uint64 `db:"id"`
	Title string `db:"title"`
}
