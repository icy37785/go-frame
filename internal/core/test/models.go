package test

import (
	"github.com/icy37785/go-frame/internal/core/test/db"
	"unsafe"
)

type Test struct {
	ID    uint64 `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func toTest(dbSearch db.Test) Test {
	pu := (*Test)(unsafe.Pointer(&dbSearch))
	return *pu
}

func toTestSlice(dbSearches []db.Test) []Test {
	searchs := make([]Test, len(dbSearches))
	for i, dbSearch := range dbSearches {
		searchs[i] = toTest(dbSearch)
	}
	return searchs
}
