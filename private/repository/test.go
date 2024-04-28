package repository

import "time"

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (t *TestRepository) Migrate() error {
	return nil
}

func (t *TestRepository) InsertHolding(h Holdings) (*Holdings, error) {
	return &h, nil
}

func (t *TestRepository) AllHoldings() ([]Holdings, error) {
	var all []Holdings

	h := Holdings{
		Id:            1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	all = append(all, h)

	h = Holdings{
		Id:            1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 2000,
	}
	all = append(all, h)

	return all, nil
}

func (t *TestRepository) GetHoldingById(id int64) (*Holdings, error) {
	h := &Holdings{
		Id:            1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	return h, nil
}

func (t *TestRepository) UpdateHolding(id int64, updated Holdings) error {
	return nil
}

func (t *TestRepository) DeleteHolding(id int64) error {
	return nil
}
