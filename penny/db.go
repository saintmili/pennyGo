package penny

import (
	"database/sql"
)

type postgresStorage struct {
	db *sql.DB
}

func (ps *postgresStorage) init() error {
	if err := ps.createRecordsTable(); err != nil {
		return err
	}
	return nil
}

func (ps *postgresStorage) createRecordsTable() error {
	query := `CREATE TABLE IF NOT EXISTS penny_records (
        id SERIAL PRIMARY KEY,
        title varchar(255) NOT NULL,
        amount REAL NOT NULL,
        user_id INTEGER NOT NULL,
        type varchar(255) NOT NULL DEFAULT 'income',
        created_at TIMESTAMPTZ DEFAULT(now())
    )`
	_, err := ps.db.Query(query)
	return err
}
func (ps *postgresStorage) getAllRecords() ([]*Record, error) {
	query := `SELECT id, title, amount, user_id, type, created_at FROM penny_records
        ORDER BY created_at DESC
    `
	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	all := []*Record{}
	for rows.Next() {
		record, err := scanIntoRecord(rows)
		if err != nil {
			return nil, err
		}
		all = append(all, record)
	}
	return all, nil

}
func (ps *postgresStorage) createRecord(record *Record) (*Record, error) {
	query := `INSERT INTO penny_records(
        title, amount, user_id, type
        ) VALUES (
            $1, $2, $3, $4
        ) RETURNING id, title, amount, user_id, type, created_at
    `
	rows, err := ps.db.Query(query, record.Title, record.Amount, record.UserID, record.Type)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoRecord(rows)
}
func (ps *postgresStorage) deleteRecordByID(id, userID int) (*Record, error) {
	query := `DELETE FROM penny_records WHERE id = $1 AND user_id = $2 RETURNING
        id, title, amount, user_id, type, created_at
    `
	rows, err := ps.db.Query(query, id, userID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoRecord(rows)
}
func (ps *postgresStorage) updateRecordByID(record *Record) (*Record, error) {
	query := `UPDATE penny_records SET
        title = $1, amount = $2, user_id = $3, type = $4
        WHERE id = $5
        RETURNING id, title, amount, user_id, type, created_at
    `
	rows, err := ps.db.Query(query, record.Title, record.Amount, record.UserID, record.Type, record.ID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoRecord(rows)
}
func (ps *postgresStorage) getAllExpenses(userID int) ([]*Record, error) {
	query := `SELECT id, title, amount, user_id, type, created_at FROM penny_records WHERE user_id = $1 AND type = "expense"`
	rows, err := ps.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	expenses := []*Record{}
	for rows.Next() {
		expense, err := scanIntoRecord(rows)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func (ps *postgresStorage) getAllIncomes(userID int) ([]*Record, error) {
	query := `SELECT id, title, amount, user_id, type, created_at FROM penny_records WHERE user_id = $1 AND type = "income"`
	rows, err := ps.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	incomes := []*Record{}
	for rows.Next() {
		income, err := scanIntoRecord(rows)
		if err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}
	return incomes, nil
}
func scanIntoRecord(rows *sql.Rows) (*Record, error) {
	record := &Record{}
	err := rows.Scan(
		&record.ID,
		&record.Title,
		&record.Amount,
		&record.UserID,
		&record.Type,
		&record.CreatedAt,
	)
	return record, err
}
