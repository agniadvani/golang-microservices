package benchmark

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	els := []int{5, 4, 3, 2, 1}
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
	assert.EqualValues(t, 5, els[4])
}
func makeElements(len int) []int {
	els := make([]int, len)
	i := 0
	for j := len - 1; j >= 0; j-- {
		els[i] = j
		i++
	}
	return els
}
func TestMakeElements(t *testing.T) {
	els := makeElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])

}
func BenchmarkBubbleSort100(b *testing.B) {

	els := makeElements(100)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {

	els := makeElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {

	els := makeElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkNativeSort100(b *testing.B) {
	els := makeElements(100)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkNativeSort1000(b *testing.B) {
	els := makeElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkNativeSort100000(b *testing.B) {
	els := makeElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort50000(b *testing.B) {

	els := makeElements(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkNativeSort50000(b *testing.B) {
	els := makeElements(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
func BenchmarkSort100(b *testing.B) {
	els := makeElements(100)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
func BenchmarkSort100000(b *testing.B) {
	els := makeElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
