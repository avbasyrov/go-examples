package list

import "testing"

func BenchmarkNativeSlice_Default(b *testing.B) {
	var slice []int

	for n := 0; n < b.N; n++ {
		slice = append(slice, n)
	}
}

func BenchmarkNativeSlice_WithCapacity(b *testing.B) {
	slice := make([]int, 0, b.N)

	for n := 0; n < b.N; n++ {
		slice = append(slice, n)
	}
}
