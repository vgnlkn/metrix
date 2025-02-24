package metrix

import "testing"

func TestMemStorage_Update(t *testing.T) {
	storage := NewMemStorage()
	type args struct {
		name  string
		value string
		vType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Update Gauge OK",
			args: args{
				name:  "someValue",
				value: "3.14",
				vType: "gauge",
			},
			wantErr: false,
		},
		{
			name: "Update Counter OK",
			args: args{
				name:  "someValue",
				value: "3",
				vType: "counter",
			},
			wantErr: false,
		},
		{
			name: "Update Gauge Error",
			args: args{
				name:  "someValue",
				value: "asd",
				vType: "gauge",
			},
			wantErr: true,
		},
		{
			name: "Update Counter Error",
			args: args{
				name:  "someValue",
				value: "asd",
				vType: "counter",
			},
			wantErr: true,
		},
		{
			name: "Update type Error",
			args: args{
				name:  "someValue",
				value: "asd",
				vType: "invalid_type",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := storage
			if err := m.Update(tt.args.name, tt.args.value, tt.args.vType); (err != nil) != tt.wantErr {
				t.Errorf("MemStorage.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
