package repo

import "position_service/genproto/position_service"

type AttributeRepoI interface {
	Create(req *position_service.CreateAttribute) (string, error)
	GetAll(req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error)
	GetById(id string) (*position_service.Attribute, error)
	Update(req *position_service.Attribute) (*position_service.AttributeAfterUpdate, error)
	Delete(req *position_service.AttributeId) (*position_service.AttributeRowsAffected, error)
}
