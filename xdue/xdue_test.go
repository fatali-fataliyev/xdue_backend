package xdue

import "github.com/fatali-fataliyev/xdue_backend/db"

type XDue struct {
	repo db.StorageRepository
}

func NewXDue(r db.StorageRepository) *XDue {
	return &XDue{
		repo: r,
	}
}

// func (x *XDue) CreateUser(u api.User){

// }
