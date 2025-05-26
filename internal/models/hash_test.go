package models

import (
	"testing"

	"github.com/RupenderSinghRathore/shortUrl/internal/assert"
)

func TestCreateHash(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			url:  "https://google.com",
			want: "99999",
		},
		{
			url:  "https://www.instagram.com",
			want: "81baa",
		},
	}

	for _, test := range tests {
		hash := CreateHash(test.url)
		assert.Equal(t, hash, test.want)
	}
}
