package helper

import "testing"

func TestConvertIntToHexString(t *testing.T) {
	type args struct {
		name  string
		input int64
		want  string
	}

	tests := []args{
		{
			name:  "valid input",
			input: 1207,
			want:  "0x4b7",
		},
		{
			name:  "valid input",
			input: 16766980,
			want:  "0xffd804",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ConvertIntToHexString(tt.input)
			if tt.want != ans {
				t.Errorf("Expected '%s', but got '%s'", tt.want, ans)
			}
		})
	}
}

func TestConvertHexStringToInt(t *testing.T) {
	type args struct {
		name   string
		input  string
		want   int64
		hasErr bool
	}

	tests := []args{
		{
			name:   "valid input",
			input:  "0x4b7",
			want:   1207,
			hasErr: false,
		},
		{
			name:   "valid input",
			input:  "0xffd804",
			want:   16766980,
			hasErr: false,
		},
		{
			name:   "invalid input",
			input:  "test",
			want:   0,
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := ConvertHexStringToInt(tt.input)
			if tt.hasErr && err == nil {
				t.Errorf("Expected has error, but got no error")
			}
			if !tt.hasErr && tt.want != ans {
				t.Errorf("Expected '%d', but got '%d'", tt.want, ans)
			}
		})
	}
}
