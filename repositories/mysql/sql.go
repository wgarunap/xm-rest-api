package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wgarunap/xm-rest-api/config"
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/domain/repository"
)

var _ repository.Company = (*ramsqlDB)(nil)

var ErrDuplicateEntry = errors.New(`company code already exist`)
var ErrEntryNotFound = errors.New(`unable to find the company code in the database`)

type ramsqlDB struct {
	db *sql.DB
}

func (r *ramsqlDB) Create(company domain.Company) error {
	exist := false
	raw := r.db.QueryRow(`SELECT EXISTS(SELECT 1 from test.company WHERE code= ? LIMIT 1);`, company.Code)
	err := raw.Scan(&exist)
	if err != nil {
		return err
	}
	if exist {
		return ErrDuplicateEntry
	}

	_, err = r.db.Exec(`INSERT INTO company (name, code,country,website,phone) VALUES (?, ?,?,?,?);`,
		company.Name,
		company.Code,
		company.Country,
		company.Website,
		company.Phone)

	if err != nil {
		return err
	}
	return nil
}
func (r *ramsqlDB) Update(company domain.Company) error {
	exist := false
	raw := r.db.QueryRow(`SELECT EXISTS(SELECT 1 from test.company WHERE code= ? LIMIT 1);`, company.Code)
	err := raw.Scan(&exist)
	if err != nil {
		return err
	}
	if !exist {
		return ErrEntryNotFound
	}

	_, err = r.db.Exec(`UPDATE company SET name=?,country=?,website=?,phone=? WHERE code=?;`,
		company.Name,
		company.Country,
		company.Website,
		company.Phone,
		company.Code)

	if err != nil {
		return err
	}
	return nil
}

func (r *ramsqlDB) Get(filters ...repository.Filter) ([]domain.Company, error) {
	query := `SELECT name, code, country,website, phone  FROM company WHERE`
	for i, filter := range filters {
		query = fmt.Sprintf(`%s %s='%v'`, query, filter.FieldName, filter.Value)
		if i != len(filters)-1 {
			query += " and "
		}
	}
	query += ";"

	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var com []domain.Company
	for res.Next() {
		tmp := domain.Company{}
		err := res.Scan(&tmp.Name, &tmp.Code, &tmp.Country, &tmp.Website, &tmp.Phone)
		if err != nil {
			return nil, err
		}
		com = append(com, tmp)
	}

	return com, err
}

func (r *ramsqlDB) Delete(filters ...repository.Filter) error {
	query := `DELETE FROM company WHERE`
	for i, filter := range filters {
		query = fmt.Sprintf(`%s %s='%v'`, query, filter.FieldName, filter.Value)
		if i != len(filters)-1 {
			query += " and "
		}
	}
	query += ";"

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *ramsqlDB) Close() error {
	return r.db.Close()
}

func NewRamDBInstance(cfg *config.Conf) repository.Company {
	db, err := sql.Open("mysql", cfg.DatabaseConnectUrl)
	if err != nil {
		panic(fmt.Sprintf("sql.Open : Error : %v", err))
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS  company (id int AUTO_INCREMENT , name VARCHAR(100) , code VARCHAR(100) , country VARCHAR(100),website VARCHAR(250), phone INTEGER(12),PRIMARY KEY(id));`)

	if err != nil {
		panic(fmt.Sprintf("sql.Exec: Error: %v", err))
	}

	return &ramsqlDB{db}
}
