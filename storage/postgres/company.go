package postgres

import (
	"position_service/genproto/position_service"
	"position_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type companyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) repo.CompanyRepoI {
	return &companyRepo{
		db: db,
	}
}

func (r *companyRepo) Create(req *position_service.CreateCompany) (string, error) {
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

	query := `INSERT INTO company ( id, name) VALUES($1,$2)`

	_, err = r.db.Exec(query, id, req.Name)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *companyRepo) GetAll(req *position_service.GetAllCompanyRequest) (*position_service.GetAllCompanyResponse, error) {
	var (
		filter   string
		count    int32
		companys []*position_service.Company
	)
	args := make(map[string]interface{})
	if req.Name != "" {
		filter += " AND name ILIKE '%' || :na || '%' "
		args["na"] = req.Name
	}
	filter += " LIMIT :limi OFFSET :offs"
	args["limi"] = req.Limit
	args["offs"] = req.Offset

	countQuery := `SELECT count(1) FROM company WHERE true` + filter

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
			FROM company WHERE true` + filter

	rows, err = r.db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var company position_service.Company
		err = rows.Scan(
			&company.Id,
			&company.Name,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		companys = append(companys, &company)
	}

	return &position_service.GetAllCompanyResponse{
		Companys: companys,
		Count:    count,
	}, nil
}

func (r *companyRepo) GetById(id string) (*position_service.Company, error) {
	var company position_service.Company
	query := `SELECT id,name,created_at,updated_at FROM company WHERE id =$1`

	rows, err := r.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	rows.Next()

	err = rows.Scan(
		&company.Id,
		&company.Name,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &position_service.Company{
		Id:        company.Id,
		Name:      company.Name,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}, nil
}

func (r *companyRepo) Update(req *position_service.Company) (*position_service.Company, error) {
	_, err := r.db.Exec(
		`UPDATE company SET name=$1,updated_at = Now() WHERE id=$2 `,
		req.Name,
		req.Id,
	)
	if err != nil {
		return nil, err
	}
	return &position_service.Company{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}

func (r *companyRepo) Delete(req *position_service.CompanyId) (*position_service.AffectedRow, error) {
	resp, err := r.db.Exec(
		`DELETE FROM company WHERE id = $1`,
		req.Id,
	)

	if err != nil {
		return nil, err
	}

	affectedrows, err := resp.RowsAffected()

	if err != nil {
		return nil, err
	}
	return &position_service.AffectedRow{
		Number: affectedrows,
	}, nil
}
