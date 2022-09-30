package repo

import "position_service/genproto/position_service"

type CompanyRepoI interface {
	Create(req *position_service.CreateCompany) (string, error)
	GetAll(req *position_service.GetAllCompanyRequest) (*position_service.GetAllCompanyResponse, error)
	GetById(id string) (*position_service.Company, error)
	Update(req *position_service.Company) (*position_service.Company, error)
	Delete(req *position_service.CompanyId) (*position_service.AffectedRow, error)
}
