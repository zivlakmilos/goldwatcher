package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepositoryMigrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepositoryInsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.Id <= 0 {
		t.Error("invalid id sent back:", result.Id)
	}
}

func TestSQLiteRepositoryAllHoldings(t *testing.T) {
	result, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed:", err)
	}

	if len(result) != 1 {
		t.Error("expected 1 row but got", len(result))
	}
}

func TestSQLiteRepositoryGetHoldingById(t *testing.T) {
	result, err := testRepo.GetHoldingById(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	if result.PurchasePrice != 1000 {
		t.Error("expected purchase price 1000 but got", result.PurchasePrice)
	}

	_, err = testRepo.GetHoldingById(2)
	if err == nil {
		t.Error("get one returned for non existing id")
	}
}

func TestSQLiteRepositoryUpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingById(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	h.PurchasePrice = 1011

	err = testRepo.UpdateHolding(1, *h)
	if err != nil {
		t.Error("update failed:", err)
	}
}

func TestSQLiteRepositoryDeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding:", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error while try to delete non existing id")
	}
}
