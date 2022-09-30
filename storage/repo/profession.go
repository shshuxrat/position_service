package repo

import "position_service/genproto/position_service"

type ProfessionRepoI interface {
	Create(req *position_service.CreateProfession) (string, error)
	GetAll(req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error)
	GetById(id string) (*position_service.Profession, error)
	Update(req *position_service.Profession) (*position_service.Profession, error)
	Delete(req *position_service.ProfessionId) (*position_service.AffectedRows, error)
}
