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

func (p *Penny) AddIncome(amount float64, title string) (*Income, error) {
	income := &Income{
		Title:  title,
		Amount: amount,
	}
	createdIncome, err := p.db.createIncome(income)
	return createdIncome, err
}

func (p *Penny) RemoveIncomeByID(id int) (*Income, error) {
	deletedIncome, err := p.db.deleteIncomeByID(id)
	return deletedIncome, err

}

// func (p *Penny) AddExpense(amount float64, title string) {
// 	p.Expenses = append(p.Expenses, &Expense{
// 		Title:     title,
// 		Amount:    amount,
// 		CreatedAt: time.Now(),
// 	})
// }

func (p *Penny) GetIncomes() ([]*Income, error) {
	incomes, err := p.db.getAllIncomes()
	if err != nil {
		return nil, err
	}
	return incomes, nil
}
