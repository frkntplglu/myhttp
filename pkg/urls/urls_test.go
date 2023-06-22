package urls

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

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{

		{
			url:  "https",
			want: false,
		},
		{
			url:  "https://",
			want: false,
		},
		{
			url:  "",
			want: false,
		},
		{
			url:  "http://www",
			want: true,
		},

		{
			url:  "adjust.com",
			want: false,
		},
		{
			url:  "http://www.adjust.com",
			want: true,
		},
		{
			url:  "https://www.adjust.com/tr/",
			want: true,
		},
		{
			url:  "/testing-path",
			want: false,
		},
		{
			url:  "alskjff#?asf//dfas",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("Test for different strings", func(t *testing.T) {
			if got := IsValidUrl(tt.url); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
