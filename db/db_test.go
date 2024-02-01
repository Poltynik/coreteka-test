package db

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInMemoryDB_SetAndCommitTransaction(t *testing.T) {
	type fields struct {
		key          string
		value        string
		expectedData map[string]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Set and commit transaction",
			fields: fields{
				key:   "key",
				value: "value",
				expectedData: map[string]any{
					"key": "value",
				},
			},
		},
		{
			name: "Set with empty value",
			fields: fields{
				key:          "",
				value:        "",
				expectedData: map[string]any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewInMemoryDB()
			db.StartTransaction()
			db.Set(tt.fields.key, tt.fields.value)
			db.CommitTransaction()

			assert.Equal(t, tt.fields.expectedData, db.mainStorage)
		})
	}
}

func TestInMemoryDB_DeleteInTransaction(t *testing.T) {
	type fields struct {
		key          string
		value        string
		expectedData map[string]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Delete in transaction",
			fields: fields{
				key:          "key",
				value:        "value",
				expectedData: map[string]any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewInMemoryDB()
			db.StartTransaction()
			db.Set(tt.fields.key, tt.fields.value)
			db.Delete(tt.fields.key)
			db.CommitTransaction()

			assert.Equal(t, tt.fields.expectedData, db.mainStorage)
		})
	}
}

func TestInMemoryDB_Delete(t *testing.T) {
	type fields struct {
		key          string
		expectedData map[string]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success delete",
			fields: fields{
				key:          "key",
				expectedData: map[string]any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewInMemoryDB()
			db.Delete(tt.fields.key)

			assert.Equal(t, tt.fields.expectedData, db.mainStorage)
		})
	}
}

func TestInMemoryDB_Get(t *testing.T) {
	type fields struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success get",
			fields: fields{
				key:   "key",
				value: "value",
			},
		},
		{
			name: "Get by empty key",
			fields: fields{
				key:   "",
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewInMemoryDB()
			db.Set(tt.fields.key, tt.fields.value)

			if got := db.Get(tt.fields.key); !reflect.DeepEqual(got, tt.fields.value) {
				t.Errorf("Get() = %v, want %v", got, tt.fields.value)
			}
		})
	}
}

func TestInMemoryDB_RollBackTransaction(t *testing.T) {
	type fields struct {
		key          string
		value        string
		newValue     string
		expectedData map[string]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success rollback",
			fields: fields{
				key:      "key",
				value:    "value",
				newValue: "new_value",
				expectedData: map[string]any{
					"key": "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewInMemoryDB()
			db.Set(tt.fields.key, tt.fields.value)
			db.StartTransaction()
			db.Set(tt.fields.key, tt.fields.newValue)
			db.RollBackTransaction()

			assert.Equal(t, tt.fields.expectedData, db.mainStorage)
		})
	}
}
