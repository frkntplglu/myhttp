package hash

import "testing"

func TestGetMd5Hash(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Non empty string",
			value: "test1234",
			want:  "16d7a4fca7442dda3ad93c9a726597e4",
		},
		{
			name:  "Numeric string",
			value: "12345",
			want:  "827ccb0eea8a706c4c34a16891f84e7b",
		},
		{
			name:  "Empty string",
			value: "",
			want:  "d41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMd5Hash(tt.value); got != tt.want {
				t.Errorf("GetMd5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
