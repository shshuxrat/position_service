package service

import (
	"context"
	"position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"

	"github.com/jmoiron/sqlx"
)

type attributeService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	position_service.UnimplementedAttributeServiceServer
}

func NewAttributeService(log logger.LoggerI, db *sqlx.DB) *attributeService {
	return &attributeService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (a *attributeService) Create(c context.Context, req *position_service.CreateAttribute) (*position_service.AttributeId, error) {
	attributeId, err := a.storage.Attribute().Create(req)
	if err != nil {
		return nil, err
	}

	return &position_service.AttributeId{
		Id: attributeId,
	}, nil

}

func (a *attributeService) GetAll(c context.Context, req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error) {
	attributes, err := a.storage.Attribute().GetAll(req)

	if err != nil {
		return nil, err
	}
	return attributes, nil
}

func (a *attributeService) GetById(c context.Context, req *position_service.AttributeId) (*position_service.Attribute, error) {
	attribute, err := a.storage.Attribute().GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return attribute, nil
}

func (a *attributeService) Update(c context.Context, req *position_service.Attribute) (*position_service.AttributeAfterUpdate, error) {
	attributes, err := a.storage.Attribute().Update(req)

	if err != nil {
		return nil, err
	}

	return attributes, nil
}

func (a *attributeService) Delete(c context.Context, req *position_service.AttributeId) (*position_service.AttributeRowsAffected, error) {
	affectedRows, err := a.storage.Attribute().Delete(req)

	if err != nil {
		return nil, err
	}

	return affectedRows, nil
}
