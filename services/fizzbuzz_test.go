package services

import (
	"testing"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/stretchr/testify/assert"
)

func TestComputeFizzbuzz(t *testing.T) {
	testCases := []struct {
		name     string
		params   models.FizzbuzzParams
		expected string
	}{
		{
			name: "Basic FizzBuzz",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz",
		},
		{
			name: "Custom Strings",
			params: models.FizzbuzzParams{
				Int1:  2,
				Int2:  3,
				Limit: 10,
				Str1:  "Hello",
				Str2:  "World",
			},
			expected: "1,Hello,World,Hello,5,HelloWorld,7,Hello,World,Hello",
		},
		{
			name: "Single Int1",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  3,
				Limit: 9,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: "1,2,FizzBuzz,4,5,FizzBuzz,7,8,FizzBuzz",
		},
		{
			name: "Single Int2",
			params: models.FizzbuzzParams{
				Int1:  5,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: "1,2,3,4,FizzBuzz,6,7,8,9,FizzBuzz,11,12,13,14,FizzBuzz",
		},
	}

	fbs := NewFizzbuzzService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fbs.ComputeFizzbuzz(tc.params)
			assert.Equal(t, tc.expected, result)
		})
	}
}
