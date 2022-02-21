package service

import (
	"go_web_server/internal/repository"
)

type InitDB struct {
	repo repository.InitDBRepository
}

func NewInitDBService(repo repository.InitDBRepository) *InitDB {
	return &InitDB{repo: repo}
}

func (i InitDB) CreateTables() error {
	return i.repo.CreateTables()
}

func (i InitDB) InsertInto() error {
	return i.repo.InsertInto()
}
