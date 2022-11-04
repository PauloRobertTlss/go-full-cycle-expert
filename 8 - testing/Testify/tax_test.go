package TestifyGo //obrigatório ter o mesmo nome no package

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// terminal$>> go test .
// terminal | gerar cobertura $>> go test --coverprofile=coverage.out
// terminal$ | gerar hmtl >> go tool cover -html=coverage.out
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)
	assert.Equal(t, expected, result)
}

func TestCalculateTaxInBatch(t *testing.T) {
	type Taxs struct {
		amount   float64
		expected float64
	}
	table := []Taxs{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1001.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		assert.Equal(t, item.expected, result)
	}
}
