package utils

import(
	"sort"
)

func BubbleSort(t []int) []int {

	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(t)-1; i++ {
			if t[i] > t[i+1] {
				t[i], t[i+1] = t[i+1], t[i]
				keepRunning = true
			}
		}
	}

	return t
}

func Sort(elements []int) {
	if(len(elements) > 1000) {
		sort.Ints(elements)
		return
	}

	BubbleSort(elements)
	return
}
