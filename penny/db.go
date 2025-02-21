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
	if err := ps.createExpensesTable(); err != nil {
		return err
	}
	return nil
}

func (ps *postgresStorage) createIncomesTable() error {
	query := `CREATE TABLE IF NOT EXISTS penny_incomes (
        id SERIAL PRIMARY KEY,
        title varchar(255) NOT NULL,
        amount REAL NOT NULL,
        user_id INTEGER NOT NULL,
        created_at TIMESTAMPTZ DEFAULT(now())
    )`
	_, err := ps.db.Query(query)
	return err
}
func (ps *postgresStorage) getAllIncomes(userID int) ([]*Income, error) {
	query := `SELECT id, title, amount, user_id, created_at FROM penny_incomes WHERE user_id = $1`
	rows, err := ps.db.Query(query, userID)
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
        title, amount, user_id
        ) VALUES (
            $1, $2, $3
        ) RETURNING id, title, amount, user_id, created_at
    `
	rows, err := ps.db.Query(query, income.Title, income.Amount, income.UserID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoIncomes(rows)
}
func (ps *postgresStorage) deleteIncomeByID(id, userID int) (*Income, error) {
	query := `DELETE FROM penny_incomes WHERE id = $1 AND user_id = $2 RETURNING
        id, title, amount, user_id, created_at
    `
	rows, err := ps.db.Query(query, id, userID)
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
		&income.UserID,
		&income.CreatedAt,
	)
	return income, err
}

func (ps *postgresStorage) createExpensesTable() error {
	query := `CREATE TABLE IF NOT EXISTS penny_expenses (
        id SERIAL PRIMARY KEY,
        title varchar(255) NOT NULL,
        amount REAL NOT NULL,
        user_id INTEGER NOT NULL,
        created_at TIMESTAMPTZ DEFAULT(now())
    )`
	_, err := ps.db.Query(query)
	return err
}
func (ps *postgresStorage) getAllExpenses(userID int) ([]*Expense, error) {
	query := `SELECT id, title, amount, user_id, created_at FROM penny_expenses WHERE user_id = $1`
	rows, err := ps.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	expenses := []*Expense{}
	for rows.Next() {
		expense, err := scanIntoExpenses(rows)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func (ps *postgresStorage) createExpense(expense *Expense) (*Expense, error) {
	query := `INSERT INTO penny_expenses(
        title, amount, user_id
        ) VALUES (
            $1, $2, $3
        ) RETURNING id, title, amount, user_id, created_at
    `
	rows, err := ps.db.Query(query, expense.Title, expense.Amount, expense.UserID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoExpenses(rows)
}
func (ps *postgresStorage) deleteExpenseByID(id, userID int) (*Expense, error) {
	query := `DELETE FROM penny_expenses WHERE id = $1 AND user_id = $2 RETURNING
        id, title, amount, user_id, created_at
    `
	rows, err := ps.db.Query(query, id, userID)
	if err != nil {
		return nil, err
	}
	rows.Next()
	return scanIntoExpenses(rows)
}
func scanIntoExpenses(rows *sql.Rows) (*Expense, error) {
	expense := &Expense{}
	err := rows.Scan(
		&expense.ID,
		&expense.Title,
		&expense.Amount,
		&expense.UserID,
		&expense.CreatedAt,
	)
	return expense, err
}

func (ps *postgresStorage) getAllTypes() ([]*AllType, error) {
	query := `SELECT id, title, amount, user_id, 'income' AS type, created_at FROM penny_incomes
        UNION ALL
        SELECT id, title, amount, user_id, 'expense' AS type, created_at FROM penny_expenses
        ORDER BY created_at DESC
    `
	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	all := []*AllType{}
	for rows.Next() {
		record, err := scanIntoAllTypes(rows)
		if err != nil {
			return nil, err
		}
		all = append(all, record)
	}
	return all, nil

}
func scanIntoAllTypes(rows *sql.Rows) (*AllType, error) {
	record := &AllType{}
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
