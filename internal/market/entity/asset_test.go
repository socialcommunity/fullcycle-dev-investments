package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewAsset(test *testing.T) {

	uuid := uuid.New().String()
	name := "test"
	marketVolume := 10000

	got := NewAsset(uuid, name, marketVolume)

	if got.ID != uuid {
		test.Errorf("NewAsset() = %v, want %v", got.ID, uuid)
	}
}
