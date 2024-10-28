package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidParams = errors.New("invalid request params")
)

var validate = validator.New()

type FizzbuzzParams struct {
	Int1  int    `json:"int1" validate:"required,min=1"`
	Int2  int    `json:"int2" validate:"required,min=1"`
	Limit int    `json:"limit" validate:"required,min=1,max=9999999"`
	Str1  string `json:"str1"  validate:"required"`
	Str2  string `json:"str2"  validate:"required"`
}

func (f *FizzbuzzParams) Validate() error {
	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("%w: parameters doesn't match the required validator parameters", ErrInvalidParams)
	}

	if f.Int1 == f.Int2 {
		return fmt.Errorf("%w: invalid int1 is equal to int2", ErrInvalidParams)
	}

	return nil
}

func (f *FizzbuzzParams) String() string {
	return fmt.Sprintf("%d-%d-%d-%s-%s", f.Int1, f.Int2, f.Limit, f.Str1, f.Str2)
}
