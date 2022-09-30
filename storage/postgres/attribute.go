package postgres

import (
	"errors"
	"fmt"
	"position_service/genproto/position_service"
	"position_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type attributeRepo struct {
	db *sqlx.DB
}

func isCorrectType(s string) bool {
	arr := [...]string{"text", "datetime", "number"}

	for _, i := range arr {
		if s == i {
			return true
		}
	}
	return false
}

func NewAttributeRepo(db *sqlx.DB) repo.AttributeRepoI {
	return &attributeRepo{
		db: db,
	}
}

func (a *attributeRepo) Create(req *position_service.CreateAttribute) (string, error) {
	if !isCorrectType(req.Type) {
		return "", errors.New("false type ")
	}

	var (
		id uuid.UUID
	)
	tx, err := a.db.Begin()
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err = uuid.NewRandom()

	if err != nil {
		return "", nil
	}

	query := `INSERT INTO attribute(id, name, "type") VALUES($1,$2,$3)`

	_, err = a.db.Exec(query, id, req.Name, req.Type)

	if err != nil {
		return "", nil
	}

	return id.String(), nil
}

func (a *attributeRepo) GetAll(req *position_service.GetAllAttributeRequest) (*position_service.GetAllAttributeResponse, error) {

	var (
		filter string
		count  int32
		arr    []*position_service.Attribute
	)

	args := make(map[string]interface{})

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :filter_name ||'%'"
		args["filter_name"] = req.Name
	}

	filter += " LIMIT :limi OFFSET :offs"
	args["limi"] = req.Limit
	args["offs"] = req.Offset
	countQuery := `SELECT count(1) FROM attribute WHERE true` + filter
	rows, err := a.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT id,name,type,created_at,updated_at FROM attribute WHERE true` + filter
	rows, err = a.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var attribute position_service.Attribute

		err = rows.Scan(&attribute.Id, &attribute.Name, &attribute.Type, &attribute.CreatedAt, &attribute.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		arr = append(arr, &attribute)
	}

	return &position_service.GetAllAttributeResponse{
		Attributes: arr,
		Count:      count,
	}, nil
}
func (a *attributeRepo) GetById(id string) (*position_service.Attribute, error) {
	query := `SELECT id,name,type, created_at, updated_at FROM attribute WHERE id = $1  `
	rows, err := a.db.Query(query, id)
	var attribute position_service.Attribute

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&attribute.Id, &attribute.Name, &attribute.Type, &attribute.CreatedAt, &attribute.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &position_service.Attribute{
		Id:        attribute.Id,
		Name:      attribute.Name,
		Type:      attribute.Type,
		CreatedAt: attribute.CreatedAt,
		UpdatedAt: attribute.UpdatedAt,
	}, nil
}

func (a *attributeRepo) Update(req *position_service.Attribute) (*position_service.AttributeAfterUpdate, error) {

	if !isCorrectType(req.Type) {
		return nil, errors.New("update wrong type")
	}

	old, err := a.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	_, err = a.db.Exec(
		`UPDATE attribute 
		SET name=$1 , type=$2 ,updated_at=Now()
		WHERE id=$3`,
		req.Name,
		req.Type,
		req.Id,
	)

	if err != nil {
		return nil, err
	}

	newone, err := a.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return &position_service.AttributeAfterUpdate{
		Old: old,
		New: newone,
	}, nil

}

func (a *attributeRepo) Delete(req *position_service.AttributeId) (*position_service.AttributeRowsAffected, error) {
	resp, err := a.db.Exec(
		`DELETE FROM attribute WHERE id = $1`,
		req.Id,
	)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &position_service.AttributeRowsAffected{
		RowsAffected: rowsAffected,
	}, nil
}
