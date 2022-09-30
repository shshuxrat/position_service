package service

import (
	"context"
	"position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"

	"github.com/jmoiron/sqlx"
)

type professionService struct {
	logger                                                logger.LoggerI
	storage                                               storage.StorageI
	position_service.UnimplementedProfessionServiseServer //bug fixer
}

func NewProfessionService(log logger.LoggerI, db *sqlx.DB) *professionService {
	return &professionService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}
func (p *professionService) Create(c context.Context, req *position_service.CreateProfession) (*position_service.ProfessionId, error) {
	id, err := p.storage.Profession().Create(req)
	if err != nil {
		return nil, err
	}

	return &position_service.ProfessionId{
		Id: id,
	}, nil
}
func (p *professionService) GetAll(c context.Context, req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error) {
	professions, err := p.storage.Profession().GetAll(req)
	if err != nil {
		return nil, err
	}
	return professions, nil

}

func (p *professionService) GetById(c context.Context, req *position_service.ProfessionId) (*position_service.Profession, error) {
	profession, err := p.storage.Profession().GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return profession, nil
}

func (p *professionService) Update(c context.Context, req *position_service.Profession) (*position_service.Profession, error) {
	profession, err := p.storage.Profession().Update(req)
	if err != nil {
		return nil, err
	}
	return profession, nil
}

func (p *professionService) Delete(c context.Context, req *position_service.ProfessionId) (*position_service.AffectedRows, error) {
	affectedrows, err := p.storage.Profession().Delete(req)

	if err != nil {
		return nil, err
	}

	return affectedrows, nil
}
