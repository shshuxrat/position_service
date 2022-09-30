package service

import (
	"context"
	"position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"

	"github.com/jmoiron/sqlx"
)

type companyService struct {
	logger                                             logger.LoggerI
	storage                                            storage.StorageI
	position_service.UnimplementedCompanyServiceServer //bug fixer
}

func NewCompanyService(log logger.LoggerI, db *sqlx.DB) *companyService {
	return &companyService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (r *companyService) Create(c context.Context, req *position_service.CreateCompany) (*position_service.CompanyId, error) {
	id, err := r.storage.Company().Create(req)

	if err != nil {
		return nil, err
	}

	return &position_service.CompanyId{
		Id: id,
	}, nil
}

func (r *companyService) GetAll(c context.Context, req *position_service.GetAllCompanyRequest) (*position_service.GetAllCompanyResponse, error) {
	companys, err := r.storage.Company().GetAll(req)
	if err != nil {
		return nil, err
	}

	return companys, nil

}

func (r *companyService) GetById(c context.Context, req *position_service.CompanyId) (*position_service.Company, error) {
	company, err := r.storage.Company().GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (r *companyService) Update(c context.Context, req *position_service.Company) (*position_service.Company, error) {
	updatedCompany, err := r.storage.Company().Update(req)
	if err != nil {
		return nil, err
	}

	return updatedCompany, nil

}

func (r *companyService) Delete(c context.Context, req *position_service.CompanyId) (*position_service.AffectedRow, error) {
	affectedrows, err := r.storage.Company().Delete(req)
	if err != nil {
		return nil, err
	}

	return affectedrows, nil

}
