package models_test

import (
	"testing"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzParams_Validate(t *testing.T) {
	testCases := []struct {
		name            string
		params          models.FizzbuzzParams
		expectedErr     error
		expectedErrMssg string
	}{
		{
			name: "Valid parameters",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr: nil,
		},
		{
			name: "Invalid Int1 (0)",
			params: models.FizzbuzzParams{
				Int1:  0,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Invalid Int2 (0)",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  0,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Invalid Limit (0)",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 0,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Invalid Limit (too high)",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 10000000,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Empty Str1",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Empty Str2",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: parameters doesn't match the required validator parameters",
		},
		{
			name: "Int1 equals Int2",
			params: models.FizzbuzzParams{
				Int1:  3,
				Int2:  3,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expectedErr:     models.ErrInvalidParams,
			expectedErrMssg: "invalid request params: invalid int1 is equal to int2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tc.expectedErr)
				assert.EqualError(t, err, tc.expectedErrMssg)
			}
		})
	}
}
