package api

import (
	"github.com/jinzhu/gorm"
)

type API struct {
	db *gorm.DB
}

// New returns a new API instance
func New(db *gorm.DB) *API {
	return &API{
		db: db,
	}
}
