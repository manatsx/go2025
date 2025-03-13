package main

import "testing"

// TestCountingValleys contiene varios casos de prueba para countingValleys
func TestCountingValleys(t *testing.T) {
	// Definimos un conjunto de casos de prueba
	tests := []struct {
		name string
		path string
		want int32
	}{
		{
			name: "Caso básico 1",
			path: "UDDDUDUU", // Cruza un valle
			want: 1,
		},
		{
			name: "Caso básico 2",
			path: "DDUUUUDD", // Cruza un valle
			want: 1,
		},
		{
			name: "Sin valles",
			path: "UUUUDDDDDD", // No cruza valles
			want: 0,
		},
		{
			name: "Un solo valle",
			path: "UDDDUU", // Cruza un solo valle
			want: 1,
		},
	}

	// Ejecutar cada prueba
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countingValleys(int32(len(tt.path)), tt.path)
			if got != tt.want {
				t.Errorf("countingValleys(%v) = %v; want %v", tt.path, got, tt.want)
			}
		})
	}
}
