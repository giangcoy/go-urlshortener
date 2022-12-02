package generator

import (
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	g := NewGenerator()
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Case success",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := g.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func Benchmark_Generate(b *testing.B) {
	g := NewGenerator()
	for n := 0; n < b.N; n++ {
		g.Generate()
	}
}
