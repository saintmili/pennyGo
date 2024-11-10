package penny

import (
	"database/sql"
	"time"
)

type Penny struct {
	db *postgresStorage
}

type Income struct {
	ID        int
	Title     string
	Amount    float64
	UserID    int
	CreatedAt time.Time
}

type Expense = Income

func New(database *sql.DB) (*Penny, error) {
	postgresDB := &postgresStorage{
		db: database,
	}
	err := postgresDB.init()
	return &Penny{
		db: postgresDB,
	}, err
}

func (p *Penny) AddIncome(amount float64, title string, userID int) (*Income, error) {
	income := &Income{
		Title:  title,
		Amount: amount,
		UserID: userID,
	}
	createdIncome, err := p.db.createIncome(income)
	return createdIncome, err
}
func (p *Penny) RemoveIncomeByID(id, userID int) (*Income, error) {
	deletedIncome, err := p.db.deleteIncomeByID(id, userID)
	return deletedIncome, err

}
func (p *Penny) GetIncomes(userID int) ([]*Income, error) {
	incomes, err := p.db.getAllIncomes(userID)
	if err != nil {
		return nil, err
	}
	return incomes, nil
}

func (p *Penny) AddExpense(amount float64, title string, userID int) (*Expense, error) {
	expense := &Expense{
		Title:  title,
		Amount: amount,
		UserID: userID,
	}
	createdExpense, err := p.db.createExpense(expense)
	return createdExpense, err
}
func (p *Penny) RemoveExpenseByID(id, userID int) (*Expense, error) {
	deletedExpense, err := p.db.deleteExpenseByID(id, userID)
	return deletedExpense, err

}
func (p *Penny) GetExpenses(userID int) ([]*Expense, error) {
	expenses, err := p.db.getAllExpenses(userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}
