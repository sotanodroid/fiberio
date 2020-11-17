package handler

import (
	repo "fiberio/app/repo"
)

// BaseHandler ...
type BaseHandler struct {
	db repo.Repository
}

// NewBasehandler ...
func NewBasehandler(db repo.Repository) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}
