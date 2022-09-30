package postgres

import (
	"position_service/genproto/position_service"
	"position_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type professionRepo struct {
	db *sqlx.DB
}

func NewProfessionRepo(db *sqlx.DB) repo.ProfessionRepoI {
	return &professionRepo{
		db: db,
	}
}

func (r *professionRepo) Create(req *position_service.CreateProfession) (string, error) {
	var (
		id uuid.UUID
	)
	tx, err := r.db.Begin()
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
		return "", err
	}

	query := `INSERT INTO "profession"( id, name) VALUES($1,$2)`

	_, err = r.db.Exec(query, id, req.Name)

	if err != nil {
		return "", err
	}

	return id.String(), nil

}

func (r *professionRepo) GetAll(req *position_service.GetAllProfessionRequest) (*position_service.GetAllProfessionResponse, error) {

	var (
		filter      string
		count       int32
		professions []*position_service.Profession
	)
	args := make(map[string]interface{})
	if req.Name != "" {
		filter += " AND name ILIKE '%' || :na || '%' "
		args["na"] = req.Name
	}
	filter += " LIMIT :limi OFFSET :offs"
	args["limi"] = req.Limit
	args["offs"] = req.Offset

	countQuery := `SELECT count(1) FROM profession WHERE true` + filter

	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT 
			id,
			name,
			created_at,
			updated_at
			FROM profession WHERE true` + filter

	rows, err = r.db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var profession position_service.Profession
		err = rows.Scan(
			&profession.Id,
			&profession.Name,
			&profession.CreatedAt,
			&profession.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		professions = append(professions, &profession)
	}

	return &position_service.GetAllProfessionResponse{
		Professions: professions,
		Count:       count,
	}, nil

}

func (r *professionRepo) GetById(id string) (*position_service.Profession, error) {

	var profession position_service.Profession

	query := `SELECT id,name,created_at,updated_at FROM profession WHERE id = $1`

	rows, err := r.db.Query(query, id)
	if err != nil {

		return nil, err
	}
	rows.Next()
	err = rows.Scan(
		&profession.Id,
		&profession.Name,
		&profession.CreatedAt,
		&profession.UpdatedAt,
	)
	if err != nil {

		return nil, err
	}

	return &position_service.Profession{
		Id:        profession.Id,
		Name:      profession.Name,
		CreatedAt: profession.CreatedAt,
		UpdatedAt: profession.UpdatedAt,
	}, nil

}

func (r *professionRepo) Update(req *position_service.Profession) (*position_service.Profession, error) {
	_, err := r.db.Exec(
		`UPDATE profession SET name=$1,updated_at=Now() WHERE id=$2`,
		req.Name,
		req.Id,
	)
	if err != nil {
		return nil, err
	}
	return &position_service.Profession{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}

func (r *professionRepo) Delete(req *position_service.ProfessionId) (*position_service.AffectedRows, error) {

	resp, err := r.db.Exec(
		`DELETE FROM profession WHERE id = $1;`,
		req.Id,
	)

	if err != nil {

		return nil, err
	}

	affectedRows, err := resp.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &position_service.AffectedRows{
		Number: int32(affectedRows),
	}, nil

}
