package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		got := Sum(given)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, given)
		}
	})
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{32, 235, -123, 2223})
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	result := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(result)
	// Output: [3 9]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 23}, []int{9899, 45}, []int{455, 231, -823})
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	result := SumAllTails([]int{1, 2}, []int{0, 2, 2, 3})
	fmt.Println(result)
	// Output: [2 7]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2}, []int{0, 2, 2, 3}, []int{12, 2, 3, 5})
	}
}
