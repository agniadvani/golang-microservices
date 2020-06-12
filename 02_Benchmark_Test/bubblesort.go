package benchmark

import "sort"

func BubbleSort(elements []int) {
	keeprunning := true
	for keeprunning {
		keeprunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keeprunning = true
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
		return
	}
	sort.Ints(elements)
}
