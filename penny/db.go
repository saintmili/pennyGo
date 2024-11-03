package penny

import (
	"database/sql"
)

type postgresStorage struct {
	db *sql.DB
}

func (ps *postgresStorage) init() error {
	if err := ps.createIncomesTable(); err != nil {
		return err
	}
	return nil
}

func (ps *postgresStorage) createIncomesTable() error {
	query := `CREATE TABLE IF NOT EXISTS penny_incomes (
        id SERIAL PRIMARY KEY,
        title varchar(255) NOT NULL,
        amount REAL NOT NULL,
        created_at TIMESTAMPTZ DEFAULT(now())
    )`
	_, err := ps.db.Query(query)
	return err
}
func (ps *postgresStorage) getAllIncomes() ([]*Income, error) {
	query := `SELECT id, title, amount, created_at FROM penny_incomes`
	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	incomes := []*Income{}
	for rows.Next() {
		income, err := scanIntoIncomes(rows)
		if err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}
	return incomes, nil
}
func (ps *postgresStorage) createIncome(income *Income) (*Income, error) {
	query := `INSERT INTO penny_incomes(
        title, amount
        ) VALUES (
            $1, $2
        ) RETURNING id, title, amount, created_at
    `
	rows, err := ps.db.Query(query, income.Title, income.Amount)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoIncomes(rows)

}
func (ps *postgresStorage) deleteIncomeByID(id int) (*Income, error) {
	query := `DELETE FROM penny_incomes WHERE id = $1 RETURNING
        id, title, amount, created_at
    `
	rows, err := ps.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoIncomes(rows)
}
func scanIntoIncomes(rows *sql.Rows) (*Income, error) {
	income := &Income{}
	err := rows.Scan(
		&income.ID,
		&income.Title,
		&income.Amount,
		&income.CreatedAt,
	)
	return income, err
}
