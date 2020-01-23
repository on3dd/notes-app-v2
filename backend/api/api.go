package api

import (
	"github.com/jinzhu/gorm"
)

type API struct {
	db *gorm.DB
	signingKey []byte
}

// New returns a new API instance
func New(db *gorm.DB, sk []byte) *API {
	return &API{
		db: db,
		signingKey: sk,
	}
}
