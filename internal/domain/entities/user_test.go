package entities

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)


func TestUser_JSONMarshaling(t *testing.T) {
	u := User{
		Id: "a123",
		Email: "test@example.com",
		PasswordHash: "hashed_password",
		Login: "test_user",
		CreatedAt: time.Date(2026, time.August, 1, 15, 8, 12, 12, time.UTC),
		UpdatedAt: time.Date(2026, time.August, 1, 15, 8, 12, 12, time.UTC),
	}

	data, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}

	var restored User
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(u, restored) {
		t.Errorf("mismatch: %+v vs %+v", u, restored)
	}
}
