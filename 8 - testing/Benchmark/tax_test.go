package BencharcTax //obrigatório ter o mesmo nome no package

import "testing"

// terminal$>> go test .
// terminal | gerar cobertura $>> go test --coverprofile=coverage.out
// terminal$ | gerar hmtl >> go tool cover -html=coverage.out
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but get %f", expected, result)
	}
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
		if result != item.expected {
			t.Errorf("Expected %f but get %f", item.expected, result)
		}
	}
}

// terminal: Rodar benchmarch $ >> got test --bench=. -run=^#
// $ >> got test --bench=. -run=^# --count=10 (ideal pegar média)
// $ >> got test --bench=. -run=^# --benchtime=3s (per periodof)
// $ >> got test --bench=. -run=^# --benchmen (analise de memoria)
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTaxSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500.0)
	}
}
