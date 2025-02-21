package penny

import (
	"database/sql"
	"time"
)

type Penny struct {
	db *postgresStorage
}

type Record struct {
	ID        int
	Title     string
	Amount    float64
	UserID    int
	Type      string
	CreatedAt time.Time
}

func New(database *sql.DB) (*Penny, error) {
	postgresDB := &postgresStorage{
		db: database,
	}
	err := postgresDB.init()
	return &Penny{
		db: postgresDB,
	}, err
}

func (p *Penny) AddRecord(amount float64, title string, userID int, recordType string) (*Record, error) {
	income := &Record{
		Title:  title,
		Amount: amount,
		UserID: userID,
		Type:   recordType,
	}
	createdIncome, err := p.db.createRecord(income)
	return createdIncome, err
}
func (p *Penny) RemoveRecordByID(id, userID int) (*Record, error) {
	deletedIncome, err := p.db.deleteRecordByID(id, userID)
	return deletedIncome, err
}
func (p *Penny) GetIncomes(userID int) ([]*Record, error) {
	incomes, err := p.db.getAllIncomes(userID)
	if err != nil {
		return nil, err
	}
	return incomes, nil
}

func (p *Penny) GetExpenses(userID int) ([]*Record, error) {
	expenses, err := p.db.getAllExpenses(userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}
func (p *Penny) GetAllRecords() ([]*Record, error) {
	records, err := p.db.getAllRecords()
	if err != nil {
		return nil, err
	}
	return records, nil
}
func (p *Penny) UpdateRecordByID(record *Record) (*Record, error) {
	updatedRecord, err := p.db.updateRecordByID(record)
	if err != nil {
		return nil, err
	}
	return updatedRecord, err
}
