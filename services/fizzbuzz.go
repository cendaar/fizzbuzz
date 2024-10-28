package services

import (
	"strconv"
	"strings"

	"github.com/cendaar/fizzbuzz/models"
)

type FizzbuzzServiceI interface {
	ComputeFizzbuzz(p models.FizzbuzzParams) string
}

type FizzbuzzService struct{}

func NewFizzbuzzService() *FizzbuzzService {
	return &FizzbuzzService{}
}

func (fbs *FizzbuzzService) ComputeFizzbuzz(p models.FizzbuzzParams) string {
	var result strings.Builder

	for i := 1; i <= p.Limit; i++ {
		isInt1Multiple := i%p.Int1 == 0
		isInt2Multiple := i%p.Int2 == 0

		switch {
		case isInt1Multiple && isInt2Multiple:
			result.WriteString(p.Str1 + p.Str2)
		case isInt1Multiple:
			result.WriteString(p.Str1)
		case isInt2Multiple:
			result.WriteString(p.Str2)
		default:
			result.WriteString(strconv.Itoa(i))
		}

		if i != p.Limit {
			result.WriteByte(',')
		}
	}

	return result.String()
}
