package services

import (
	"strconv"
	"strings"

	"github.com/cendaar/fizzbuzz/models"
)

// FizzbuzzServiceI defines an interface for the FizzbuzzService.
// This makes it easy to mock and test the service separately.
type FizzbuzzServiceI interface {
	ComputeFizzbuzz(p models.FizzbuzzParams) string
}

// FizzbuzzService is a struct implementing the FizzbuzzServiceI interface.
// It's the concrete implementation of the fizzbuzz logic.
type FizzbuzzService struct{}

// NewFizzbuzzService creates and returns a new instance of FizzbuzzService.
// This function allows for dependency injection and separation of concerns
func NewFizzbuzzService() *FizzbuzzService {
	return &FizzbuzzService{}
}

// ComputeFizzbuzz generates a FizzBuzz sequence up to the specified limit.
// It replaces multiples of Int1 and Int2 with the corresponding strings Str1 and Str2, respectively.
// If a number is a multiple of both Int1 and Int2, it combines Str1 and Str2.
func (fbs *FizzbuzzService) ComputeFizzbuzz(p models.FizzbuzzParams) string {
	// Use a strings.Builder to efficiently build the final result as a single string.
	var result strings.Builder

	// Loop from 1 up to the limit specified in the FizzbuzzParams.
	for i := 1; i <= p.Limit; i++ {
		// Determine if the current number is a multiple of Int1 and/or Int2.
		isInt1Multiple := i%p.Int1 == 0
		isInt2Multiple := i%p.Int2 == 0

		// Append the appropriate string to the result based on the FizzBuzz rules.
		switch {
		// If the number is a multiple of both Int1 and Int2, append Str1Str2.
		case isInt1Multiple && isInt2Multiple:
			result.WriteString(p.Str1 + p.Str2)
		// If the number is a multiple of Int1 only, append Str1.
		case isInt1Multiple:
			result.WriteString(p.Str1)
		// If the number is a multiple of Int2 only, append Str2.
		case isInt2Multiple:
			result.WriteString(p.Str2)
		// If the number is not a multiple of Int1 or Int2, append the number itself.
		default:
			result.WriteString(strconv.Itoa(i))
		}

		// Add a comma after each entry, except for the last one.
		if i != p.Limit {
			result.WriteByte(',')
		}
	}

	return result.String()
}
