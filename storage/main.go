package storage

import (
	"position_service/storage/postgres"
	"position_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Profession() repo.ProfessionRepoI
	Company() repo.CompanyRepoI
	Attribute() repo.AttributeRepoI
}

type storagePG struct {
	profession repo.ProfessionRepoI
	company    repo.CompanyRepoI
	attribute  repo.AttributeRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		profession: postgres.NewProfessionRepo(db),
		company:    postgres.NewCompanyRepo(db),
		attribute:  postgres.NewAttributeRepo(db),
	}

}

func (s *storagePG) Profession() repo.ProfessionRepoI {
	return s.profession
}

func (s *storagePG) Company() repo.CompanyRepoI {
	return s.company
}

func (s *storagePG) Attribute() repo.AttributeRepoI {
	return s.attribute
}
