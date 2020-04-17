package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	list := []int{9,8,7,6,5}

	BubbleSort(list)

	assert.NotNil(t, list)
	assert.Equal(t, 5, len(list))
	assert.Equal(t, 5, list[0])
	assert.Equal(t, 6, list[1])
	assert.Equal(t, 7, list[2])
	assert.Equal(t, 8, list[3])
	assert.Equal(t, 9, list[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	list := []int{5,6,7,8,9}

	BubbleSort(list)

	assert.NotNil(t, list)
	assert.Equal(t, 5, len(list))
	assert.Equal(t, 5, list[0])
	assert.Equal(t, 6, list[1])
	assert.Equal(t, 7, list[2])
	assert.Equal(t, 8, list[3])
	assert.Equal(t, 9, list[4])
}

func GetElements(n int) []int {
	els := make([]int, n)

	i := 0

	for j := n-1; j >= 0; j-- {
		els[i] = j
		i++
	}

	return els
}

func BenchmarkGetElements10(t *testing.B) {
	els := GetElements(10)

	for i := 0; i < t.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkGetElements1000(b *testing.B) {
	elements := GetElements(1000)

	for i := 0; i<b.N; i++ {
		
		BubbleSort(elements)
	}
}

func BenchmarkGetElements100000(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i< b.N; i++ {
		BubbleSort(elements)
	}
}

//Sorting using GoLang Native Sort method

func BenchmarkSort10(b *testing.B) {
	elements := GetElements(10)

	for i := 0 ; i<b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort1000(t *testing.B) {
	elements := GetElements(1000)

	for i := 0; i< t.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100000(t *testing.B) {
	elements := GetElements(100000)

	for i := 0; i<t.N; i++ {
		Sort(elements)
	}
}