package memstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgnlkn/metrix/internal/entity"
)

func TestMemStorageUpdate(t *testing.T) {
	storage := NewMemStorage()
	tests := []struct {
		name    string
		wantErr bool
		metric  *entity.Metrics
	}{
		{
			name:    "UpdateGaugeOK",
			wantErr: false,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3.14"),
				Type: entity.TypeGauge,
			},
		},
		{
			name:    "UpdateCounterOK",
			wantErr: false,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3"),
				Type: entity.TypeCounter,
			},
		},
		{
			name:    "UpdateUnknownType",
			wantErr: true,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3.14"),
				Type: "unknown",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := storage.UpdateMetrics(tt.metric); (err != nil) != tt.wantErr {
				t.Errorf("storage.UpdateMetrics() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemStorageCreate(t *testing.T) {
	storage := NewMemStorage()
	tests := []struct {
		name    string
		wantErr bool
		metric  *entity.Metrics
	}{
		{
			name:    "CreateOK",
			wantErr: false,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3.14"),
				Type: entity.TypeGauge,
			},
		},
		{
			name:    "CreateWithExistingName",
			wantErr: true,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3.14"),
				Type: entity.TypeGauge,
			},
		},
		{
			name:    "CreateWithInvalidType",
			wantErr: true,
			metric: &entity.Metrics{
				Name: "some",
				Val:  entity.NewCounterValue("3.14"),
				Type: "invalid type",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := storage.CreateMetrics(tt.metric); (err != nil) != tt.wantErr {
				t.Errorf("storage.CreateMetrics() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemStorageFind(t *testing.T) {
	storage := NewMemStorage()
	storage.CreateMetrics(&entity.Metrics{
		Name: "some",
		Val:  entity.NewCounterValue("3"),
		Type: entity.TypeCounter,
	})

	tests := []struct {
		name        string
		wantErr     bool
		mName       string
		mType       string
		expectedVal string
	}{
		{
			name:        "FindNotExisting",
			wantErr:     true,
			mName:       "not exist",
			expectedVal: "",
			mType:       entity.TypeCounter,
		},
		{
			name:        "FindOK",
			wantErr:     false,
			mName:       "some",
			expectedVal: "3",
			mType:       entity.TypeCounter,
		},
		{
			name:        "FindInvalidType",
			wantErr:     true,
			mName:       "some",
			expectedVal: "",
			mType:       "invalid type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := storage.FindMetrics(tt.mName, tt.mType)
			if (err != nil) != tt.wantErr {
				t.Errorf("storage.FindMetrics() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, val, tt.expectedVal)
		})
	}
}

func TestMemStorageAll(t *testing.T) {
	storage := NewMemStorage()
	err1 := storage.CreateMetrics(&entity.Metrics{
		Name: "some",
		Val:  entity.NewCounterValue("3"),
		Type: entity.TypeCounter,
	})

	err2 := storage.CreateMetrics(&entity.Metrics{
		Name: "some",
		Val:  entity.NewCounterValue("3"),
		Type: entity.TypeGauge,
	})

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Len(t, storage.All(), 2)
}
