package penny

import "time"

type Penny struct {
	Incomes  []*Income
	Expenses []*Expense
}

type Income struct {
	ID        int
	Title     string
	Amount    float64
	CreatedAt time.Time
}

type Expense = Income

func (p *Penny) Init() *Penny {
	return &Penny{}
}

func (p *Penny) AddIncome(amount float64, title string) {
	p.Incomes = append(p.Incomes, &Income{
		Title:     title,
		Amount:    amount,
		CreatedAt: time.Now(),
	})
}

func (p *Penny) AddExpense(amount float64, title string) {
	p.Expenses = append(p.Expenses, &Expense{
		Title:     title,
		Amount:    amount,
		CreatedAt: time.Now(),
	})
}
