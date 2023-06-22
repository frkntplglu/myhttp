package url

import (
	"reflect"
	"testing"
)

func TestProtocolCheck(t *testing.T) {

	tests := []struct {
		name    string
		urlList []string
		want    []string
	}{
		{
			name:    "test for both without protocol and http",
			urlList: []string{"google.com", "example.com", "http://facebook.com", "http://twitter.com"},
			want:    []string{"http://google.com", "http://example.com", "http://facebook.com", "http://twitter.com"},
		},
		{
			name:    "test for both with http and https",
			urlList: []string{"http://example.org", "https://example.com"},
			want:    []string{"http://example.org", "https://example.com"},
		},
		{
			name:    "test for both without protocol and http and https",
			urlList: []string{"example.com", "http://example.com", "https://example.com"},
			want:    []string{"http://example.com", "http://example.com", "https://example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtocolCheck(tt.urlList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProtocolCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
