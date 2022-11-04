package TestFuzzing //obrigatório ter o mesmo nome no package

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

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 100.0, 1501.0}
	for _, am := range seed {
		f.Add(am)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f expectedd 0", result)
		}
	})
}
